package netapp

import (
	"encoding/xml"
	"net/http"
)

type Volume struct {
	Base
	Params struct {
		XMLName xml.Name
		*VolumeOptions
	}
}

type VolumeOptions struct {
	DesiredAttributes *VolumeInfo `xml:"desired-attributes,omitempty"`
	MaxRecords        int         `xml:"max-records,omitempty"`
	Query             *VolumeInfo `xml:"query,omitempty"`
	Tag               string      `xml:"tag,omitempty"`
}
type VolumeAntivirusAttributes struct {
	OnAccessPolicy string `xml:"on-access-policy"`
}
type VolumeAutobalanceAttributes struct {
	IsAutobalanceEligible string `xml:"is-autobalance-eligible"`
}
type VolumeAutosizeAttributes struct {
	GrowThresholdPercent   string `xml:"grow-threshold-percent"`
	IsEnabled              string `xml:"is-enabled"`
	MaximumSize            string `xml:"maximum-size"`
	MinimumSize            string `xml:"minimum-size"`
	Mode                   string `xml:"mode"`
	ShrinkThresholdPercent string `xml:"shrink-threshold-percent"`
}
type VolumeDirectoryAttributes struct {
	I2PEnabled string `xml:"i2p-enabled"`
	MaxDirSize string `xml:"max-dir-size"`
	RootDirGen string `xml:"root-dir-gen"`
}
type VolumeExportAttributes struct {
	Policy string `xml:"policy"`
}
type VolumeHybridCacheAttributes struct {
	CacheRetentionPriority string `xml:"cache-retention-priority"`
	CachingPolicy          string `xml:"caching-policy"`
	Eligibility            string `xml:"eligibility"`
}
type VolumeIDAttributes struct {
	AggrList struct {
		AggrName string `xml:"aggr-name"`
	} `xml:"aggr-list"`
	Comment                 string `xml:"comment"`
	ContainingAggregateName string `xml:"containing-aggregate-name"`
	ContainingAggregateUUID string `xml:"containing-aggregate-uuid"`
	CreationTime            string `xml:"creation-time"`
	Dsid                    string `xml:"dsid"`
	Fsid                    string `xml:"fsid"`
	InstanceUUID            string `xml:"instance-uuid"`
	JunctionParentName      string `xml:"junction-parent-name"`
	JunctionPath            string `xml:"junction-path"`
	Msid                    string `xml:"msid"`
	Name                    string `xml:"name"`
	NameOrdinal             string `xml:"name-ordinal"`
	Node                    string `xml:"node"`
	Nodes                   struct {
		NodeName string `xml:"node-name"`
	} `xml:"nodes"`
	OwningVserverName string `xml:"owning-vserver-name"`
	OwningVserverUUID string `xml:"owning-vserver-uuid"`
	ProvenanceUUID    string `xml:"provenance-uuid"`
	Style             string `xml:"style"`
	StyleExtended     string `xml:"style-extended"`
	Type              string `xml:"type"`
	UUID              string `xml:"uuid"`
}
type VolumeInodeAttributes struct {
	BlockType                string `xml:"block-type"`
	FilesPrivateUsed         string `xml:"files-private-used"`
	FilesTotal               string `xml:"files-total"`
	FilesUsed                string `xml:"files-used"`
	InodefilePrivateCapacity string `xml:"inodefile-private-capacity"`
	InodefilePublicCapacity  string `xml:"inodefile-public-capacity"`
	InofileVersion           string `xml:"inofile-version"`
}
type VolumeLanguageAttributes struct {
	IsConvertUcodeEnabled string `xml:"is-convert-ucode-enabled"`
	IsCreateUcodeEnabled  string `xml:"is-create-ucode-enabled"`
	Language              string `xml:"language"`
	LanguageCode          string `xml:"language-code"`
	NfsCharacterSet       string `xml:"nfs-character-set"`
	OemCharacterSet       string `xml:"oem-character-set"`
}
type VolumeMirrorAttributes struct {
	IsDataProtectionMirror   string `xml:"is-data-protection-mirror"`
	IsLoadSharingMirror      string `xml:"is-load-sharing-mirror"`
	IsMoveMirror             string `xml:"is-move-mirror"`
	IsReplicaVolume          string `xml:"is-replica-volume"`
	MirrorTransferInProgress string `xml:"mirror-transfer-in-progress"`
	RedirectSnapshotID       string `xml:"redirect-snapshot-id"`
}
type VolumePerformanceAttributes struct {
	ExtentEnabled        string `xml:"extent-enabled"`
	FcDelegsEnabled      string `xml:"fc-delegs-enabled"`
	IsAtimeUpdateEnabled string `xml:"is-atime-update-enabled"`
	MaxWriteAllocBlocks  string `xml:"max-write-alloc-blocks"`
	MinimalReadAhead     string `xml:"minimal-read-ahead"`
	ReadRealloc          string `xml:"read-realloc"`
}
type VolumeSecurityAttributes struct {
	Style                        string `xml:"style"`
	VolumeSecurityUnixAttributes struct {
		GroupID     string `xml:"group-id"`
		Permissions string `xml:"permissions"`
		UserID      string `xml:"user-id"`
	} `xml:"volume-security-unix-attributes"`
}
type VolumeSisAttributes struct {
	CompressionSpaceSaved             string `xml:"compression-space-saved"`
	DeduplicationSpaceSaved           string `xml:"deduplication-space-saved"`
	DeduplicationSpaceShared          string `xml:"deduplication-space-shared"`
	IsSisLoggingEnabled               string `xml:"is-sis-logging-enabled"`
	IsSisStateEnabled                 string `xml:"is-sis-state-enabled"`
	IsSisVolume                       string `xml:"is-sis-volume"`
	PercentageCompressionSpaceSaved   string `xml:"percentage-compression-space-saved"`
	PercentageDeduplicationSpaceSaved string `xml:"percentage-deduplication-space-saved"`
	PercentageTotalSpaceSaved         string `xml:"percentage-total-space-saved"`
	TotalSpaceSaved                   string `xml:"total-space-saved"`
}
type VolumeSnaplockAttributes struct {
	SnaplockType string `xml:"snaplock-type"`
}
type VolumeSnapshotAttributes struct {
	AutoSnapshotsEnabled           string `xml:"auto-snapshots-enabled"`
	SnapdirAccessEnabled           string `xml:"snapdir-access-enabled"`
	SnapshotCloneDependencyEnabled string `xml:"snapshot-clone-dependency-enabled"`
	SnapshotCount                  string `xml:"snapshot-count"`
	SnapshotPolicy                 string `xml:"snapshot-policy"`
}
type VolumeSnapshotAutodeleteAttributes struct {
	Commitment          string `xml:"commitment"`
	DeferDelete         string `xml:"defer-delete"`
	DeleteOrder         string `xml:"delete-order"`
	DestroyList         string `xml:"destroy-list"`
	IsAutodeleteEnabled string `xml:"is-autodelete-enabled"`
	Prefix              string `xml:"prefix"`
	TargetFreeSpace     string `xml:"target-free-space"`
	Trigger             string `xml:"trigger"`
}
type VolumeSpaceAttributes struct {
	FilesystemSize                  string `xml:"filesystem-size"`
	IsFilesysSizeFixed              string `xml:"is-filesys-size-fixed"`
	IsSpaceGuaranteeEnabled         string `xml:"is-space-guarantee-enabled"`
	IsSpaceSloEnabled               string `xml:"is-space-slo-enabled"`
	OverwriteReserve                string `xml:"overwrite-reserve"`
	OverwriteReserveRequired        string `xml:"overwrite-reserve-required"`
	OverwriteReserveUsed            string `xml:"overwrite-reserve-used"`
	OverwriteReserveUsedActual      string `xml:"overwrite-reserve-used-actual"`
	PercentageFractionalReserve     string `xml:"percentage-fractional-reserve"`
	PercentageSizeUsed              string `xml:"percentage-size-used"`
	PercentageSnapshotReserve       string `xml:"percentage-snapshot-reserve"`
	PercentageSnapshotReserveUsed   string `xml:"percentage-snapshot-reserve-used"`
	PhysicalUsed                    string `xml:"physical-used"`
	PhysicalUsedPercent             string `xml:"physical-used-percent"`
	Size                            string `xml:"size"`
	SizeAvailable                   string `xml:"size-available"`
	SizeAvailableForSnapshots       string `xml:"size-available-for-snapshots"`
	SizeTotal                       string `xml:"size-total"`
	SizeUsed                        string `xml:"size-used"`
	SizeUsedBySnapshots             string `xml:"size-used-by-snapshots"`
	SnapshotReserveSize             string `xml:"snapshot-reserve-size"`
	SpaceFullThresholdPercent       string `xml:"space-full-threshold-percent"`
	SpaceGuarantee                  string `xml:"space-guarantee"`
	SpaceMgmtOptionTryFirst         string `xml:"space-mgmt-option-try-first"`
	SpaceNearlyFullThresholdPercent string `xml:"space-nearly-full-threshold-percent"`
	SpaceSlo                        string `xml:"space-slo"`
}
type VolumeStateAttributes struct {
	BecomeNodeRootAfterReboot string `xml:"become-node-root-after-reboot"`
	ForceNvfailOnDr           string `xml:"force-nvfail-on-dr"`
	IgnoreInconsistent        string `xml:"ignore-inconsistent"`
	InNvfailedState           string `xml:"in-nvfailed-state"`
	IsClusterVolume           string `xml:"is-cluster-volume"`
	IsConstituent             string `xml:"is-constituent"`
	IsFlexgroup               string `xml:"is-flexgroup"`
	IsInconsistent            string `xml:"is-inconsistent"`
	IsInvalid                 string `xml:"is-invalid"`
	IsJunctionActive          string `xml:"is-junction-active"`
	IsMoving                  string `xml:"is-moving"`
	IsNodeRoot                string `xml:"is-node-root"`
	IsNvfailEnabled           string `xml:"is-nvfail-enabled"`
	IsQuiescedInMemory        string `xml:"is-quiesced-in-memory"`
	IsQuiescedOnDisk          string `xml:"is-quiesced-on-disk"`
	IsUnrecoverable           string `xml:"is-unrecoverable"`
	IsVolumeInCutover         string `xml:"is-volume-in-cutover"`
	IsVserverRoot             string `xml:"is-vserver-root"`
	State                     string `xml:"state"`
}
type VolumeTransitionAttributes struct {
	IsCftPrecommit        string `xml:"is-cft-precommit"`
	IsCopiedForTransition string `xml:"is-copied-for-transition"`
	IsTransitioned        string `xml:"is-transitioned"`
	TransitionBehavior    string `xml:"transition-behavior"`
}

type VolumeInfo struct {
	Encrypt                            string                             `xml:"encrypt"`
	KeyID                              string                             `xml:"key-id,omitempty"`
	VolumeAntivirusAttributes          VolumeAntivirusAttributes          `xml:"volume-antivirus-attributes,omitempty"`
	VolumeAutobalanceAttributes        VolumeAutobalanceAttributes        `xml:"volume-autobalance-attributes,omitempty"`
	VolumeAutosizeAttributes           VolumeAutosizeAttributes           `xml:"volume-autosize-attributes"`
	VolumeDirectoryAttributes          VolumeDirectoryAttributes          `xml:"volume-directory-attributes"`
	VolumeExportAttributes             VolumeExportAttributes             `xml:"volume-export-attributes,omitempty"`
	VolumeHybridCacheAttributes        VolumeHybridCacheAttributes        `xml:"volume-hybrid-cache-attributes"`
	VolumeIDAttributes                 VolumeIDAttributes                 `xml:"volume-id-attributes"`
	VolumeInodeAttributes              VolumeInodeAttributes              `xml:"volume-inode-attributes"`
	VolumeLanguageAttributes           VolumeLanguageAttributes           `xml:"volume-language-attributes"`
	VolumeMirrorAttributes             VolumeMirrorAttributes             `xml:"volume-mirror-attributes"`
	VolumePerformanceAttributes        VolumePerformanceAttributes        `xml:"volume-performance-attributes"`
	VolumeSecurityAttributes           VolumeSecurityAttributes           `xml:"volume-security-attributes"`
	VolumeSisAttributes                VolumeSisAttributes                `xml:"volume-sis-attributes"`
	VolumeSnaplockAttributes           VolumeSnaplockAttributes           `xml:"volume-snaplock-attributes,omitempty"`
	VolumeSnapshotAttributes           VolumeSnapshotAttributes           `xml:"volume-snapshot-attributes"`
	VolumeSnapshotAutodeleteAttributes VolumeSnapshotAutodeleteAttributes `xml:"volume-snapshot-autodelete-attributes"`
	VolumeSpaceAttributes              VolumeSpaceAttributes              `xml:"volume-space-attributes"`
	VolumeStateAttributes              VolumeStateAttributes              `xml:"volume-state-attributes"`
	VolumeTransitionAttributes         VolumeTransitionAttributes         `xml:"volume-transition-attributes"`
}

type VolumeListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			VolumeAttributes []VolumeInfo `xml:"volume-attributes"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

func (v *Volume) List(options *VolumeOptions) (*VolumeListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "volume-get-iter"}
	v.Params.VolumeOptions = options
	r := VolumeListResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}

type VolumeSpaceInfo struct {
	FilesystemMetadata         string `xml:"filesystem-metadata"`
	FilesystemMetadataPercent  string `xml:"filesystem-metadata-percent"`
	Inodes                     string `xml:"inodes"`
	InodesPercent              string `xml:"inodes-percent"`
	PerformanceMetadata        string `xml:"performance-metadata"`
	PerformanceMetadataPercent string `xml:"performance-metadata-percent"`
	PhysicalUsed               string `xml:"physical-used"`
	PhysicalUsedPercent        string `xml:"physical-used-percent"`
	SnapshotReserve            string `xml:"snapshot-reserve"`
	SnapshotReservePercent     string `xml:"snapshot-reserve-percent"`
	TotalUsed                  string `xml:"total-used"`
	TotalUsedPercent           string `xml:"total-used-percent"`
	UserData                   string `xml:"user-data"`
	UserDataPercent            string `xml:"user-data-percent"`
	Volume                     string `xml:"volume"`
	Vserver                    string `xml:"vserver"`
}
type VolumeSpacesInfo []VolumeSpaceInfo

func (v VolumeSpacesInfo) Len() int {
	return len(v)
}

func (v VolumeSpacesInfo) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (p VolumeSpacesInfo) Less(i, j int) bool {
	return p[i].TotalUsedPercent < p[j].TotalUsedPercent
}

type VolumeSpace struct {
	Base
	Params struct {
		XMLName xml.Name
		*VolumeSpaceOptions
	}
}

type VolumeSpaceListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			SpaceInfo VolumeSpacesInfo `xml:"space-info"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

type VolumeSpaceOptions struct {
	DesiredAttributes *VolumeSpaceInfo `xml:"desired-attributes,omitempty"`
	MaxRecords        int              `xml:"max-records,omitempty"`
	Query             *VolumeSpaceInfo `xml:"query,omitempty"`
	Tag               string           `xml:"tag,omitempty"`
}

func (v *VolumeSpace) List(options *VolumeSpaceOptions) (*VolumeSpaceListResponse, *http.Response, error) {
	v.Params.XMLName = xml.Name{Local: "volume-space-get-iter"}
	v.Params.VolumeSpaceOptions = options
	r := VolumeSpaceListResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}
