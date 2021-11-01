#!/bin/bash

CreateTable () {
    TableName=$1
    KeyName=$2

    aws --endpoint-url http://dynamodb:8000 dynamodb create-table \
    --table-name $TableName \
    --key-schema AttributeName=$KeyName,KeyType=HASH \
    --attribute-definitions AttributeName=$KeyName,AttributeType=S \
    --provisioned-throughput ReadCapacityUnits=3,WriteCapacityUnits=3
}

InitTableDatas () {
    file=$1
    aws --endpoint-url http://dynamodb:8000 dynamodb batch-write-item --request-items $file
}

## Create table
#### IG
CreateTable atlas_customs_tmp id
CreateTable atlas_customs_target id
CreateTable atlas_customs_account account
#### FB
# CreateTable atlas_customs_fb_tmp id
CreateTable atlas_customs_fb_target2 pageID
CreateTable atlas_customs_fb_account account
CreateTable atlas_fb_bot_account account

## Create default data of 'atlas_customs_target'
InitTableDatas file://default_ig_customs_target.json
InitTableDatas file://default_ig_customs_account.json

InitTableDatas file://default_fb_customs_target.json
InitTableDatas file://default_fb_customs_account.json
# InitTableDatas file://default_fb_bot_account.json
