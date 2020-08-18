#! /bin/bash

cd /opt/crawler_common/bin

ulimit -c unlimited
export LD_LIBRARY_PATH=.:./
./crawler_common -c ../config/download.ini

