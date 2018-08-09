package netapp

import (
	"encoding/xml"
	"net/http"
)

type Net struct {
	Base
}

type NetPortQuery struct {
	NetPortInfo *NetPortInfo `xml:"net-port-info,omitempty"`
}

type NetPortOptions struct {
	DesiredAttributes *NetPortQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int           `xml:"max-records,omitempty"`
	Query             *NetPortQuery `xml:"query,omitempty"`
	Tag               string        `xml:"tag,omitempty"`
}

type NetPortInfo struct {
	AdministrativeDuplex          string `xml:"administrative-duplex"`
	AdministrativeFlowcontrol     string `xml:"administrative-flowcontrol"`
	AdministrativeSpeed           string `xml:"administrative-speed"`
	AutorevertDelay               int    `xml:"autorevert-delay"`
	IfgrpDistributionFunction     string `xml:"ifgrp-distribution-function"`
	IfgrpMode                     string `xml:"ifgrp-mode"`
	IfgrpNode                     string `xml:"ifgrp-node"`
	IfgrpPort                     string `xml:"ifgrp-port"`
	IsAdministrativeAutoNegotiate bool   `xml:"is-administrative-auto-negotiate"`
	IsAdministrativeUp            bool   `xml:"is-administrative-up"`
	IsOperationalAutoNegotiate    bool   `xml:"is-operational-auto-negotiate"`
	LinkStatus                    string `xml:"link-status"`
	MacAddress                    string `xml:"mac-address"`
	Mtu                           int    `xml:"mtu"`
	Node                          string `xml:"node"`
	OperationalDuplex             string `xml:"operational-duplex"`
	OperationalFlowcontrol        string `xml:"operational-flowcontrol"`
	OperationalSpeed              string `xml:"operational-speed"`
	Port                          string `xml:"port"`
	PortType                      string `xml:"port-type"`
	RemoteDeviceId                string `xml:"remote-device-id"`
	Role                          string `xml:"role"`
	VlanId                        int    `xml:"vlan-id"`
	VlanNode                      string `xml:"vlan-node"`
	VlanPort                      string `xml:"vlan-port"`
}

type NetPortGetIterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			NetPortAttributes []NetPortInfo `xml:"net-port-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

type NetPortPageResponse struct {
	Response    *NetPortGetIterResponse
	Error       error
	RawResponse *http.Response
}

type NetPortPageHandler func(NetPortPageResponse) (shouldContinue bool)

func (n *Net) NetPortGetIter(options *NetPortOptions) (*NetPortGetIterResponse, *http.Response, error) {
	params := newNetPortGetIterParams(options, n.Base)
	r := NetPortGetIterResponse{}
	res, err := n.get(params, &r)
	return &r, res, err
}

func (n *Net) NetPortGetAll(options *NetPortOptions, fn NetPortPageHandler) {

	requestOptions := options

	for shouldContinue := true; shouldContinue; {
		netPortGetIterResponse, res, err := n.NetPortGetIter(requestOptions)
		handlerResponse := false

		handlerResponse = fn(NetPortPageResponse{Response: netPortGetIterResponse, Error: err, RawResponse: res})

		nextTag := ""
		if err == nil {
			nextTag = netPortGetIterResponse.Results.NextTag
			requestOptions = &NetPortOptions{
				Tag:        nextTag,
				MaxRecords: options.MaxRecords,
			}
		}
		shouldContinue = nextTag != "" && handlerResponse
	}

}

type netPortGetIterParams struct {
	Base
	Params struct {
		XMLName xml.Name
		*NetPortOptions
	}
}

func newNetPortGetIterParams(options *NetPortOptions, base Base) *netPortGetIterParams {
	params := netPortGetIterParams{
		Base: base,
	}
	params.Params.XMLName = xml.Name{Local: "net-port-get-iter"}
	params.Params.NetPortOptions = options
	return &params
}

type NetInterfaceQuery struct {
	NetInterfaceInfo *NetInterfaceInfo `xml:"net-interface-info,omitempty"`
}

type NetInterfaceOptions struct {
	DesiredAttributes *NetInterfaceQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                `xml:"max-records,omitempty"`
	Query             *NetInterfaceQuery `xml:"query,omitempty"`
	Tag               string             `xml:"tag,omitempty"`
}

type NetInterfaceInfo struct {
	Address              string `xml:"address"`
	AdministrativeStatus string `xml:"administrative-status"`
	Comment              string `xml:"comment"`
	CurrentNode          string `xml:"current-node"`
	CurrentPort          string `xml:"current-port"`
	DnsDomainName        string `xml:"dns-domain-name"`
	FailoverGroup        string `xml:"failover-group"`
	FailoverPolicy       string `xml:"failover-policy"`
	FirewallPolicy       string `xml:"firewall-policy"`
	HomeNode             string `xml:"home-node"`
	HomePort             string `xml:"home-port"`
	InterfaceName        string `xml:"interface-name"`
	IsAutoRevert         bool   `xml:"is-auto-revert"`
	IsHome               bool   `xml:"is-home"`
	IsIpv4LinkLocal      bool   `xml:"is-ipv4-link-local"`
	Netmask              string `xml:"netmask"`
	NetmaskLength        int    `xml:"netmask-length"`
	OperationalStatus    string `xml:"operational-status"`
	Role                 string `xml:"role"`
	RoutingGroupName     string `xml:"routing-group-name"`
	UseFailoverGroup     string `xml:"use-failover-group"`
	Vserver              string `xml:"vserver"`
}

type NetInterfaceGetIterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			NetInterfaceAttributes []NetInterfaceInfo `xml:"net-interface-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

type NetInterfacePageResponse struct {
	Response    *NetInterfaceGetIterResponse
	Error       error
	RawResponse *http.Response
}

type NetInterfacePageHandler func(NetInterfacePageResponse) (shouldContinue bool)

func (n *Net) NetInterfaceGetIter(options *NetInterfaceOptions) (*NetInterfaceGetIterResponse, *http.Response, error) {
	params := newNetInterfaceGetIterParams(options, n.Base)
	r := NetInterfaceGetIterResponse{}
	res, err := n.get(params, &r)
	return &r, res, err
}

func (n *Net) NetInterfaceGetAll(options *NetInterfaceOptions, fn NetInterfacePageHandler) {

	requestOptions := options

	for shouldContinue := true; shouldContinue; {
		netInterfaceGetIterResponse, res, err := n.NetInterfaceGetIter(requestOptions)
		handlerResponse := false

		handlerResponse = fn(NetInterfacePageResponse{Response: netInterfaceGetIterResponse, Error: err, RawResponse: res})

		nextTag := ""
		if err == nil {
			nextTag = netInterfaceGetIterResponse.Results.NextTag
			requestOptions = &NetInterfaceOptions{
				Tag:        nextTag,
				MaxRecords: options.MaxRecords,
			}
		}
		shouldContinue = nextTag != "" && handlerResponse
	}

}

type netInterfaceGetIterParams struct {
	Base
	Params struct {
		XMLName xml.Name
		*NetInterfaceOptions
	}
}

func newNetInterfaceGetIterParams(options *NetInterfaceOptions, base Base) *netInterfaceGetIterParams {
	params := netInterfaceGetIterParams{
		Base: base,
	}
	params.Params.XMLName = xml.Name{Local: "net-interface-get-iter"}
	params.Params.NetInterfaceOptions = options
	return &params
}
