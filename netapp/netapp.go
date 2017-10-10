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
)

const (
	libraryVersion = "1"
	ServerURL      = `/servlets/netapp.servlets.admin.XMLrequest_filer`
	userAgent      = "go-netapp/" + libraryVersion
	XMLNs          = "http://www.netapp.com/filer/admin"
)

// A Client manages communication with the GitHub API.
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	options   *ClientOptions
	VServer   *VServer
	Quota     *Quota
	Job       *Job
	Qtree     *Qtree
}

type ClientOptions struct {
	BasicAuthUser     string
	BasicAuthPassword string
	SSLVerify         bool
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
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !options.SSLVerify,
			},
		},
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

	c.VServer = &VServer{
		Base: b,
	}
	c.Quota = &Quota{
		Base: b,
	}

	c.Job = &Job{
		Base: b,
	}

	c.Qtree = &Qtree{
		Base: b,
	}

	return c
}

func (c *Client) NewRequest(method string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
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
