#!/bin/bash

#Script to execute and run a sqlserver linux version with name parameterised in!

echo "Executing script!"

command=`docker run --rm -d --hostname my-rabbit -p 5672:5672 --name some-rabbit rabbitmq:3`
echo $command
sleep 5
commandManager=`docker run --rm -d --hostname my-rabbit -p 15672:15672 --name rabbit-man rabbitmq:3-management`
echo $commandManager