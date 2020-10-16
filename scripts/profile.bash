#!/bin/bash

source test_profile

go tool pprof -alloc_space $TESTDIR/mem.out
