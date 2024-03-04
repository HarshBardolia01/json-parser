#!/bin/zsh

# removing older build
rm ./coverage-test/json-parser.test

# removing old coverage files
rm ./coverage-test/reports/counts/*
rm ./coverage-test/reports/functional/*

# compile
go test ./... -cover -c -o coverage-test

echo "Running Tests..."

for file in ./test/*.json; do
    fileName=${file##*/}
	coverFileName=${fileName%.*}".txt"
    
    ./coverage-test/json-parser.test -test.coverprofile=./coverage-test/reports/counts/$coverFileName $file
    go tool cover -func=./coverage-test/reports/counts/$coverFileName > ./coverage-test/reports/functional/$coverFileName
done

echo ""

