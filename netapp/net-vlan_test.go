package netapp_test

import (
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestNet_VlanGetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.GetVlan("a0a-3555", "test-cluster-01-01")
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	info := call.Results.Info
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Interface Name", info.InterfaceName, "a0a-3555"},
		{"Node", info.Node, "test-cluster-01-01"},
		{"Parent Interface", info.ParentInterface, "a0a"},
		{"VLanID", info.VlanID, 3555},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Net.GetVlan() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestNet_VlanGetFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.GetVlan("a0a-3555", "test-cluster-01-01")

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(13115, `Invalid value specified for "node" element within "net-vlan-get": "test-cluster-01-01".`, results, t)
}

func TestNet_VlanListSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	options := &netapp.NetVlanInfo{
		ParentInterface: "a0a",
		VlanID:          3555,
	}
	call, _, err := c.Net.ListVlans(options)
	checkResponseSuccess(&call.Results.ResultBase, err, t)

	info := call.Results.Info[0]
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Interface Name", info.InterfaceName, "a0a-3555"},
		{"Node", info.Node, "test-cluster-01-01"},
		{"Parent Interface", info.ParentInterface, "a0a"},
		{"VLanID", info.VlanID, 3555},
		{"Port String", info.ToString(), "test-cluster-01-01:a0a-3555"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Net.GetVlan() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}
func TestNet_VlanCreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.CreateVlan(&netapp.NetVlanInfo{
		InterfaceName:   "I shouldn't be sent to the server",
		ParentInterface: "a0a",
		Node:            "test-cluster-01-01",
		VlanID:          3555,
	})
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestNet_VlanCreateFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.CreateVlan(&netapp.NetVlanInfo{
		ParentInterface: "a0a",
		Node:            "test-cluster-01-01",
		VlanID:          3555,
	})

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(13115, `Invalid value specified for "node" element within "vlan-info": "test-cluster-01-01".`, results, t)
}

func TestNet_VlanDeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteVlan(&netapp.NetVlanInfo{
		ParentInterface: "a0a",
		Node:            "test-cluster-01-01",
		VlanID:          3555,
	})
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestNet_VlanDeleteFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteVlan(&netapp.NetVlanInfo{
		ParentInterface: "a0a",
		Node:            "test-cluster-01-01",
		VlanID:          3455,
	})

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(13115, `Invalid value specified for "node" element within "vlan-info": "test-cluster-01-01".`, results, t)
}
