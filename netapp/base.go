package netapp

import (
	"encoding/xml"
	"net/http"
)

type Base struct {
	XMLName xml.Name `xml:"netapp"`
	Version string   `xml:"version,attr"`
	XMLNs   string   `xml:"xmlsns,attr"`
	Name    string   `xml:"vfiler,attr,omitempty"`
	client  *Client
}

type ResultBase struct {
	Status     string `xml:"status,attr"`
	Reason     string `xml:"reason,attr"`
	NumRecords int    `xml:"num-records"`
	ErrorNo    int    `xml:"errno,attr"`
}

type SingleResultBase struct {
	Status  string `xml:"status,attr"`
	Reason  string `xml:"reason,attr"`
	ErrorNo int    `xml:"errno,attr"`
}

type AsyncResultBase struct {
	ErrorCode    int    `xml:"result-error-code"`
	ErrorMessage string `xml:"result-error-message"`
	JobID        int    `xml:"result-jobid"`
	JobStatus    string `xml:"result-status"`
	Status       string `xml:"status,attr"`
}

func (r *ResultBase) Passed() bool {
	return r.Status == "passed"
}

func (r *SingleResultBase) Passed() bool {
	return r.Status == "passed"
}

func (r *AsyncResultBase) Passed() bool {
	return r.Status == "passed"
}

func (b *Base) get(base interface{}, r interface{}) (*http.Response, error) {
	req, err := b.client.NewRequest("POST", &base)
	if err != nil {
		return nil, err
	}

	res, err := b.client.Do(req, r)
	if err != nil {
		return nil, err
	}

	return res, nil
}
