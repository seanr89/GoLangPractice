#!/bin/bash

#Script to execute and run a sqlserver linux version with name parameterised in!

echo "Executing script!"

#should be 1 argument, the name of the container
if [ $# -gt 0 ]; then
    echo "Your command line contains $# arguments"

    #1. Create Postgres Container
    containerName=$1;
    command=`docker run --rm --name $containerName -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres`
    echo $command
    #2. added sleep to ensure that the container is stable before the database is created
    sleep 2
    #3. Connect and create a sample database
    dbCreate=`docker exec -it $containerName psql -U postgres -c "CREATE DATABASE demo;"`
    echo $dbCreate

else
    echo "Your command line contains no arguments. Please provide the container name"
fi