

# aws
`import "github.com/cognusion/go-jar/aws"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func GetAwsRegion() (region string)](#GetAwsRegion)
* [func GetAwsRegionE() (region string, err error)](#GetAwsRegionE)
* [func InitAWS(awsRegion, awsAccessKey, awsSecretKey string) (*session.Session, error)](#InitAWS)
* [func S3urlToParts(url string) (bucket, filePath, filename string)](#S3urlToParts)
* [type Session](#Session)
  * [func NewSession(awsRegion, awsAccessKey, awsSecretKey string) (*Session, error)](#NewSession)
  * [func (s *Session) BucketToFile(bucket, bucketPath, filename string) (size int64, err error)](#Session.BucketToFile)
  * [func (s *Session) BucketToWriter(bucket, bucketPath string, out io.Writer) (size int64, err error)](#Session.BucketToWriter)
  * [func (s *Session) BucketToWriterWithContext(ctx context.Context, bucket, bucketPath string, out io.Writer) (size int64, err error)](#Session.BucketToWriterWithContext)
  * [func (s *Session) BucketUpload(bucket, bucketPath string, file io.Reader) error](#Session.BucketUpload)
  * [func (s *Session) BucketUploadWithContext(ctx context.Context, bucket, bucketPath string, file io.Reader) error](#Session.BucketUploadWithContext)
  * [func (s *Session) GetInstanceAZByIP(ip string) (string, error)](#Session.GetInstanceAZByIP)
  * [func (s *Session) GetInstancesAZByIP(ips []*string) (*map[string]string, error)](#Session.GetInstancesAZByIP)


#### <a name="pkg-files">Package files</a>
[aws.go](https://github.com/cognusion/go-jar/tree/master/aws/aws.go)



## <a name="pkg-variables">Variables</a>
``` go
var (
    // DebugOut is a log.Logger for debug messages
    DebugOut = log.New(io.Discard, "", 0)
    // TimingOut is a log.Logger for timing-related debug messages. DEPRECATED
    TimingOut = log.New(io.Discard, "[TIMING] ", 0)
)
```


## <a name="GetAwsRegion">func</a> [GetAwsRegion](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=7391:7426#L274)
``` go
func GetAwsRegion() (region string)
```
GetAwsRegion returns the region as a string,
first consulting the well-known environment variables,
then falling back EC2 metadata calls



## <a name="GetAwsRegionE">func</a> [GetAwsRegionE](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=7630:7677#L282)
``` go
func GetAwsRegionE() (region string, err error)
```
GetAwsRegionE returns the region as a string and and error,
first consulting the well-known environment variables,
then falling back EC2 metadata calls



## <a name="InitAWS">func</a> [InitAWS](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=1669:1753#L66)
``` go
func InitAWS(awsRegion, awsAccessKey, awsSecretKey string) (*session.Session, error)
```
InitAWS optionally takes a region, accesskey and secret key,
setting AWSSession to the resulting session. If values aren't
provided, the well-known environment variables (WKE) are
consulted. If they're not available, and running in an EC2
instance, then it will use the local IAM role



## <a name="S3urlToParts">func</a> [S3urlToParts](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=8028:8093#L298)
``` go
func S3urlToParts(url string) (bucket, filePath, filename string)
```
S3urlToParts explodes an s3://bucket/path/file url into its parts




## <a name="Session">type</a> [Session](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=762:907#L33)
``` go
type Session struct {
    // AWS is the raw, hopefully initialized AWS Session
    AWS *session.Session
    Me  *ec2metadata.EC2InstanceIdentityDocument
}

```
Session is a container around an AWS Session, to make AWS operations easier







### <a name="NewSession">func</a> [NewSession](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=953:1032#L40)
``` go
func NewSession(awsRegion, awsAccessKey, awsSecretKey string) (*Session, error)
```
NewSession returns a Session or an error





### <a name="Session.BucketToFile">func</a> (\*Session) [BucketToFile](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=2741:2832#L107)
``` go
func (s *Session) BucketToFile(bucket, bucketPath, filename string) (size int64, err error)
```
BucketToFile copies a file from an S3 bucket to a local file




### <a name="Session.BucketToWriter">func</a> (\*Session) [BucketToWriter](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=3394:3492#L131)
``` go
func (s *Session) BucketToWriter(bucket, bucketPath string, out io.Writer) (size int64, err error)
```
BucketToWriter copies a file from an S3 bucket to a Writer




### <a name="Session.BucketToWriterWithContext">func</a> (\*Session) [BucketToWriterWithContext](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=3654:3784#L136)
``` go
func (s *Session) BucketToWriterWithContext(ctx context.Context, bucket, bucketPath string, out io.Writer) (size int64, err error)
```
BucketToWriterWithContext copies a file from an S3 bucket to a Writer




### <a name="Session.BucketUpload">func</a> (\*Session) [BucketUpload](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=4289:4368#L155)
``` go
func (s *Session) BucketUpload(bucket, bucketPath string, file io.Reader) error
```
BucketUpload uploads the file to the bucket/bucketPath




### <a name="Session.BucketUploadWithContext">func</a> (\*Session) [BucketUploadWithContext](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=4553:4664#L160)
``` go
func (s *Session) BucketUploadWithContext(ctx context.Context, bucket, bucketPath string, file io.Reader) error
```
BucketUploadWithContext uploads the file to the bucket/bucketPath, with the specified context




### <a name="Session.GetInstanceAZByIP">func</a> (\*Session) [GetInstanceAZByIP](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=5572:5634#L186)
``` go
func (s *Session) GetInstanceAZByIP(ip string) (string, error)
```
GetInstanceAZByIP returns an Availability Zone or an error




### <a name="Session.GetInstancesAZByIP">func</a> (\*Session) [GetInstancesAZByIP](https://github.com/cognusion/go-jar/tree/master/aws/aws.go?s=6322:6401#L222)
``` go
func (s *Session) GetInstancesAZByIP(ips []*string) (*map[string]string, error)
```
GetInstancesAZByIP returns a map of IPs to Availability Zones or an error








- - -
Generated by [godoc2md](http://godoc.org/github.com/cognusion/godoc2md)