package netapp_test

import (
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestQosPolicy_GetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.QosPolicy.Get("qos-policy-test", &netapp.QosPolicyInfo{})
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	info := call.Results.QosPolicyInfo

	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"IOPS", info.MaxThroughput, "1000IOPS"},
		{"NumWorkloads", info.NumWorkloads, 1},
		{"Pgid", info.PgID, 12966},
		{"Name", info.PolicyGroup, "qos-policy-test"},
		{"UUID", info.UUID, "9112f935-9b41-11e8-bf6a-00a0983afb38"},
		{"Vserver", info.VServer, "test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("QosPolicy.GetSuccess() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestQosPolicy_GetNotFound(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.QosPolicy.Get("qos-policy-test-not-found", &netapp.QosPolicyInfo{})

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(15661, "entry doesn't exist", results, t)
}

func TestQosPolicy_CreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := netapp.QosPolicyInfo{
		PolicyGroup:   "qos-policy-create-test",
		MaxThroughput: "3000iops",
		VServer:       "test-vserver",
	}
	call, _, err := c.QosPolicy.Create(&opts)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestQosPolicy_CreateFailureBecauseQosPolicyExists(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := netapp.QosPolicyInfo{
		VServer:       "test-vserver",
		PolicyGroup:   "qos-policy-create-test",
		MaxThroughput: "3000iops",
	}
	call, _, err := c.QosPolicy.Create(&opts)
	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)
	testFailureResult(13130, "duplicate entry", results, t)
}

func TestQosPolicy_RenameSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := netapp.QosPolicyRenameInfo{
		CurrentPolicyGroup: "qos-policy-rename-test",
		NewPolicyGroup:     "qos-policy-rename-test-new",
	}
	call, _, err := c.QosPolicy.Rename(&opts)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestQosPolicy_RenameFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := netapp.QosPolicyRenameInfo{
		CurrentPolicyGroup: "qos-policy-rename-test-fail",
		NewPolicyGroup:     "qos-policy-rename-test-fail-new",
	}
	call, _, err := c.QosPolicy.Rename(&opts)
	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)
	testFailureResult(18339, "Policy with new name already exists.", results, t)

}

func TestQosPolicy_ChangeIopsSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.QosPolicy.ChangeIops("1000IOPS", "qos-policy-iops-change")
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestQosPolicy_ChangeIopsFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.QosPolicy.ChangeIops("1000 dalmations", "qos-policy-iops-change")
	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)
	testFailureResult(13115, `Invalid value specified for "max-throughput" element within "qos-policy-group-modify": "1000dalmations".`, results, t)
}

func TestQosPolicy_DeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.QosPolicy.Delete("qos-policy-iops-delete", false)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestQosPolicy_DeleteFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.QosPolicy.Delete("qos-policy-iops-delete-fail", false)
	checkResponseFailure(&call.Results.SingleResultBase, err, t)
}
