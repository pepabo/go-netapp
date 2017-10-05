package netapp

import (
	"encoding/xml"
	"net/http"
)

type VServer struct {
	XMLName     xml.Name `xml:"netapp"`
	Version     string   `xml:"version,attr"`
	XMLNs       string   `xml:"xmlsns,attr"`
	VServerName string   `xml:"vfiler,attr"`

	Params struct {
		XMLName xml.Name
		*VServerOptions
	}
	client *Client
}

type VServerInfo struct {
	AntivirusOnAccessPolicy string `xml:"antivirus-on-access-policy"`
	Comment                 string `xml:"comment"`
	Ipspace                 string `xml:"ipspace,omitempty"`
	IsRepositoryVserver     string `xml:"is-repository-vserver"`
	SnapshotPolicy          string `xml:"snapshot-policy,omitempty"`
	UUID                    string `xml:"uuid"`
	VserverName             string `xml:"vserver-name"`
	VserverType             string `xml:"vserver-type"`
	AllowedProtocols        struct {
		Protocol []string `xml:"protocol"`
	} `xml:"allowed-protocols,omitempty"`
	DisallowedProtocols struct {
		Protocol []string `xml:"protocol"`
	} `xml:"disallowed-protocols,omitempty"`
	IsConfigLockedForChanges string `xml:"is-config-locked-for-changes,omitempty"`
	Language                 string `xml:"language,omitempty"`
	MaxVolumes               string `xml:"max-volumes,omitempty"`
	NameMappingSwitch        struct {
		Nmswitch string `xml:"nmswitch"`
	} `xml:"name-mapping-switch,omitempty"`
	NameServerSwitch struct {
		Nsswitch string `xml:"nsswitch"`
	} `xml:"name-server-switch,omitempty"`
	OperationalState           string `xml:"operational-state,omitempty"`
	QuotaPolicy                string `xml:"quota-policy,omitempty"`
	RootVolume                 string `xml:"root-volume,omitempty"`
	RootVolumeAggregate        string `xml:"root-volume-aggregate,omitempty"`
	RootVolumeSecurityStyle    string `xml:"root-volume-security-style,omitempty"`
	State                      string `xml:"state,omitempty"`
	VolumeDeleteRetentionHours string `xml:"volume-delete-retention-hours,omitempty"`
	VserverSubtype             string `xml:"vserver-subtype,omitempty"`
}

type VServerOptions struct {
	DesiredAttributes *VServerInfo `xml:"desired-attributes,omitempty"`
	MaxRecords        int          `xml:"max-records,omitempty"`
	Query             *VServerInfo `xml:"query,omitempty"`
	Tag               string       `xml:"tag,omitempty"`
}

type VServerListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		Status         string `xml:"status,attr"`
		AttributesList struct {
			VserverInfo []VServerInfo `xml:"vserver-info"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

type VServerResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		Status         string `xml:"status,attr"`
		AttributesList struct {
			VserverInfo []VServerInfo `xml:"vserver-info"`
		} `xml:"attributes"`
	} `xml:"results"`
}

func (v *VServer) get(options *VServerOptions, vr interface{}) (*http.Response, error) {
	v.Version = v.client.Version
	v.XMLNs = v.client.XMLNs
	v.Params.VServerOptions = options
	req, err := v.client.NewRequest("POST", v)
	if err != nil {
		return nil, err
	}

	res, err := v.client.Do(req, vr)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (v *VServer) Get(name string, options *VServerOptions) (*VServerResponse, *http.Response, error) {
	v.VServerName = name
	v.Params.XMLName = xml.Name{Local: "vserver-get"}

	r := VServerResponse{}
	res, err := v.get(options, &r)
	return &r, res, err
}

func (v *VServer) List(options *VServerOptions) (*VServerListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "vserver-get-iter"}

	r := VServerListResponse{}
	res, err := v.get(options, &r)
	return &r, res, err
}
