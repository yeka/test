#/bin/bash

g++ -std=c++11 $1 -o $1.out && \
./$1.out && \
rm -f $1.out
