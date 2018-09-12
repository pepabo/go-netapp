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
	AntivirusOnAccessPolicy    string   `xml:"antivirus-on-access-policy,omitempty"`
	AggregateList              []string `xml:"aggr-list>aggr-name"`
	Comment                    string   `xml:"comment,omitempty"`
	Ipspace                    string   `xml:"ipspace,omitempty"`
	IsRepositoryVserver        string   `xml:"is-repository-vserver,omitempty"`
	SnapshotPolicy             string   `xml:"snapshot-policy,omitempty"`
	UUID                       string   `xml:"uuid,omitempty"`
	VserverName                string   `xml:"vserver-name,omitempty"`
	VserverType                string   `xml:"vserver-type,omitempty"`
	AllowedProtocols           []string `xml:"allowed-protocols>protocol,omitempty"`
	DisallowedProtocols        []string `xml:"disallowed-protocols>protocol,omitempty"`
	IsConfigLockedForChanges   string   `xml:"is-config-locked-for-changes,omitempty"`
	Language                   string   `xml:"language,omitempty"`
	MaxVolumes                 string   `xml:"max-volumes,omitempty"`
	NameMappingSwitch          []string `xml:"name-mapping-switch>nmswitch,omitempty"`
	NameServerSwitch           []string `xml:"name-server-switch>nsswitch,omitempty"`
	OperationalState           string   `xml:"operational-state,omitempty"`
	QuotaPolicy                string   `xml:"quota-policy,omitempty"`
	RootVolume                 string   `xml:"root-volume,omitempty"`
	RootVolumeAggregate        string   `xml:"root-volume-aggregate,omitempty"`
	RootVolumeSecurityStyle    string   `xml:"root-volume-security-style,omitempty"`
	State                      string   `xml:"state,omitempty"`
	VolumeDeleteRetentionHours int      `xml:"volume-delete-retention-hours,omitempty"`
	VserverSubtype             string   `xml:"vserver-subtype,omitempty"`
}

type VServerQuery struct {
	VServerInfo *VServerInfo `xml:"vserver-info,omitempty"`
}
type VServerOptions struct {
	DesiredAttributes *VServerQuery `xml:"desired-attributes,omitempty"`
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
		SingleResultBase
		VServerInfo VServerInfo `xml:"attributes>vserver-info"`
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