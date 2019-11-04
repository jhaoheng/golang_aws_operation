#!/bin/bash

CreateTable () {
    TableName=$1
    KeyName=$2

    awscli --endpoint-url http://dynamodb:8000 dynamodb create-table \
    --table-name $TableName \
    --key-schema AttributeName=$KeyName,KeyType=HASH \
    --attribute-definitions AttributeName=$KeyName,AttributeType=S \
    --provisioned-throughput ReadCapacityUnits=3,WriteCapacityUnits=3
}

InitTableDatas () {
    file=$1

    awscli --endpoint-url http://dynamodb:8000 dynamodb batch-write-item --request-items $file
}

## Create table
CreateTable test id

## Create default data of 'atlas_customs_target'
InitTableDatas file://default_datas.json



## 
/bin/bash