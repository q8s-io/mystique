#!/usr/bin/env bash

ulimit -c unlimited
export LD_LIBRARY_PATH=.:./

nohup ./crawler_common -c ../config/download.ini | ./route_common "./route.config" & 
