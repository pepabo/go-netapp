package netapp_test

import (
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestSnapmirror_GetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Snapmirror.Get("C666", "C666:c666_root", "C666:c666_root_m1", nil)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)

	info := call.Results.Info

	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Destination Cluster", info.DestinationCluster, "lab-cluster01"},
		{"Destination Location", info.DestinationLocation, "lab-cluster01://C666/c666_root_m1"},
		{"Destination Volume", info.DestinationVolume, "c666_root_m1"},
		{"Destination VServer", info.DestinationVServer, "C666"},
		{"Is Healthy", info.IsHealthy, true},
		{"Relationship Type", info.RelationshipType, "load_sharing"},
		{"Source Cluster", info.SourceCluster, "lab-cluster01"},
		{"Source Location", info.SourceLocation, "lab-cluster01://C666/c666_root"},
		{"Source Volume", info.SourceVolume, "c666_root"},
		{"Source VServer", info.SourceVServer, "C666"},
		{"VServer", info.VServer, "C666"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Snapmirror.Get() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestSnapmirror_CreateSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	opts := &netapp.SnapmirrorInfo{
		SourceLocation:      "C666:c666_root",
		DestinationLocation: "C666:c666_root_m1",
		RelationshipType:    netapp.SnapmirrorRelationshipLS,
	}
	call, _, err := c.Snapmirror.Create("C666", opts)
	checkResponseSuccess(&call.Results.SingleResultBase, err, t)
}

func TestSnapmirror_InitializeLSSetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Snapmirror.InitializeLSSet("C666", "C666:c666_root")
	checkResponseSuccess(&call.Results.AsyncResultBase, err, t)

	job := call.Results.AsyncResultBase
	expectedJob := 27167
	if job.JobID != expectedJob {
		t.Errorf("Incorrect Job Id. Expected %d, got %d", expectedJob, job.JobID)
	}
	expectedStatus := "in_progress"
	if job.JobStatus != expectedStatus {
		t.Errorf("Incorrect job status. Exepcted %s, got %s", expectedStatus, job.JobStatus)
	}
}

func TestSnapmirror_UpdateLSSetSuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	call, _, err := c.Snapmirror.UpdateLSSet("C666", "C666:c666_root")
	checkResponseSuccess(&call.Results.AsyncResultBase, err, t)

	job := call.Results.AsyncResultBase
	expectedJob := 27168
	if job.JobID != expectedJob {
		t.Errorf("Incorrect Job Id. Expected %d, got %d", expectedJob, job.JobID)
	}
	expectedStatus := "in_progress"
	if job.JobStatus != expectedStatus {
		t.Errorf("Incorrect job status. Exepcted %s, got %s", expectedStatus, job.JobStatus)
	}
}
