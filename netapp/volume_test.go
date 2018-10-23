package netapp_test

import (
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestVolume_ListSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	tru := new(bool)
	*tru = true
	f := new(bool)
	*f = false
	options := &netapp.VolumeOptions{
		Query: &netapp.VolumeQuery{
			VolumeInfo: &netapp.VolumeInfo{
				VolumeIDAttributes: &netapp.VolumeIDAttributes{
					UUID: "98595b5a-cb3d-11e8-bf6a-00a0983afb38",
				},
			},
		},
	}
	call, _, err := c.Volume.List(options)
	checkResponseSuccess(&call.Results.ResultBase, err, t)

	info := call.Results.AttributesList[0]
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Volume Name", info.VolumeIDAttributes.Name, "test_volume_01"},
		{"QOS Name", info.VolumeQosAttributes.PolicyGroupName, "test_volume_01-1000iops"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Volume.List() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}
