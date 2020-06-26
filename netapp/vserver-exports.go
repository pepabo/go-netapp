package netapp

import (
	"encoding/xml"
	"net/http"
)

type vServerExportsRequest struct {
	Base
	Params struct {
		XMLName               xml.Name
		VServerExportRuleInfo `xml:",innerxml"`
	}
}type  VServerExportRuleQuery struct {
	VServerExportRuleInfo *LunInfo `xml:"export-rule-info,omitempty"`
}



// VServerExportRuleInfo sets all different options for Export Rules
type VServerExportRuleInfo struct {
	AnonymousUserID           int       `xml:"anonymous-user-id,omitempty"`
	ClientMatch               string    `xml:"client-match,omitempty"`
	ExportChownMode           string    `xml:"export-chown-mode,omitempty"`
	ExportNTFSUnixSecurityOps string    `xml:"export-ntfs-unix-security-ops,omitempty"`
	AllowCreateDevices        bool      `xml:"is-allow-dev-is-enabled,omitempty"`
	AllowSetUID               bool      `xml:"is-allow-set-uid-enabled,omitempty"`
	PolicyName                string    `xml:"policy-name,omitempty"`
	Protocol                  *[]string `xml:"protocol>access-protocol,omitempty"`
	ReadOnlyRule              *[]string `xml:"ro-rule>security-flavor,omitempty"`
	RuleIndex                 int       `xml:"rule-index,omitempty"`
	ReadWriteRule             *[]string `xml:"rw-rule>security-flavor,omitempty"`
	SuperUserSecurity         *[]string `xml:"super-user-security>security-flavor,omitempty"`
}

// VServerExportsResponse creates correct response obj
type VServerExportsResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		SingleResultBase
	} `xml:"results"`
}

type VServerExportRuleListResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			VserverExportRuleInfo []VServerExportRuleInfo `xml:"export-rule-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

type VServerExportRuleListPagesResponse struct {
	Response    *LunListResponse
	Error       error
	RawResponse *http.Response
}

// ListExportRules list the rules of an export policy
func (v VServer) ListExportRules(vServerName string, exportPolicy string, options *VServerExportRuleInfo) (*VServerExportRuleListResponse, *http.Response, error) {
	v.Name = vServerName
	v.Params.XMLName = xml.Name{Local: "export-rule-get-iter"}
	v.Params.VServerExportRuleInfo = VServerExportRuleInfo{
		PolicyName: exportPolicy,
	}

	r := VServerExportRuleListResponse{}
	res, err := v.get(v, &r)
	return &r, res, err
}

// CreateExportRule creates a new export rule for a given vserver
func (v VServer) CreateExportRule(vServerName string, options *VServerExportRuleInfo) (*VServerExportsResponse, *http.Response, error) {
	req := v.newVServerExportsRequest()
	req.Base.Name = vServerName
	req.Params.XMLName = xml.Name{Local: "export-rule-create"}
	req.Params.VServerExportRuleInfo = *options

	r := &VServerExportsResponse{}
	res, err := v.get(req, r)
	return r, res, err
}

// DeleteExportRule removes an export rule for a given vserver, policy and rule index
func (v VServer) DeleteExportRule(vServerName string, policyName string, ruleIndex int) (*VServerExportsResponse, *http.Response, error) {
	req := v.newVServerExportsRequest()
	req.Base.Name = vServerName
	req.Params.XMLName = xml.Name{Local: "export-rule-destroy"}
	req.Params.VServerExportRuleInfo = VServerExportRuleInfo{
		PolicyName: policyName,
		RuleIndex:  ruleIndex,
	}

	r := &VServerExportsResponse{}
	res, err := v.get(req, r)
	return r, res, err
}

func (v VServer) newVServerExportsRequest() *vServerExportsRequest {
	return &vServerExportsRequest{
		Base: v.Base,
	}
}
