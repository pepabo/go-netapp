package netapp_test

import (
	"reflect"
	"testing"

	"github.com/pepabo/go-netapp/netapp"
)

func TestJob_GetHistorySuccess(t *testing.T) {
	c, teardown := createTestClientWithFixtures(t)
	defer teardown()

	query := &netapp.JobHistoryOptions{
		Query: &netapp.JobHistoryInfo{
			JobID: 27292,
		},
		MaxRecords: 20,
	}
	call, _, err := c.Job.GetHistory(query)
	checkResponseSuccess(&call.Results.ResultBase, err, t)

	info := call.Results.HistoryInfo[0]
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"UUID", info.JobUUID, "0106994d-c8c9-11e8-a192-00a0983afb00"},
		{"Job Event Type", info.JobEventType, "succeeded"},
		{"Job Username", info.JobUsername, "swautomation"},
		{"Job Node", info.JobNode, "test-cluster01-02"},
		{"Job Name", info.JobName, "Vserver Create "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Job.GetHistory() got = %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}
