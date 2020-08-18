#!/bin/bash

cur_num=$(ps -ef | grep "crawler_common" | grep -v "grep" | grep -v "check.sh" | wc -l)
crawler_num=1

if [ ${cur_num} != ${crawler_num} ];then
    ps aux | grep crawler_common | awk '{print $2}' | xargs -I {} kill -9 {}
    ps aux | grep route_common | awk '{print $2}' | xargs -I {} kill -9 {}
    for((i=1;i<=crawler_num;i++));
    do
        nohup sh -x route.sh &
    done
fi
