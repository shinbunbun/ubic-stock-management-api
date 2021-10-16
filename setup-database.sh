#!/bin/bash

aws dynamodb create-table \
    --table-name UBIC-FOOD-test \
    --attribute-definitions \
        AttributeName=ID,AttributeType=S \
        AttributeName=DataType,AttributeType=S \
        AttributeName=Data,AttributeType=S \
        AttributeName=DataKind,AttributeType=S \
    --key-schema \
        AttributeName=ID,KeyType=HASH \
    --global-secondary-indexes \
        IndexName=Data-DataType-index,KeySchema=[{'AttributeName=Data,KeyType=HASH'},{'AttributeName=DataType,KeyType=RANGE'}],ProvisionedThroughput='{ReadCapacityUnits=1,WriteCapacityUnits=1}',Projection={ProjectionType=ALL} \
        IndexName=DataKind-index,KeySchema=[{'AttributeName=DataKind,KeyType=HASH'}],ProvisionedThroughput='{ReadCapacityUnits=1,WriteCapacityUnits=1}',Projection={ProjectionType=ALL} \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --region ap-northeast-1 \
    --endpoint-url http://localhost:8000

aws dynamodb put-item --table-name UBIC-FOOD-test --item '{"ID": {"S":"1"},"DataType": {"S":"2"},"Data": {"S":"3"},"DataKind": {"S":"1"},"IntData": {"N":"1"}}' --region ap-northeast-1 --endpoint-url http://localhost:8000
