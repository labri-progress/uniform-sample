#!/bin/bash
echo "Experiment "
times=10
p="data/resultJul95"
#p="data/resultss_log.txt" #"data/resultaccess_log"
#p="data/resultess_1"
for i in 2
do
    for j in {1..2}
    do
        ./main -t=$i -n=$j -p=$p
    done
done

