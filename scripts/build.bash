#!/bin/bash

source test_profile

cd ..
go build -a -trimpath 
go build -a -trimpath examples/simple/simple.go
go build -a -trimpath examples/rvlisten/rvlisten.go
go build -a -trimpath examples/rvsend/rvsend.go
go build -a -trimpath examples/sendrequest/sendrequest.go
go build -a -trimpath examples/sendreply/sendreply.go
go build -a -trimpath examples/json/json.go
