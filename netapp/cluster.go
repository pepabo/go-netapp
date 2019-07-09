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
	ClusterContact          *ClusterContactInfo     `xml:"cluster-contact,omitempty"`
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
	c.Params.NodeDetailOptions = options
	r := ClusterResponse{}
	res, err := s.get(s, &r)
	return &r, res, err
}


type ClusterClusterContact struct {
	Base
	Params struct {
		XMLName xml.Name
		*ClusterContactOptions
	}
}


type ClusterContactInfo struct {
	address 				  string 
	business-name	 		string 
	city 						  string 
	country						string 
	primary-alt-phone	string 
	primary-email			string
	primary-name			string 
	primary-phone			string 
	secondary-alt-phone		string 
	secondary-email				string 
	secondary-name				string 	
	secondary-phone				string 	
	state									string 
	zip-code							string 
}



type  ClusterContactOptions struct {
	DesiredAttributes *ClusterContactInfo `xml:"desired-attributes,omitempty"`
}

type   ClusterContactResponse struct {
	DesiredAttributes *ClusterContactInfo `xml:"desired-attributes,omitempty"`
}


func (c *ClusterClusterContact) List(options *ClusterContactOptions) (*ClusterContactResponse, *http.Response, error) {
	c.Params.XMLName = xml.Name{Local: "cluster-contact-get"}
	c.Params.ClusterContactOptions = options
	r := ClusterResponse{}
	res, err := s.get(s, &r)
	return &r, res, err
}
