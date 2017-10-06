package netapp

import (
	"encoding/xml"
	"net/http"
)

type Quota struct {
	Base
	Params struct {
		XMLName xml.Name
		*QuotaOptions
	}
}

type QuotaOptions struct {
	DesiredAttributes *QuotaEntry `xml:"desired-attributes,omitempty"`
	MaxRecords        int         `xml:"max-records,omitempty"`
	Query             *QuotaEntry `xml:"query,omitempty"`
	Tag               string      `xml:"tag,omitempty"`
	*QuotaEntry
}

type QuotaEntry struct {
	DiskLimit          string  `xml:"disk-limit,omitempty"`
	FileLimit          string  `xml:"file-limit,omitempty"`
	PerformUserMapping string  `xml:"perform-user-mapping,omitempty"`
	Policy             string  `xml:"policy,omitempty"`
	Qtree              *string `xml:"qtree,omitempty"`
	QuotaTarget        string  `xml:"quota-target,omitempty"`
	QuotaType          string  `xml:"quota-type,omitempty"`
	SoftDiskLimit      string  `xml:"soft-disk-limit,omitempty"`
	SoftFileLimit      string  `xml:"soft-file-limit,omitempty"`
	Threshold          string  `xml:"threshold,omitempty"`
	Volume             string  `xml:"volume,omitempty"`
	Vserver            string  `xml:"vserver,omitempty"`
}

type QuotaResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		Status string `xml:"status,attr"`
		QuotaEntry
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

type QuotaListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		Status         string `xml:"status,attr"`
		AttributesList struct {
			QuotaEntry QuotaEntry `xml:"quota-entry"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

func (q *Quota) Get(name string, options *QuotaOptions) (*QuotaResponse, *http.Response, error) {
	q.Name = name
	q.Params.XMLName = xml.Name{Local: "quota-get-entry"}
	q.Params.QuotaOptions = options
	r := QuotaResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}

func (q *Quota) List(options *QuotaOptions) (*QuotaListResponse, *http.Response, error) {
	q.Params.XMLName = xml.Name{Local: "quota-list-entries-iter"}
	q.Params.QuotaOptions = options

	r := QuotaListResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}
