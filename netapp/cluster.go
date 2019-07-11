package netapp

import (
	"encoding/xml"
	"net/http"
)

type Cluster struct {
	Base
	Params struct {
		XMLName xml.Name
		*ClusterOptions
	}
}



type ClusterInfo struct {
	ClusterContact          string `xml:"cluster-contact,omitempty"`
	ClusterLocation         string `xml:"cluster-location"`
	ClusterName  					  int    `xml:"cluster-name"`
	ClusterSerialNumber 	  string `xml:"cluster-serial-number"`
	RdbUuid                 string `xml:"rdb-uuid"`
	UUID                    string `xml:"uuid"`
}




type  ClusterOptions struct {
	DesiredAttributes *ClusterInfo `xml:"desired-attributes,omitempty"`
}

type  ClusterResponse struct {
	DesiredAttributes *ClusterInfo `xml:"desired-attributes,omitempty"`
}


func (c *Cluster) List(options *ClusterOptions) (*ClusterResponse, *http.Response, error) {
	c.Params.XMLName = xml.Name{Local: "cluster-identity-get"}
	c.Params.ClusterOptions = options
	r := ClusterResponse{}
	res, err := c.get(c, &r)
	return &r, res, err
}
