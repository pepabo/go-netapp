package netapp_test

import (
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestNet_BroadcastDomainGetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.GetBroadcastDomain("test-bdomain", "test-net-ipspace")
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	info := call.Results.Info
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Broadcast Domains", info.BroadcastDomain, "test-bdomain"},
		{"Failover Groups", info.FailoverGroups, []string{"test-bdomain"}},
		{"IPSpace", info.IPSpace, "test-net-ipspace"},
		{"MTU", info.MTU, 1500},
		{"Port", (*info.Ports)[0], netapp.NetPortUpdateInfo{
			Port:                    "lab-cluster01-02:a0a-3555",
			PortUpdateStatus:        "complete",
			PortUpdateStatusDetails: "complete",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Net.GetBroadcastDomain() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestNet_BroadcastDomainGetFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.GetBroadcastDomain("other-bdomain", "test-ip-space")

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(15661, `entry doesn"t exist`, results, t)
}

func TestNet_BroadcastDomainCreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.CreateBroadcastDomain(&netapp.NetBroadcastDomainCreateOptions{
		BroadcastDomain: "test-bcastdomain",
		IPSpace:         "test-ip-space",
		MTU:             1500,
		Ports:           &[]string{"lab-cluster01-01:a0a-3555", "lab-cluster01-02:a0a-3555"},
	})
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	expected := "complete"
	if call.Results.CombinedPortUpdateStatus != expected {
		t.Errorf("Response was %s, expected %s", call.Results.CombinedPortUpdateStatus, expected)
	}
}

func TestNet_BroadcastDomainCreateFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.CreateBroadcastDomain(&netapp.NetBroadcastDomainCreateOptions{
		BroadcastDomain: "test-bcastdomain",
		IPSpace:         "test-ip-space",
		MTU:             1500,
	})

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(18603, "Specified IPspace not found", results, t)
}

func TestNet_BroadcastDomainDeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteBroadcastDomain("test-bcastdomain", "test-ip-space")
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	expected := "complete"
	if call.Results.CombinedPortUpdateStatus != expected {
		t.Errorf("Response was %s, expected %s", call.Results.CombinedPortUpdateStatus, expected)
	}
}

func TestNet_BroadcastDomainDeleteFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteBroadcastDomain("test-bcastdomain", "test-ip-space")

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(18604, `Broadcast domain "test-bcastdomain" in IPspace "test-ip-space" not found.`, results, t)
}
