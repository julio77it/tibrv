# gotibrv
A CGO wrapper for [TIBCO](https://www.tibco.com/) [RendezVous C API](https://docs.tibco.com/pub/rendezvous/8.5.0/doc/html/wwhelp/wwhimpl/js/html/wwhelp.htm)

## Introduction

It's a CGO wrapper to C library. For building and using this package a valid TIBCO Rendezvous installation is needed.
This package permits to send and receive messages with regular, certified message delivery and distribuited queue transports.
The dispatcher library section is not included, use goroutines instead.

## Configuration

This file:

scripts/test_profile

references 2 enviroments variables :

**TEST_DIR**   = temporary test directory

**TIBRV_HOME** = path to valid tibrv C api installation

## Test & Build

### Test and coverage information:
scripts/cover.bash

### Benchmarks:
scripts/benchmark.bash

### Profiling:
scripts/profile.bash

### Examples and tools:
scripts/build.bash

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Disclaimer

This project has been an exercise for improving my GO skills, wrapping up things I already knew.

The package has never been used, it needs deep testing.

