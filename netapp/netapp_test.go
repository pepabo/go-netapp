package netapp_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/pepabo/go-netapp/netapp"
)

func setup() (baseURL string, mux *http.ServeMux, teardownFn func()) {
	mux = http.NewServeMux()
	srv := httptest.NewServer(mux)
	return srv.URL, mux, srv.Close
}

func fixture(path string, t *testing.T) []byte {
	r, err := os.OpenFile("fixtures/"+path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func createTestClientWithFixtures(t *testing.T) (c *netapp.Client, teardownFn func()) {
	baseURL, mux, teardown := setup()

	requestFixture := bytes.TrimSpace(fixture(fmt.Sprintf("%s_%s", t.Name(), "request.xml"), t))
	responseFixture := bytes.TrimSpace(fixture(fmt.Sprintf("%s_%s", t.Name(), "response.xml"), t))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		bs, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("Got error reading body %s", err)
		}

		if !bytes.Equal(bs, requestFixture) {
			t.Errorf("%s: result not as expected:\n%v", t.Name(), diff.LineDiff(string(requestFixture), string(bs)))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseFixture)
	})

	c, _ = netapp.NewClient(baseURL, "1.10", nil)
	log.SetOutput(ioutil.Discard)
	return c, teardown
}

func checkResponseSuccess(result netapp.Result, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Should not have gotten an error: '%s'", err)
	}

	if !result.Passed() {
		t.Fatalf("Got the failure response, expected success, reason: %s", result.Result().Reason)
	}
}

func checkResponseFailure(result netapp.Result, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Should not have gotten an error %s", err)
	}

	if result.Passed() {
		t.Fatal("Got the successful response, expecting failure")
	}
}

func testFailureResult(errorNo int, expectedReason string, r netapp.Result, t *testing.T) {
	result := r.Result()
	if result.ErrorNo != errorNo {
		t.Errorf("%s got = %+v, want %+v", t.Name(), result.ErrorNo, errorNo)
	}

	if result.Reason != expectedReason {
		t.Errorf("%s got = %+v, want %+v", t.Name(), result.Reason, expectedReason)
	}
}

// debugTableItems is used to get reflected vaules of 2 items so its easier to tell why reflect.DeepEqual() fails
func debugItems(v1 interface{}, v2 interface{}) {
	val1 := reflect.ValueOf(v1)
	val2 := reflect.ValueOf(v2)
	fmt.Printf("v1: %v, v2: %v", val1, val2)
}

func TestClientCerts(t *testing.T) {
	c, _ := netapp.NewClient("", "1.10", &netapp.ClientOptions{
		CertFile: "test_cert.pem",
		KeyFile:  "test_key.pem",
	},
	)

	if c == nil {
		t.Error(`NewClient with certs failed`)
	}
}

func TestMissingClientCertArgs(t *testing.T) {
	c, _ := netapp.NewClient("", "1.10", &netapp.ClientOptions{
		CertFile: "test_cert.pem",
	},
	)

	if c != nil {
		t.Error(`NewClient with invalid client key arguments should fail`)
	}

	c, _ = netapp.NewClient("", "1.10", &netapp.ClientOptions{
		KeyFile: "test_key.pem",
	},
	)

	if c != nil {
		t.Error(`NewClient with invalid client cert argument should fail`)
	}
}
func TestBadClientCertArgs(t *testing.T) {
	c, _ := netapp.NewClient("", "1.10", &netapp.ClientOptions{
		CertFile: "test_certxxxxxxx.pem",
		KeyFile:  "test_cert.pem",
	},
	)

	if c != nil {
		t.Error(`NewClient with invalid client key arguments should fail`)
	}

	c, _ = netapp.NewClient("", "1.10", &netapp.ClientOptions{
		CertFile: "test_cert.pem",
		KeyFile:  "test_keyxxxxxx.pem",
	},
	)

	if c != nil {
		t.Error(`NewClient with invalid client cert argument should fail`)
	}

}
func TestCACert(t *testing.T) {
	c, _ := netapp.NewClient("", "1.10", &netapp.ClientOptions{
		CAFile: "test_cert.pem",
	},
	)

	if c == nil {
		t.Error(`NewClient with CAFile failed`)
	}
}
func TestBadCACert(t *testing.T) {
	c, _ := netapp.NewClient("", "1.10", &netapp.ClientOptions{
		CAFile: "test_certxxxxxxxxx.pem",
	},
	)

	if c != nil {
		t.Error(`NewClient with invalid CAFile arguments should fail`)
	}
}
