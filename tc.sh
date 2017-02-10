#!/bin/sh

rm -rf bin
make build images

contrib/hack/stop-server.sh
contrib/hack/start-server.sh

go build -v -i  client.go && ./client  -v 10 -stderrthreshold 10 -logtostderr
