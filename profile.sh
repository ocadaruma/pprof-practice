#!/bin/bash

set -eu

go build slow_app.go && ./slow_app &
go build app.go && ./app &

wrk -t5 -c10 -d45s http://localhost:8080/blocking &
wrk -t5 -c10 -d45s http://localhost:8080/cpu-intensive &

# wait for warming up
sleep 5

pprof -http=":8088" app localhost:8080/debug/pprof/profile
