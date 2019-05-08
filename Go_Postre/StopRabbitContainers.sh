#!/bin/bash

#Script to attempt to stop all rabbit containers

echo "Executing script!"

stoprabbit=`docker stop some-rabbit`
echo $stoprabbit
sleep 2
stopmanager=`docker stop rabbit-man`
echo $stopmanager