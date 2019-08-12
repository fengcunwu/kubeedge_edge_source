#!/bin/bash
while((1))
do
    go run resty-test.go  
    cat file.txt | curl --data-binary @- http://192.168.226.139:9091/metrics/job/some_job/instance/some_instance
    echo "complete"
    sleep 10s
done
