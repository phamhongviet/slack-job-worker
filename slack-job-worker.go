/*
A template for building workers to work with slack-jobs
https://github.com/phamhongviet/slack-jobs
*/
package main

import (
	"fmt"
	"github.com/benmanns/goworker"
	"strings"
)

// struct Job represent the job received from resque
type Job struct {
	Request   string
	User      string
	Channel   string
	Timestamp string
}

// ParseJob create a new Job from arguments
func ParseJob(args []interface{}) Job {
	var job Job
	for _, v := range args {
		v := v.(string)
		switch {
		case strings.HasPrefix(v, "request="):
			job.Request = strings.TrimPrefix(v, "request=")
		case strings.HasPrefix(v, "user="):
			job.User = strings.TrimPrefix(v, "user=")
		case strings.HasPrefix(v, "timestamp="):
			job.Timestamp = strings.TrimPrefix(v, "timestamp=")
		case strings.HasPrefix(v, "channel_name="):
			job.Channel = strings.TrimPrefix(v, "channel_name=")
		}
	}
	return job
}

// Parse has the same logic as ParseJob
// Parse can update an existing Job
func (job *Job) Parse(args []interface{}) {
	for _, v := range args {
		v := v.(string)
		switch {
		case strings.HasPrefix(v, "request="):
			job.Request = strings.TrimPrefix(v, "request=")
		case strings.HasPrefix(v, "user="):
			job.User = strings.TrimPrefix(v, "user=")
		case strings.HasPrefix(v, "timestamp="):
			job.Timestamp = strings.TrimPrefix(v, "timestamp=")
		case strings.HasPrefix(v, "channel_name="):
			job.Channel = strings.TrimPrefix(v, "channel_name=")
		}
	}
}

// register worker with class and handle function
func init() {
	goworker.Register("SlackOPS", slackOPSWorker)
}

// simple handle function
func slackOPSWorker(queue string, args ...interface{}) error {
	// parse job and print its argument
	job := ParseJob(args)
	fmt.Printf("JOB: %s\n", job)
	return nil
}

// let the worker do its job :)
func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}
