# Frequently Unasked Questions

## Q: How many requests per second can be handled?

You're asking the wrong the question. The strict answer to your question is "many thousands on a 2CPU/4G instance". Various handlers add their own latencies, but in general it's the endpoint roundtrip that serves the request content that adds the most.

## Q: How are Paths chosen?

Given a Request, Paths are walked sequentially based on their configuration, with the first match winning. The **path**, **headers**, **host**, **hosts**, and **methods** options are used in determining Path selection.

## Q: Why are there two ways of expressing a Path host?

Clarity *and* flexibility. Most often, a single **host** suffices, when not a list of **hosts** is available.

## Q: How do I serve a simple folder from the filesystem?

There doesn't exist a facility to "serve files" as this is a request router, first. It would be trivial to embed JAR into an application that did, however. (Or write a file-serving finisher)

## Q: Finisher or Pool?

In general, if the Path is going to end, you want a Finisher, if it's going to be proxied, you want a Pool. There are reasons where you need more logic than a Pool can offer, so a Finisher may work in conjunction with Pools to handle complex proxying (see URLswitch).

## Q: Handler or Finisher?

In general, if you're operating *on* a Request or a Response, versus serving the content of a Request, you want a Handler. Especially if you want something proxied, let a Pool do the heavy-lifting there, and write a Handler.

## Q: Why do Handlers take an http.Handler and return an http.Handler, instead of just *be* an http.Handler?

Wrapping or closure. By operating this way, the func that does all of the work isn't executed until it's needed, only the "wrapper" is executed to create the unique closed func.

## Q: Shouldn't the *Recoverer* handler be included by default?

Make it the first global handler, and it is. "Probably" is the answer, however it is not.

```yaml
handlers:
  - Recoverer
```

## Q: By default, how does a Pool work?

A minimally-configured Pool, as shown below, will operate as a fair load-balancer between all members, regardless of health, availability, or performance. 100/N percent of all requests will go to each member of an N-member pool, sequentially. Of note: the order of members is not necessarily indicative of the order requests are doled out.

```yaml
  somepool:
    Members:
      - http://10.0.0.1:80/
      - http://10.0.0.2:80/
```

## Q: Can I have a Pool with one member?

Absolutely. Pools are the basic "proxy" primitives so they can support any number of members with ease, and scale as needed transparently to the Paths that use them.

## Q: How do I configure a Pool, where I *do not* want existing sessions to fail over, but would rather they return an error?

You want the Pool to be **Sticky**, but don't set **Buffered** and **BufferedFails**, and don't **Prune**.

```yaml
  somepool:
    HealthCheckUri: /healthcheck
    Sticky: true
    Members:
      - http://10.0.0.1:80/
      - http://10.0.0.2:80/
```

## Q: How do I configure a pool, where failures are handled elegantly and the user *at most* sees a delay? Does this work with Sticky?

You want the Pool to be **Buffered** with **BufferedFails** set to at least *2*, and **Prune**. If using **Sticky**, this will also work (the cookie will be reissued on failover) as long as the underlying members can tolerate idempotent requests.

```yaml
  somepool:
    HealthCheckUri: /healthcheck
    Buffered: true
    BufferedFails: 2
    Prune: true
    Members:
      - http://10.0.0.1:80/
      - http://10.0.0.2:80/
```

## Q: Why should I "obfuscate" the stickycookie?

The value of the stickycookie is exactly what you have in the *Members* list. If it's clear, then it's at best telling people what your internal names or IP
addresses are, and at worst giving a potential attacker a point to start manipulating requests and seeing what they can break. Because of how member lookups are
boxed, requests will be limited to only the listed members, but in conjunction with other manipulations that on their own provide little-to-no value, could allow
someone to figure out a lot about your underlying infrastructure.

Different levels of obfuscation are available to meet your security desires.

## Q: This is mostly more than I need, is anything *required*?

Nope. Keep it simple, use the constructs you want, and reap the benefits around them. Everything has a default, and every default is intelligent. Below is my typical test config, with all of the debugging on and a single Path, which I extend as needed.

```yaml
debug: true
debugrequests: true
debugresponses: true
debugtimings: true

paths:
  -
    Path: /
    Name: catchall
    Finisher: Test
```

## Q: InitFuncs?

Go has a great facility in the ``func init()`` that is called *after* all of the ``var``s and ``const``s are built, but *before* ``func main()``, etc. In your ``func init()`` you may want to do some assignment or bootstrap that may depend on configuration or that other inits are executed. ``InitFuncs.Add()`` is your friend, because you can pass it a func that will get called at bootstrap time, which is *after* all of the inits, *after* the configs are loaded, but *before* the JAR server is instantiated. InitFuncs you add will be executed **at most once**.

## Q: When I use *Pool.EC2Affinity* without *Pool.Prune*, I see an error message: Why?

*Pool.Prune* empowers the Pool Lifeguard to quickly identify and take unhealthy or failed members out of the Pool, and put them back in after they've stabilized. This is extra important when using various EC2-awareness facilities, as the preference for AZ-local members is so high that you could lose large amounts of traffic if your AZ-local member fails or destabilizes and isn't pruned.

## Q: You use an awful lot of ``const``s instead of string values, why?

If I, for example, call an option ``"somethingspeledright"`` and you get it wrong (see what I did there?) it may be inobvious whether the "thing" is working or not, but if I define ``const ConfigSomethingSpelledRight`` and *you* mis-spell it, you'll get a compile-time error. Strategic use of ``const``s hardens the subsystems utilizing them (e.g. the configuration subsystem).

## Q: How do S3 Pools work?

If a pool has a member of ``s3://bucket`` then the Pool type is automatically set to S3, and EC2 IAM profiles are used. This implies running on an EC2 instance, and that the instance profile allows at least read access to the specified bucket.

## Q: How do websockets work?

If a pool has a member of ``ws://host:port/path`` then the Pool type is automatically set to WS, and websockets should be properly proxied. Not all Handlers may be websocket-safe, however, so less is more.

## Q: ZOMG JAR IS USING SO MUCH RAM?! WTF?!!

Calm down. While it is possible there's an engineering defect, it's unlikely. The test batteries are pretty keen on finding goro leaks, and in a goro-heavy
architecture like JAR, having a heap leak without a goro leak is gonna be pretty tough. More likely, this is a result of Go changing how it marks free memory. 

### TL;DR

Freed memory is marked as reclaimable, but puts no pressure on the kernel to *actually* free the memory. This can be switched to a different marking style, that fairly immediately reclaims memory. There's nothing wrong with leaving it as-is. JAR will continue to reuse the overallocated memory as-needed instead of requesting new pages, unless it actually needs new pages.

If this is concerning, you can change how the Go runtime marks free memory by setting ``GODEBUG=madvdontneed=1`` before you run your Go application.

### Why

Many languages like to give off an air of being "memory efficient". What this means to most people is "use little memory", which from an engineering perspective means "free memory quickly and get the RSS size down, so someone staring at ``top`` won't freak out." *Actual* memory efficiency is to let the memory manager decide what to do and when, and not waste time requesting N pages when you just freed >=N pages a moment ago. That's hard to sell to someone staring at ``top``, however.

#### DONTNEED

When you advise the kernel that you "don't need" a range of memory, the memory manager wipes it from your allocation (RSS no longer reports the freed range) even though you can still write to it (and thus have it added back to your RSS) it until the memory manager *actually* reclaims the memory at some point in the future. This is the most universally-compatible memory "freeing" advice on fairly modern memory management systems. Go prior to 1.12 used **MADV_DONTNEED**.

#### FREE

When you advise the kernel that a range is "free", the memory manager doesn't do a damn thing until it decides it needs to reclaim memory. The "free" memory can absolutely be written to again until the memory manager *actually* reclaims the memory. Because this advisory is a noop under normal conditions, applications using it appear to bloat, as they will stay at peak allocation until the memory manager decides it needs more pages, and *then* goes and reclaims it (and *then* RSS no longer report the freed memory). Go 1.12 uses **MADV_FREE** on supporting systems.

An interesting side effect of this, is it makes peak memory usage more apparent, as even a momentary need for a burst of memory will persist indefinitely on a healthy system, effectively showing a HWM for memory usage.