#!/bin/sh
go build -o chat
if [ -e chat ]
then
./chat
else
echo "build failed."
fi
