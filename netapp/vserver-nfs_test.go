package netapp_test

import (
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestVServer_NfsCreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	nfsOpts := &netapp.VServerNfsCreateOptions{
		NfsAccessEnabled: true,
		NfsV3Enabled:     true,
		NfsV4Enabled:     true,
		VStorageEnabled:  true,
	}

	call, _, err := c.VServer.CreateNfsService("C666", nfsOpts)

	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}
