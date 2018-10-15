package netapp_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestVServer_GetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.VServer.Get("G555", &netapp.VServerOptions{})
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	info := call.Results.VServerInfo

	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Aggregate List", info.AggregateList, &[]string{
			"aggr0_root_cluster01_01",
			"aggr0_root_cluster01_02",
			"n01_aggrfp_sas01",
			"n02_aggr_sata01",
		}},
		{"Allowed Protocols", info.AllowedProtocols, &[]string{"nfs"}},
		{"Anti Virus on Access Policy", info.AntivirusOnAccessPolicy, "default"},
		{"Disallowed Protocols", info.DisallowedProtocols, &[]string{"cifs", "fcp", "iscsi", "ndmp"}},
		{"Ip Space", info.Ipspace, "g555"},
		{"Is locked for changes", info.IsConfigLockedForChanges, false},
		{"Language", info.Language, "c.utf_8"},
		{"Max volumes", info.MaxVolumes, "unlimited"},
		{"Operational State", info.OperationalState, "running"},
		{"Quota Policy", info.QuotaPolicy, "default"},
		{"Root Volume", info.RootVolume, "g555_root"},
		{"Root Volume Aggregate", info.RootVolumeAggregate, "n01_aggrfp_sas01"},
		{"Vserver Name", info.VserverName, "G555"},
		{"UUID", info.UUID, "48fc46c1-9b2e-11e8-bf6a-00a0983afb38"},
		{"Vserver Type", info.VserverType, "data"},
		{"Vserver Subtype", info.VserverSubtype, "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Vserver.Get() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestVServer_GetFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.VServer.Get("non-existent-vserver", &netapp.VServerOptions{})

	results := &call.Results.SingleResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(15698, "Specified vserver not found", results, t)
}

func TestVServer_CreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	vserverSettings := &netapp.VServerInfo{
		VserverName:             "G554",
		VserverSubtype:          "default",
		RootVolume:              "g554_root",
		RootVolumeSecurityStyle: "unix",
		RootVolumeAggregate:     "test_aggr_01",
		SnapshotPolicy:          "none",
		Language:                "C.UTF-8",
		Ipspace:                 "g554",
	}

	call, _, err := c.VServer.Create(vserverSettings)
	checkResponseSuccess(&call.Results.AsyncResultBase, err, t)

	info := call.Results.VServerInfo
	job := call.Results.AsyncResultBase
	expectedJob := 27008
	if job.JobID != expectedJob {
		t.Errorf("Incorrect Job Id. Expected %d, got %d", expectedJob, job.JobID)
	}

	if info.VserverName != vserverSettings.VserverName {
		t.Errorf("Incorrect VServer name. Expected %s, got %s", vserverSettings.VserverName, info.VserverName)
	}
}

func TestVServer_CreateFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := &netapp.VServerInfo{
		VserverName:             "TEST",
		Ipspace:                 "test",
		RootVolume:              fmt.Sprintf("%s_root", "TEST"),
		RootVolumeAggregate:     "aggregate",
		RootVolumeSecurityStyle: "unix",
		SnapshotPolicy:          "none",
	}

	call, _, err := c.VServer.Create(opts)

	results := &call.Results.AsyncResultBase
	checkResponseFailure(results, err, t)

	testFailureResult(1017, `Ipspace "test" does not exist.`, results, t)
}

func TestVServer_ModifySuccess(t *testing.T) {

	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	modifyOpts := &netapp.VServerInfo{
		AllowedProtocols: &[]string{"nfs"},
	}

	call, _, err := c.VServer.Modify("T100", modifyOpts)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestVServer_DeleteSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.VServer.Delete("T100")
	checkResponseSuccess(&call.Results.ResultBase, err, t)
}
