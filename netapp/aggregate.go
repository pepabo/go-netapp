package netapp

import (
	"encoding/xml"
	"net/http"
)

type AggregateSpace struct {
	Base
	Params struct {
		XMLName xml.Name
		*AggrSpaceOptions
	}
}

type AggrSpaceOptions struct {
	DesiredAttributes *AggrSpaceInfo `xml:"desired-attributes,omitempty"`
	MaxRecords        int            `xml:"max-records,omitempty"`
	Query             *AggrSpaceInfo `xml:"query,omitempty"`
	Tag               string         `xml:"tag,omitempty"`
}

type AggrSpaceInfo struct {
	Aggregate 							string 	`xml:"aggregate"`
	AggregateMetadata 					string 	`xml:"aggregate-metadata"`
	AggregateMetadataPercent 			string 	`xml:"aggregate-metadata-percent"`
	AggregateSize 						string 	`xml:"aggregate-size"`
	PercentSnapshotSpace 				string 	`xml:"percent-snapshot-space"`
	PhysicalUsed 						string 	`xml:"physical-used"`
	PhysicalUsedPercent 				string 	`xml:"physical-used-percent"`
	SnapSizeTotal 						string 	`xml:"snap-size-total"`
	SnapshotReserveUnusable 			string 	`xml:"snapshot-reserve-unusable"`
	SnapshotReserveUnusablePercent 		string 	`xml:"snapshot-reserve-unusable-percent"`
	UsedIncludingSnapshotReserve 		string 	`xml:"used-including-snapshot-reserve"`
	UsedIncludingSnapshotReservePercent string 	`xml:"used-including-snapshot-reserve-percent"`
	VolumeFootprints 					string 	`xml:"volume-footprints"`
	VolumeFootprintsPercent 			string 	`xml:"volume-footprints-percent"`
}

type AggrSpaceListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			AggrAttributes []AggrSpaceInfo `xml:"space-information"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

func (a *AggregateSpace) List(options *AggrSpaceOptions) (*AggrSpaceListResponse, *http.Response, error) {
	a.Params.XMLName = xml.Name{Local: "aggr-space-get-iter"}
	a.Params.AggrSpaceOptions = options
	r := AggrSpaceListResponse{}
	res, err := a.get(a, &r)
	return &r, res, err
}