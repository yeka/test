#/bin/sh
rustc $1 -o temp-bin
./temp-bin
rm -f temp-bin