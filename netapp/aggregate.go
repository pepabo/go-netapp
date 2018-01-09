package netapp

import (
	"encoding/xml"
	"net/http"
)

type Aggregate struct {
	Base
	Params struct {
		XMLName xml.Name
		*AggrOptions
	}
}
type AggrQuery struct {
	AggrEntry *AggrInfo `xml:"aggr-attributes,omitempty"`
}

type AggrOptions struct {
	DesiredAttributes *AggrQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int        `xml:"max-records,omitempty"`
	Query             *AggrQuery `xml:"query,omitempty"`
	Tag               string     `xml:"tag,omitempty"`
}

type AggrListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			AggrAttributes []AggrInfo `xml:"aggr-attributes"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (a *Aggregate) List(options *AggrOptions) (*AggrListResponse, *http.Response, error) {
	a.Params.XMLName = xml.Name{Local: "aggr-get-iter"}
	a.Params.AggrOptions = options
	r := AggrListResponse{}
	res, err := a.get(a, &r)
	return &r, res, err
}

type AggrListPagesResponse struct {
	Response    *AggrListResponse
	Error       error
	RawResponse *http.Response
}

type AggregatePageHandler func(AggrListPagesResponse) (shouldContinue bool)

func (a *Aggregate) ListPages(options *AggrOptions, fn AggregatePageHandler) {

	requestOptions := options

	for shouldContinue := true; shouldContinue; {
		aggregateResponse, res, err := a.List(requestOptions)
		handlerResponse := false

		handlerResponse = fn(AggrListPagesResponse{Response: aggregateResponse, Error: err, RawResponse: res})

		nextTag := ""
		if err == nil {
			nextTag = aggregateResponse.Results.NextTag
			requestOptions = &AggrOptions{
				Tag:        nextTag,
				MaxRecords: options.MaxRecords,
			}
		}
		shouldContinue = nextTag != "" && handlerResponse
	}

}

type AggrInfo struct {
	AggregateName       string              `xml:"aggregate-name"`
	AggrInodeAttributes AggrInodeAttributes `xml:"aggr-inode-attributes"`
	AggrSpaceAttributes AggrSpaceAttributes `xml:"aggr-space-attributes"`
}

type AggrInodeAttributes struct {
	FilesPrivateUsed         string `xml:"files-private-used"`
	FilesTotal               string `xml:"files-total"`
	FilesUsed                string `xml:"files-used"`
	InodefilePrivateCapacity string `xml:"inodefile-private-capacity"`
	InodefilePublicCapacity  string `xml:"inodefile-public-capacity"`
	MaxfilesAvailable        string `xml:"maxfiles-available"`
	MaxfilesPossible         string `xml:"maxfiles-possible"`
	MaxfilesUsed             string `xml:"maxfiles-used"`
	PercentInodeUsedCapacity string `xml:"percent-inode-used-capacity"`
}

type AggrSpaceAttributes struct {
	AggregateMetadata            string `xml:"aggregate-metadata"`
	HybridCacheSizeTotal         string `xml:"hybrid-cache-size-total"`
	PercentUsedCapacity          string `xml:"percent-used-capacity"`
	PhysicalUsed                 string `xml:"physical-used"`
	PhysicalUsedPercent          string `xml:"physical-used-percent"`
	SizeAvailable                string `xml:"size-available"`
	SizeTotal                    string `xml:"size-total"`
	SizeUsed                     string `xml:"size-used"`
	TotalReservedSpace           string `xml:"total-reserved-space"`
	UsedIncludingSnapshotReserve string `xml:"used-including-snapshot-reserve"`
	VolumeFootprints             string `xml:"volume-footprints"`
}

type AggregateSpace struct {
	Base
	Params struct {
		XMLName xml.Name
		*AggrSpaceOptions
	}
}
type AggrSpaceInfoQuery struct {
	AggrSpaceInfo *AggrSpaceInfo `xml:"space-information,omitempty"`
}

type AggrSpaceOptions struct {
	DesiredAttributes *AggrSpaceInfoQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                 `xml:"max-records,omitempty"`
	Query             *AggrSpaceInfoQuery `xml:"query,omitempty"`
	Tag               string              `xml:"tag,omitempty"`
}

type AggrSpaceInfo struct {
	Aggregate                           string `xml:"aggregate"`
	AggregateMetadata                   string `xml:"aggregate-metadata"`
	AggregateMetadataPercent            string `xml:"aggregate-metadata-percent"`
	AggregateSize                       string `xml:"aggregate-size"`
	PercentSnapshotSpace                string `xml:"percent-snapshot-space"`
	PhysicalUsed                        string `xml:"physical-used"`
	PhysicalUsedPercent                 string `xml:"physical-used-percent"`
	SnapSizeTotal                       string `xml:"snap-size-total"`
	SnapshotReserveUnusable             string `xml:"snapshot-reserve-unusable"`
	SnapshotReserveUnusablePercent      string `xml:"snapshot-reserve-unusable-percent"`
	UsedIncludingSnapshotReserve        string `xml:"used-including-snapshot-reserve"`
	UsedIncludingSnapshotReservePercent string `xml:"used-including-snapshot-reserve-percent"`
	VolumeFootprints                    string `xml:"volume-footprints"`
	VolumeFootprintsPercent             string `xml:"volume-footprints-percent"`
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

type AggregateSpares struct {
	Base
	Params struct {
		XMLName xml.Name
		*AggrSparesOptions
	}
}

func (a *AggregateSpares) List(options *AggrSparesOptions) (*AggrSparesListResponse, *http.Response, error) {
	a.Params.XMLName = xml.Name{Local: "aggr-spare-get-iter"}
	a.Params.AggrSparesOptions = options
	r := AggrSparesListResponse{}
	res, err := a.get(a, &r)
	return &r, res, err
}

func (a *AggregateSpares) ListPages(options *AggrSparesOptions, fn AggregateSparesPageHandler) {

	requestOptions := options

	for shouldContinue := true; shouldContinue; {
		aggregateResponse, res, err := a.List(requestOptions)
		handlerResponse := false

		handlerResponse = fn(AggrSparesListPagesResponse{Response: aggregateResponse, Error: err, RawResponse: res})

		nextTag := ""
		if err == nil {
			nextTag = aggregateResponse.Results.NextTag
			requestOptions = &AggrSparesOptions{
				Tag:        nextTag,
				MaxRecords: options.MaxRecords,
			}
		}
		shouldContinue = nextTag != "" && handlerResponse
	}
}

type AggrSpareDiskInfoQuery struct {
	AggrSpareDiskInfo *AggrSpareDiskInfo `xml:"aggr-spare-disk-info,omitempty"`
}

type AggrSparesOptions struct {
	DesiredAttributes *AggrSpareDiskInfoQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                     `xml:"max-records,omitempty"`
	Query             *AggrSpareDiskInfoQuery `xml:"query,omitempty"`
	Tag               string                  `xml:"tag,omitempty"`
}

type AggregateSparesPageHandler func(AggrSparesListPagesResponse) (shouldContinue bool)

type AggrSpareDiskInfo struct {
	ChecksumStyle           string `xml:"checksum-style"`
	Disk                    string `xml:"disk"`
	DiskRpm                 int    `xml:"disk-rpm"`
	DiskType                string `xml:"disk-type"`
	EffectiveDiskRpm        int    `xml:"effective-disk-rpm"`
	EffectiveDiskType       string `xml:"effective-disk-type"`
	IsDiskLeftBehind        bool   `xml:"is-disk-left-behind"`
	IsDiskShared            bool   `xml:"is-disk-shared"`
	IsDiskZeroed            bool   `xml:"is-disk-zeroed"`
	IsDiskZeroing           bool   `xml:"is-disk-zeroing"`
	IsSparecore             bool   `xml:"is-sparecore"`
	LocalUsableDataSize     int    `xml:"local-usable-data-size"`
	LocalUsableDataSizeBlks int    `xml:"local-usable-data-size-blks"`
	LocalUsableRootSize     int    `xml:"local-usable-root-size"`
	LocalUsableRootSizeBlks int    `xml:"local-usable-root-size-blks"`
	OriginalOwner           string `xml:"original-owner"`
	SyncmirrorPool          string `xml:"syncmirror-pool"`
	TotalSize               int    `xml:"total-size"`
	UsableSize              int    `xml:"usable-size"`
	UsableSizeBlks          int    `xml:"usable-size-blks"`
	ZeroingPercent          int    `xml:"zeroing-percent"`
}

type AggrSparesListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			AggrAttributes []AggrSpareDiskInfo `xml:"aggr-spare-disk-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

type AggrSparesListPagesResponse struct {
	Response    *AggrSparesListResponse
	Error       error
	RawResponse *http.Response
}
