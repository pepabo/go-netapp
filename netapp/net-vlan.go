package netapp

import (
	"encoding/xml"
	"net/http"
)

type netVlanRequest struct {
	Base
	Params struct {
		XMLName     xml.Name
		NetVlanInfo `xml:",innerxml"`
		VlanInfo    *NetVlanInfo `xml:"vlan-info,omitempty"`
	}
}

// NetVlanOptions get/list options for getting broadcast domains
type NetVlanOptions struct {
	DesiredAttributes *NetVlanInfo `xml:"desired-attributes,omitempty"`
	MaxRecords        int          `xml:"max-records,omitempty"`
	Query             *NetVlanInfo `xml:"query,omitempty"`
	Tag               string       `xml:"tag,omitempty"`
}

// NetVlanInfo is the Broadcast Domain data
type NetVlanInfo struct {
	InterfaceName   string `xml:"interface-name,omitempty"`
	Node            string `xml:"node,omitempty"`
	ParentInterface string `xml:"parent-interface,omitempty"`
	VlanID          int    `xml:"vlanid,omitempty"`
}

// NetVlanResponse returns results for broadcast domains
type NetVlanResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		SingleResultBase
		Info NetVlanInfo `xml:"attributes>vlan-info"`
	} `xml:"results"`
}

// CreateVlan creates a new broadcast domain
func (n Net) CreateVlan(options *NetVlanInfo) (*SingleResultResponse, *http.Response, error) {
	req := n.newNetVlanRequest()
	req.Params.XMLName = xml.Name{Local: "net-vlan-create"}

	options.InterfaceName = ""
	req.Params.VlanInfo = options

	r := SingleResultResponse{}
	res, err := n.get(req, &r)
	return &r, res, err
}

// GetVlan grabs a single named broadcast domain
func (n Net) GetVlan(interfaceName string, node string) (*NetVlanResponse, *http.Response, error) {
	req := n.newNetVlanRequest()
	req.Params.XMLName = xml.Name{Local: "net-vlan-get"}
	req.Params.InterfaceName = interfaceName
	req.Params.Node = node
	r := NetVlanResponse{}
	res, err := n.get(req, &r)
	return &r, res, err
}

// DeleteVlan removes vlan from existence
func (n Net) DeleteVlan(options *NetVlanInfo) (*SingleResultResponse, *http.Response, error) {
	req := n.newNetVlanRequest()
	req.Params.XMLName = xml.Name{Local: "net-vlan-delete"}
	options.InterfaceName = ""
	req.Params.VlanInfo = options
	r := SingleResultResponse{}
	res, err := n.get(req, &r)
	return &r, res, err
}

func (n Net) newNetVlanRequest() *netVlanRequest {
	return &netVlanRequest{
		Base: n.Base,
	}
}
