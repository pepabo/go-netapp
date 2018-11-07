package netapp_test

import (
	"reflect"
	"strings"
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

func TestSnapmirror_DestroyByFailure(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	query := &netapp.SnapmirrorInfo{
		VServer: "T100",
	}

	call, _, err := c.Snapmirror.DestroyBy(query, true)

	if err != nil {
		t.Fatalf("Should not have gotten an error %s", err)
	}

	results := call.Results.FailureList

	if len(results) == 0 || call.Results.NumFailed == 0 {
		t.Error("Got back 0 failures, should've had 2")
	}

	if call.Results.NumFailed != 2 {
		t.Errorf("%s got = %+v, want %+v", t.Name(), call.Results.NumFailed, 2)
	}

	if call.Results.NumSucceeded != 0 {
		t.Errorf("%s got = %+v, want %+v", t.Name(), call.Results.NumSucceeded, 0)
	}

	for _, result := range results {
		if result.ErrorNo != 13001 {
			t.Errorf("%s got = %+v, want %+v", t.Name(), result.ErrorNo, 13001)
		}
		msg := "SnapMirror: error: Failed to change the volume T100_root"
		if !strings.Contains(result.Reason, msg) {
			t.Errorf("%s got = %+v, want to contain %+v", t.Name(), result.Reason, msg)
		}

		if result.Info == nil {
			t.Errorf("%s got empty snapmirror object, should have values", t.Name())
		}
	}
}

func TestSnapmirror_DestroyBySuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	query := &netapp.SnapmirrorInfo{
		VServer: "T100",
	}

	call, _, err := c.Snapmirror.DestroyBy(query, true)

	if err != nil {
		t.Fatalf("Should not have gotten an error %s", err)
	}

	results := call.Results.SuccessList

	if len(results) == 0 {
		t.Error("Got back 0 results")
	}

	if call.Results.NumSucceeded != 2 {
		t.Errorf("%s got = %+v, want %+v", t.Name(), call.Results.NumSucceeded, 2)
	}

	if call.Results.NumFailed != 0 {
		t.Errorf("%s got = %+v, want %+v", t.Name(), call.Results.NumFailed, 0)
	}
}

func TestSnapmirror_AbortBySuccess(t *testing.T) {

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
