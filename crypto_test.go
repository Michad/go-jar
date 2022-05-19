package jar

import (
	. "github.com/smartystreets/goconvey/convey"

	//"log"
	//"os"
	//"encoding/base64"
	"testing"
)

func init() {
	//DebugOut = log.New(os.Stderr, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func TestMd5Compare(t *testing.T) {

	hpassword := "$1$150542$R/qCXbVm8MWwLwAu20dRA."
	password := "CVFDSQAYTRVCGDCNGHBG"

	Convey("When hash is is compared against the known source, it is treated correctly", t, func() {
		So(compareMD5HashAndPassword([]byte(hpassword), []byte(password)), ShouldBeNil)
	})
}

func TestMd5CompareLong(t *testing.T) {

	hpassword := "$1$150542$R/qCXbVm8MWwLwAu20dRA."
	password := "CVFDSQAYTRVCGDCNGHBGISWAYTOOLONG"

	Convey("When hash is is compared against the known source extended past the usual characters, it is treated correctly", t, func() {
		So(compareMD5HashAndPassword([]byte(hpassword), []byte(password)), ShouldNotBeNil)
	})
}

func TestMd5CompareFail(t *testing.T) {

	hpassword := "$1$150542$R/qCXbVm8MWwLwAu20dRA."
	password := "CVFDSQAYTRVCGDCNGHBH"

	Convey("When hash is is compared against the known source, it is treated correctly", t, func() {
		So(compareMD5HashAndPassword([]byte(hpassword), []byte(password)), ShouldNotBeNil)
	})
}

func TestShaCompare(t *testing.T) {

	hpassword := "{SHA}ZyY/sOKNPWfFX8PnaapCyCnPtA0="
	password := "CVFDSQAYTRVCGDCNGHBH"

	Convey("When hash is is compared against the known source, it is treated correctly", t, func() {
		So(compareShaHashAndPassword([]byte(hpassword), []byte(password)), ShouldBeNil)
	})

}

func TestShaCompareFail(t *testing.T) {

	hpassword := "{SHA}ZyY/sOKNPWfFX8PnaapCyCnPtA0="
	password := "CVFDSQAYTRVCGDCNGHB"

	Convey("When hash is is compared against the known source, it is treated correctly", t, func() {
		So(compareShaHashAndPassword([]byte(hpassword), []byte(password)), ShouldNotBeNil)
	})
}

func TestECBDecrypt(t *testing.T) {

	key := "ZWYxU0FPd2xFSmFmOVhaQg=="
	token := "h4wz0hWRBT-G03bXzoqeN_PYFmH5r_LJRKBFScaLaIj1ANCgjWl_ZyUDyf6_ohIBzT0mUdzJDuvVCR1Wi-UvtTJoMbgdvKoknjjbcTvOnbtp8exWQ4gadtBcKQgvHREytibUlGU1hINLKPF2zkclXitYMnKHFav38kg23htuOkU"
	clear := "catnetglobalcoastalflooduw_2017_v1_demo|catnet_coastal_flood_combined|112ca1af-bfac-44e3-9354-5160ab7af820|1554304914"

	Convey("When a token is decrypted with a key, the clear token is as expected", t, func() {
		cleartoken, err := ECBDecrypt(key, token)
		So(err, ShouldBeNil)
		So(string(cleartoken), ShouldEqual, clear)
	})

}

func TestECBEncrypt(t *testing.T) {

	key := "ZWYxU0FPd2xFSmFmOVhaQg=="
	token := "h4wz0hWRBT-G03bXzoqeN_PYFmH5r_LJRKBFScaLaIj1ANCgjWl_ZyUDyf6_ohIBzT0mUdzJDuvVCR1Wi-UvtTJoMbgdvKoknjjbcTvOnbtp8exWQ4gadtBcKQgvHREytibUlGU1hINLKPF2zkclXitYMnKHFav38kg23htuOkU"
	clear := "catnetglobalcoastalflooduw_2017_v1_demo|catnet_coastal_flood_combined|112ca1af-bfac-44e3-9354-5160ab7af820|1554304914"

	Convey("When a token is encrypted with a key, the clear token is as expected", t, func() {
		cryptoken, err := ECBEncrypt(key, []byte(clear))
		So(err, ShouldBeNil)
		So(cryptoken, ShouldEqual, token)
	})

}
