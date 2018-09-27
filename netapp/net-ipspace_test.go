package netapp_test

import (
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestNet_IPSpaceCreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.CreateIPSpace("test-ip-space", true)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	info := call.Results.NetIPSpaceCreateInfo
	var nilSlice []string
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Broadcast Domains", info.BroadcastDomains, nilSlice},
		{"ID", info.ID, 39},
		{"IP Space", info.IPSpace, "test-ip-space"},
		{"Ports", info.Ports, nilSlice},
		{"UUID", info.UUID, "af9fa457-c02d-11e8-bf6a-00a0983afb38"},
		{"VServers", info.VServers, []string{"test-ip-space"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Net.CreateIPSpace() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestNet_IPSpaceCreateFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.CreateIPSpace("test-ip-space", false)

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(13001, `Invalid IPspace name. The name "test-ip-space" is already in use by a cluster node, Vserver, or is the name of the local cluster.`, results, t)
}

func TestNet_IPSpaceGetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.GetIPSpace("test-net-ipspace")
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	var createInfo *netapp.NetIPSpaceInfo
	info := call.Results.NetIPSpaceInfo
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Broadcast Domains", info.BroadcastDomains, []string{"test-net-ipspace"}},
		{"Ports", info.Ports, []string{"lab-cluster01-02:a0a-3555", "lab-cluster01-01:a0a-3555"}},
		{"VServers", info.VServers, []string{"Test-Vserver", "test-vserver"}},
		{"Create Info is empty", call.Results.NetIPSpaceCreateInfo, createInfo},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Net.GetIPSpace() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestNet_IPSpaceGetFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.GetIPSpace("test-ip-space")

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(15661, `entry doesn"t exist`, results, t)

}

func TestNet_IPSpaceRenameSuccess(t *testing.T) {

	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.RenameIPSpace("test-net-ipspace", "test-net-ipspace-new")
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestNet_IPSpaceRenameFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.RenameIPSpace("test-net-ipspace", "test-net-ipspace-new")
	checkResponseFailure(&call.Results.SingleResultBase, err, t)

	testFailureResult(13001, "IPspace test-net-ipspace does not exist.", &call.Results.SingleResultBase, t)
}

func TestNet_IPSpaceDeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteIPSpace("test-net-ipspace-new")
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestNet_IPSpaceDeleteFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteIPSpace("test-net-ipspace-new")
	checkResponseFailure(&call.Results.SingleResultBase, err, t)

	testFailureResult(15661, `entry doesn"t exist`, &call.Results.SingleResultBase, t)
}
