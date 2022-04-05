#!/bin/bash

source test_profile

cd .. 

# G103 rules excluded because unsafe package is needed
gosec ./...