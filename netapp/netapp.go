package netapp

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	libraryVersion = "1"
	ServerURL      = `/servlets/netapp.servlets.admin.XMLrequest_filer`
	userAgent      = "go-netapp/" + libraryVersion
	XMLNs          = "http://www.netapp.com/filer/admin"
)

// A Client manages communication with the GitHub API.
type Client struct {
	client          *http.Client
	BaseURL         *url.URL
	UserAgent       string
	options         *ClientOptions
	Aggregate       *Aggregate
	AggregateSpace  *AggregateSpace
	AggregateSpares *AggregateSpares
	Cf              *Cf
	Diagnosis       *Diagnosis
	Fcp             *Fcp
	Fcport          *Fcport
	Job             *Job
	Lun             *Lun
	Net             *Net
	Qtree           *Qtree
	Quota           *Quota
	QuotaReport     *QuotaReport
	QuotaStatus     *QuotaStatus
	Snapshot        *Snapshot
	StorageDisk     *StorageDisk
	System          *System
	Volume          *Volume
	VolumeSpace     *VolumeSpace
	VServer         *VServer
}

type ClientOptions struct {
	BasicAuthUser     string
	BasicAuthPassword string
	SSLVerify         bool
	Timeout           time.Duration
}

func DefaultOptions() *ClientOptions {
	return &ClientOptions{
		SSLVerify: true,
	}
}

func NewClient(endpoint string, version string, options *ClientOptions) *Client {
	if options == nil {
		options = DefaultOptions()
	}

	httpClient := &http.Client{
		Timeout: options.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !options.SSLVerify,
			},
		},
	}
	if !strings.HasSuffix(endpoint, "/") {
		endpoint = endpoint + "/"
	}

	baseURL, _ := url.Parse(endpoint)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
		options:   options,
	}

	b := Base{
		client:  c,
		XMLNs:   XMLNs,
		Version: version,
	}

	c.Aggregate = &Aggregate{
		Base: b,
	}

	c.AggregateSpace = &AggregateSpace{
		Base: b,
	}

	c.AggregateSpares = &AggregateSpares{
		Base: b,
	}

	c.Cf = &Cf{
		Base: b,
	}

	c.Diagnosis = &Diagnosis{
		Base: b,
	}

	c.Fcp = &Fcp{
		Base: b,
	}

	c.Fcport = &Fcport{
		Base: b,
	}

	c.Job = &Job{
		Base: b,
	}

	c.Lun = &Lun{
		Base: b,
	}

	c.Net = &Net{
		Base: b,
	}

	c.Qtree = &Qtree{
		Base: b,
	}

	c.Quota = &Quota{
		Base: b,
	}

	c.QuotaReport = &QuotaReport{
		Base: b,
	}

	c.QuotaStatus = &QuotaStatus{
		Base: b,
	}

	c.Snapshot = &Snapshot{
		Base: b,
	}

	c.StorageDisk = &StorageDisk{
		Base: b,
	}

	c.System = &System{
		Base: b,
	}

	c.Volume = &Volume{
		Base: b,
	}

	c.VolumeSpace = &VolumeSpace{
		Base: b,
	}

	c.VServer = &VServer{
		Base: b,
	}

	return c
}

func (c *Client) NewRequest(method string, body interface{}) (*http.Request, error) {
	u, _ := c.BaseURL.Parse(ServerURL)

	buf, err := xml.MarshalIndent(body, "", "  ")
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Printf("request xml =====================================\n%v\n=================================================\n", string(buf))
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "text/xml")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	if c.options.BasicAuthUser != "" && c.options.BasicAuthPassword != "" {
		req.SetBasicAuth(c.options.BasicAuthUser, c.options.BasicAuthPassword)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") != "" {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bs))
		fmt.Printf("response xml ====================================\n%v\n=================================================\n", string(bs))
	}
	if v != nil {
		defer resp.Body.Close()
		err = xml.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}
