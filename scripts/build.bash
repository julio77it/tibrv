#!/bin/bash

source test_profile

cd ..
go build -trimpath 
go build -trimpath examples/simple/simple.go
go build -trimpath examples/rvlisten/rvlisten.go
go build -trimpath examples/rvsend/rvsend.go
go build -trimpath examples/sendrequest/sendrequest.go
go build -trimpath examples/sendreply/sendreply.go
go build -trimpath examples/json/json.go

