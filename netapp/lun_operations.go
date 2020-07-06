package netapp

import (
	"encoding/xml"
	"net/http"
)

// LUN Operations

const (
	LunOfflineOperation = "lun-offline"
	LunDestroyOperation = "lun-destroy"
)

type LunOperation struct {
	Base
	Params struct {
		XMLName xml.Name
		LunPath *lunPath
	}
}

type lunPath struct {
	XMLName xml.Name
	Path    string `xml:",innerxml"`
}

func (l LunOperation) Operation(vserverName string, lunPathName string, operation string) (*SingleResultResponse, *http.Response, error) {
	l.Params.XMLName = xml.Name{Local: operation}
	l.Name = vserverName
	elementName := "name"
	l.Params.LunPath = &lunPath{
		XMLName: xml.Name{Local: elementName},
		Path:    lunPathName,
	}
	r := SingleResultResponse{}
	res, err := l.get(l, &r)
	return &r, res, err
}
