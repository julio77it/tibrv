#!/bin/bash

source test_profile

cd ..
go test . -timeout 5s -coverprofile $TESTDIR/cover.out -trace $TESTDIR/trace.out -v $1
go tool cover -html=$TESTDIR/cover.out
