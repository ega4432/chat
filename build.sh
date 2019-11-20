#!/bin/sh
go build -o chat
if [ -e chat ]
then
./chat -host=":8080"
else
echo "build failed."
fi
