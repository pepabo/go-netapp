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
	NumRecords string `json:"num-records"`
}

func (r *ResultBase) Passed() bool {
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
