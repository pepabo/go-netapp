package netapp_test

import (
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestNet_NetInterfaceCreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := &netapp.NetInterfaceInfo{
		Vserver:              "C666",
		InterfaceName:        "C666-v3666-1.1",
		Role:                 "data",
		Address:              "172.30.1.1",
		DataProtocols:        &[]string{"nfs"},
		HomeNode:             "lab-cluster01-01",
		HomePort:             "a0a-3666",
		Netmask:              "255.255.255.0",
		AdministrativeStatus: "up",
	}

	call, _, err := c.Net.CreateNetInterface(opts)

	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestNet_NetInterfaceDeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteNetInterface("C666", "C666-v3666-1.1")

	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestNet_NetRouteCreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := &netapp.NetRoutesInfo{
		DestinationAddress: "0.0.0.0/0",
		GatewayAddress:     "172.30.1.199",
		Metric:             20,
		ReturnRecord:       true,
	}

	call, _, err := c.Net.CreateRoute("C666", opts)

	results := call.Results
	checkResponseSuccess(&results.SingleResultBase, err, t)

	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Address Family", results.Info.AddressFamily, "ipv4"},
		{"Destination Address", results.Info.DestinationAddress, opts.DestinationAddress},
		{"Gateway Address", results.Info.GatewayAddress, opts.GatewayAddress},
		{"Metric", results.Info.Metric, opts.Metric},
		{"VServer", results.Info.VServer, "C666"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Vserver.Get() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestNet_NetRouteDeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Net.DeleteRoute("C666", "0.0.0.0/0", "172.30.1.199")

	results := call.Results
	checkResponseSuccess(&results.SingleResultBase, err, t)
}
