#!/bin/bash

source test_profile

cd ../tibrv
go test . -timeout 2s -coverprofile $TESTDIR/cover.out -trace $TESTDIR/trace.out -v $1
go tool cover -html=$TESTDIR/cover.out
