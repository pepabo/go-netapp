package netapp_test

import (
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestVServer_ExportsRuleCreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := &netapp.VServerExportRuleInfo{
		PolicyName:         "default",
		SuperUserSecurity:  &[]string{"any"},
		Protocol:           &[]string{"any"},
		AllowSetUID:        true,
		AllowCreateDevices: true,
		ClientMatch:        "0.0.0.0/0",
		ReadWriteRule:      &[]string{"any"},
		ReadOnlyRule:       &[]string{"any"},
	}

	call, _, err := c.VServer.CreateExportRule("C666", opts)

	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestVServer_ExportsRuleDeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.VServer.DeleteExportRule("C666", "default", 1)

	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}
