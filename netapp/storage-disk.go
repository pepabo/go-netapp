package netapp

import (
	"encoding/xml"
	"net/http"
)

type StorageDisk struct {
	Base
	Params struct {
		XMLName xml.Name
		*StorageDiskOptions
	}
}

type StorageDiskInfo struct {
	DiskInventoryInfo DiskInventoryInfo `xml:"disk-inventory-info"`
	DiskName          string            `xml:"disk-name"`
	DiskOwnershipInfo DiskOwnershipInfo `xml:"disk-ownership-info"`
}

type DiskInventoryInfo struct {
	BytesPerSector                 int    `xml:"bytes-per-sector"`
	CapacitySectors                int    `xml:"capacity-sectors"`
	ChecksumCompatibility          string `xml:"checksum-compatibility"`
	DiskClusterName                string `xml:"disk-cluster-name"`
	DiskType                       string `xml:"disk-type"`
	DiskUid                        string `xml:"disk-uid"`
	FirmwareRevision               string `xml:"firmware-revision"`
	GrownDefectListCount           int    `xml:"grown-defect-list-count"`
	HealthMonitorTimeInterval      int    `xml:"health-monitor-time-interval"`
	ImportInProgress               bool   `xml:"import-in-progress"`
	IsDynamicallyQualified         bool   `xml:"is-dynamically-qualified"`
	IsMultidiskCarrier             bool   `xml:"is-multidisk-carrier"`
	IsShared                       bool   `xml:"is-shared"`
	MediaScrubCount                int    `xml:"media-scrub-count"`
	MediaScrubLastDoneTimeInterval int    `xml:"media-scrub-last-done-time-interval"`
	Model                          string `xml:"model"`
	ReservationKey                 string `xml:"reservation-key"`
	ReservationType                string `xml:"reservation-type"`
	RightSizeSectors               int    `xml:"right-size-sectors"`
	Rpm                            int    `xml:"rpm"`
	SerialNumber                   string `xml:"serial-number"`
	Shelf                          string `xml:"shelf"`
	ShelfBay                       string `xml:"shelf-bay"`
	ShelfUid                       string `xml:"shelf-uid"`
	StackID                        int    `xml:"stack-id"`
	Vendor                         string `xml:"vendor"`
}

type DiskOwnershipInfo struct {
	DiskUid          string `xml:"disk-uid"`
	DrHomeNodeId     int    `xml:"dr-home-node-id"`
	DrHomeNodeName   string `xml:"dr-home-node-name"`
	HomeNodeId       int    `xml:"home-node-id"`
	HomeNodeName     string `xml:"home-node-name"`
	IsFailed         bool   `xml:"is-failed"`
	OwnerNodeId      int    `xml:"owner-node-id"`
	OwnerNodeName    string `xml:"owner-node-name"`
	Pool             int    `xml:"pool"`
	ReservedByNodeId int    `xml:"reserved-by-node-id"`
}

type StorageDiskGetIterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			StorageDiskInfo []StorageDiskInfo `xml:"storage-disk-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

type StorageDiskInfoPageResponse struct {
	Response    *StorageDiskGetIterResponse
	Error       error
	RawResponse *http.Response
}

type StorageDiskOptions struct {
	MaxRecords int    `xml:"max-records,omitempty"`
	Tag        string `xml:"tag,omitempty"`
}

func (s *StorageDisk) StorageDiskGetIter(options *StorageDiskOptions) (*StorageDiskGetIterResponse, *http.Response, error) {
	s.Params.XMLName = xml.Name{Local: "storage-disk-get-iter"}
	s.Params.StorageDiskOptions = options
	r := StorageDiskGetIterResponse{}
	res, err := s.get(s, &r)
	return &r, res, err
}

type StorageDiskGetAllPageHandler func(StorageDiskInfoPageResponse) (shouldContinue bool)

func (s *StorageDisk) StorageDiskGetAll(options *StorageDiskOptions, fn StorageDiskGetAllPageHandler) {

	requestOptions := options
	for shouldContinue := true; shouldContinue; {
		storageDiskGetIterResponse, res, err := s.StorageDiskGetIter(requestOptions)
		handlerResponse := false

		handlerResponse = fn(StorageDiskInfoPageResponse{Response: storageDiskGetIterResponse, Error: err, RawResponse: res})

		nextTag := ""
		if err == nil {
			nextTag = storageDiskGetIterResponse.Results.NextTag
			requestOptions = &StorageDiskOptions{
				Tag:        nextTag,
				MaxRecords: requestOptions.MaxRecords,
			}
		}
		shouldContinue = nextTag != "" && handlerResponse
	}
}
