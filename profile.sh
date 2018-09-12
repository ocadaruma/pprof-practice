#!/bin/bash

set -eu

trap "exit" INT TERM
trap "kill 0" EXIT

go build slow_app.go && ./slow_app &
go build app.go && ./app &

# wait for startup
sleep 5

wrk -t5 -c10 -d40s http://localhost:8080/blocking &
wrk -t5 -c10 -d40s http://localhost:8080/cpu-intensive &

# wait for warming up
sleep 5

pprof -http=":8088" app localhost:8080/debug/pprof/profile
