package netapp

import (
	"encoding/xml"
	"net/http"
)

type netIPSpaceRequest struct {
	Base
	CreateParams *netIPSpaceCreateParams `xml:"net-ipspaces-create,omitempty"`
	GetParams    *netIPSpaceGetParams    `xml:"net-ipspaces-get,omitempty"`
	RenameParams *netIPSpaceRenameParams `xml:"net-ipspaces-rename,omitempty"`
	DeleteParams *netIPSpaceGetParams    `xml:"net-ipspaces-destroy,omitempty"`
}

type netIPSpaceCreateParams struct {
	IPSpace      string `xml:"ipspace"`
	ReturnRecord bool   `xml:"return-record"`
}

type netIPSpaceGetParams struct {
	IPSpace string `xml:"ipspace"`
}

type netIPSpaceRenameParams struct {
	IPSpace string `xml:"ipspace"`
	NewName string `xml:"new-name"`
}

// NetIPSpaceInfo holds newly created ipspace variables
type NetIPSpaceInfo struct {
	BroadcastDomains []string `xml:"broadcast-domains>broadcast-domain-name,omitempty"`
	ID               int      `xml:"id"`
	IPSpace          string   `xml:"ipspace"`
	Ports            []string `xml:"ports>net-qualified-port-name,omitempty"`
	UUID             string   `xml:"uuid"`
	VServers         []string `xml:"vservers>vserver-name"`
}

// NetIPSpaceResponse is return type for net ip space requests
type NetIPSpaceResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		SingleResultBase
		NetIPSpaceInfo       `xml:",innerxml"`
		NetIPSpaceCreateInfo *NetIPSpaceInfo `xml:"result>net-ipspaces-info"`
	} `xml:"results"`
}

// CreateIPSpace creates a new ipspace on the cluster
func (n Net) CreateIPSpace(name string, returnRecord bool) (*NetIPSpaceResponse, *http.Response, error) {
	req := n.newNetIPSpaceRequest()
	req.CreateParams = &netIPSpaceCreateParams{
		IPSpace:      name,
		ReturnRecord: returnRecord,
	}
	return n.newNetIPSpaceResponse(req)
}

// GetIPSpace grabs data for an ip space
func (n Net) GetIPSpace(name string) (*NetIPSpaceResponse, *http.Response, error) {
	req := n.newNetIPSpaceRequest()
	req.GetParams = &netIPSpaceGetParams{
		IPSpace: name,
	}

	return n.newNetIPSpaceResponse(req)
}

// RenameIPSpace changes the name of an ipspace
func (n Net) RenameIPSpace(name string, newName string) (*NetIPSpaceResponse, *http.Response, error) {
	req := n.newNetIPSpaceRequest()
	req.RenameParams = &netIPSpaceRenameParams{
		IPSpace: name,
		NewName: newName,
	}

	return n.newNetIPSpaceResponse(req)
}

// DeleteIPSpace deletes an IPSpace
func (n Net) DeleteIPSpace(name string) (*NetIPSpaceResponse, *http.Response, error) {
	req := n.newNetIPSpaceRequest()
	req.DeleteParams = &netIPSpaceGetParams{
		IPSpace: name,
	}

	return n.newNetIPSpaceResponse(req)
}

func (n Net) newNetIPSpaceRequest() *netIPSpaceRequest {
	return &netIPSpaceRequest{
		Base: n.Base,
	}
}

func (n Net) newNetIPSpaceResponse(req *netIPSpaceRequest) (*NetIPSpaceResponse, *http.Response, error) {
	r := NetIPSpaceResponse{}
	res, err := n.get(req, &r)
	return &r, res, err
}
