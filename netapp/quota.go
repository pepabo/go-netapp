package netapp

import (
	"encoding/xml"
	"net/http"
)

const (
	QuotaStatusCorrupt      = "corrupt"
	QuotaStatusInitializing = "initializing"
	QuotaStatusMixed        = "mixed"
	QuotaStatusOff          = "off"
	QuotaStatusOn           = "on"
	QuotaStatusResizing     = "resizing"
	QuotaStatusReverting    = "reverting"
	QuotaStatusUnknown      = "unknown"
	QuotaStatusUpgrading    = "upgrading"
)

type Quota struct {
	Base
	Params struct {
		XMLName xml.Name
		*QuotaOptions
	}
}

type QuotaQuery struct {
	QuotaEntry *QuotaEntry `xml:"quota-entry,omitempty"`
}

type QuotaOptions struct {
	DesiredAttributes *QuotaEntry `xml:"desired-attributes,omitempty"`
	MaxRecords        int         `xml:"max-records,omitempty"`
	Query             *QuotaQuery `xml:"query,omitempty"`
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
		ResultBase
		QuotaEntry
	} `xml:"results"`
}

type QuotaListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			QuotaEntry []QuotaEntry `xml:"quota-entry"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

type QuotaStatusResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		QuotaStatus    string `xml:"status"`
		QuotaSubStatus string `xml:"substatus"`
		ResultJobid    string `xml:"result-jobid"`
		ResultStatus   string `xml:"result-status"`
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

func (q *Quota) Create(serverName, target, quotaType, qtree string, entry *QuotaEntry) (*QuotaListResponse, *http.Response, error) {
	q.Name = serverName
	q.Params.XMLName = xml.Name{Local: "quota-add-entry"}

	if entry == nil {
		entry = &QuotaEntry{}
	}

	entry.QuotaTarget = target
	entry.QuotaType = quotaType
	entry.Qtree = &qtree

	q.Params.QuotaOptions = &QuotaOptions{
		QuotaEntry: entry,
	}

	r := QuotaListResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}

func (q *Quota) Update(serverName string, entry *QuotaEntry) (*QuotaListResponse, *http.Response, error) {
	q.Name = serverName
	q.Params.XMLName = xml.Name{Local: "quota-modify-entry"}
	q.Params.QuotaOptions = &QuotaOptions{
		QuotaEntry: entry,
	}

	r := QuotaListResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}

func (q *Quota) Delete(serverName, target, quotaType, volume, qtree string) (*QuotaListResponse, *http.Response, error) {
	q.Name = serverName
	q.Params.XMLName = xml.Name{Local: "quota-delete-entry"}
	q.Params.QuotaOptions = &QuotaOptions{
		QuotaEntry: &QuotaEntry{
			QuotaType:   quotaType,
			QuotaTarget: target,
			Volume:      volume,
			Qtree:       &qtree,
		},
	}

	r := QuotaListResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}

func (q *Quota) Off(serverName, volumeName string) (*QuotaStatusResponse, *http.Response, error) {
	q.Name = serverName
	q.Params.XMLName = xml.Name{Local: "quota-off"}
	q.Params.QuotaOptions = &QuotaOptions{
		QuotaEntry: &QuotaEntry{
			Volume: volumeName,
		},
	}

	r := QuotaStatusResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}

func (q *Quota) On(serverName, volumeName string) (*QuotaStatusResponse, *http.Response, error) {
	q.Name = serverName
	q.Params.XMLName = xml.Name{Local: "quota-on"}
	q.Params.QuotaOptions = &QuotaOptions{
		QuotaEntry: &QuotaEntry{
			Volume: volumeName,
		},
	}

	r := QuotaStatusResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}

func (q *Quota) Status(serverName, volumeName string) (*QuotaStatusResponse, *http.Response, error) {
	q.Name = serverName
	q.Params.XMLName = xml.Name{Local: "quota-status"}
	q.Params.QuotaOptions = &QuotaOptions{
		QuotaEntry: &QuotaEntry{
			Volume: volumeName,
		},
	}

	r := QuotaStatusResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}
