package netapp_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	b, err := ioutil.ReadFile("fixtures/" + path)
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
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseFixture)
	})

	c = netapp.NewClient(baseURL, "1.10", nil)

	return c, teardown
}
