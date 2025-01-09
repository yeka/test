#/bin/sh
rustc $1 -o temp-bin
shift
./temp-bin $@
rm -f temp-bin