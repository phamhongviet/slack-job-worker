#!/bin/bash

./slack-job-worker -concurrency=1 -queues=examplequeue -uri=redis://localhost:6379/ -use-number=true \
| while read job; do
	request=$(awk -F', ' '{ print $1 }' <<<$job)
	user=$(awk -F', ' '{ print $2 }' <<<$job)
	channel=$(awk -F', ' '{ print $3 }' <<<$job)
	timestamp=$(awk -F', ' '{ print $4 }' <<<$job)
	echo "$user request to $request from $channel at $timestamp"
done
