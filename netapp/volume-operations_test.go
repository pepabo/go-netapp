package netapp_test

import (
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestVolumeOperation_CreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := &netapp.VolumeCreateOptions{
		Volume:                  "c666_root_m1",
		ContainingAggregateName: "n01_aggrfp_sas01",
		Size:                    "1G",
		SnapshotPolicy:          "none",
		VolumeType:              "dp",
	}
	call, _, err := c.VolumeOperations.Create("C666", opts)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}
