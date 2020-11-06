#!/bin/bash

source test_profile

cd ..
go build -ldflags "-extldflags '$CGO_FLAGS'" -a -trimpath 
go build -ldflags "-extldflags '$CGO_FLAGS'" -a -trimpath examples/simple/simple.go
go build -ldflags "-extldflags '$CGO_FLAGS'" -a -trimpath examples/rvlisten/rvlisten.go
go build -ldflags "-extldflags '$CGO_FLAGS'" -a -trimpath examples/rvsend/rvsend.go
go build -ldflags "-extldflags '$CGO_FLAGS'" -a -trimpath examples/sendrequest/sendrequest.go
go build -ldflags "-extldflags '$CGO_FLAGS'" -a -trimpath examples/sendreply/sendreply.go
go build -ldflags "-extldflags '$CGO_FLAGS'" -a -trimpath examples/json/json.go
