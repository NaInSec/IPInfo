#!/bin/bash

# Lint all files in the project.

golint                                                                        \
    ./lib/...                                                                 \
    ./ipinfo/...                                                              \
    ./grepip/...                                                              \
    ./grepdomain/...                                                          \
    ./matchip/...                                                             \
    ./prips/...                                                               \
    ./cidr2range/...                                                          \
    ./range2cidr/...                                                          \
    ./range2ip/...                                                            \
    ./cidr2ip/...                                                             \
    ./splitcidr/...                                                           \
    ./randip/...
