#!/bin/bash

source test_profile

cd ..
go test -run none -gcflags "-m=1" -bench . -benchtime 10s -benchmem -memprofile $TESTDIR/mem.out -cpuprofile $TESTDIR/cpu.out -trace=$TESTDIR/trace.out 2> $TESTDIR/benchmark.out
