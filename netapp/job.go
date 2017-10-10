package netapp

import (
	"encoding/xml"
	"net/http"
)

type Job struct {
	Base
	Params struct {
		XMLName xml.Name
		*JobOptions
	}
}

type JobOptions struct {
	DesiredAttributes *JobEntry `xml:"desired-attributes,omitempty"`
	MaxRecords        int       `xml:"max-records,omitempty"`
	Tag               string    `xml:"tag,omitempty"`
	*JobEntry
}

type JobResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		Attributes struct {
			JobInfo struct {
				IsRestarted              string `xml:"is-restarted"`
				JobAffinity              string `xml:"job-affinity"`
				JobCategory              string `xml:"job-category"`
				JobCompletion            string `xml:"job-completion"`
				JobDescription           string `xml:"job-description"`
				JobID                    string `xml:"job-id"`
				JobName                  string `xml:"job-name"`
				JobNode                  string `xml:"job-node"`
				JobPriority              string `xml:"job-priority"`
				JobProcess               string `xml:"job-process"`
				JobProgress              string `xml:"job-progress"`
				JobQueueTime             string `xml:"job-queue-time"`
				JobRestartIsOrWasDelayed string `xml:"job-restart-is-or-was-delayed"`
				JobSchedule              string `xml:"job-schedule"`
				JobStartTime             string `xml:"job-start-time"`
				JobState                 string `xml:"job-state"`
				JobStatusCode            string `xml:"job-status-code"`
				JobType                  string `xml:"job-type"`
				JobUsername              string `xml:"job-username"`
				JobUUID                  string `xml:"job-uuid"`
				JobVserver               string `xml:"job-vserver"`
			} `xml:"job-info"`
		} `xml:"attributes"`
	} `xml:"results"`
}

type JobEntry struct {
	ID int `xml:"job-id"`
}

func (j *JobResponse) JobState() string {
	return j.Results.Attributes.JobInfo.JobState
}

func (j *JobResponse) Success() bool {
	return j.JobState() == "success"
}

func (q *Job) Get(vserverName string, id int, options *JobOptions) (*JobResponse, *http.Response, error) {
	q.Name = vserverName
	if options == nil {
		options = &JobOptions{
			JobEntry: &JobEntry{},
		}
	}
	options.JobEntry.ID = id
	q.Params.JobOptions = options
	q.Params.XMLName = xml.Name{Local: "job-get"}
	r := JobResponse{}
	res, err := q.get(q, &r)
	return &r, res, err
}
