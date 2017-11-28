package netapp

import (
	"encoding/xml"
	"net/http"
)

type VServer struct {
	Base
	Params struct {
		XMLName xml.Name
		*VServerOptions
	}
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

type VServerQuery struct {
	VServerInfo *VServerInfo `xml:"vserver-info,omitempty"`
}
type VServerOptions struct {
	DesiredAttributes *VServerInfo  `xml:"desired-attributes,omitempty"`
	MaxRecords        int           `xml:"max-records,omitempty"`
	Query             *VServerQuery `xml:"query,omitempty"`
	Tag               string        `xml:"tag,omitempty"`
}

type VServerListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			VserverInfo []VServerInfo `xml:"vserver-info"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

type VServerResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			VserverInfo []VServerInfo `xml:"vserver-info"`
		} `xml:"attributes"`
	} `xml:"results"`
}

func (v *VServer) Get(name string, options *VServerOptions) (*VServerResponse, *http.Response, error) {
	v.Name = name
	v.Params.XMLName = xml.Name{Local: "vserver-get"}
	v.Params.VServerOptions = options
	r := VServerResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}

func (v *VServer) List(options *VServerOptions) (*VServerListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "vserver-get-iter"}
	v.Params.VServerOptions = options

	r := VServerListResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}
