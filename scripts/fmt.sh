#!/bin/bash

DIR=`dirname $0`
ROOT=$DIR/..

# Format code in project.

gofmt -w                                                                      \
    $ROOT/lib                                                                 \
    $ROOT/ipinfo                                                              \
    $ROOT/grepip                                                              \
    $ROOT/grepdomain                                                          \
    $ROOT/matchip                                                             \
    $ROOT/prips                                                               \
    $ROOT/cidr2range                                                          \
    $ROOT/range2cidr                                                          \
    $ROOT/range2ip                                                            \
    $ROOT/cidr2ip                                                             \
    $ROOT/splitcidr                                                           \
    $ROOT/randip
