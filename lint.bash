#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# gometalinter

if [ ! $(command -v gometalinter) ]
then
	go get -u github.com/alecthomas/gometalinter
	gometalinter --update --install
fi

echo "gometalinter:"
time gometalinter \
	--exclude='error return value not checked.*(Close|Log|Print).*\(errcheck\)$' \
	--exclude='.*_test\.go:.*error return value not checked.*\(errcheck\)$' \
	--exclude='/thrift/' \
	--exclude='/pb/' \
	--exclude='no args in Log call \(vet\)' \
	--disable=dupl \
	--disable=aligncheck \
	--disable=gotype \
	--cyclo-over=20 \
	--tests \
	--concurrency=2 \
	--deadline=300s \
	./...
echo 

# golangci-lint

if [ ! $(command -v golangci-lint) ]
then
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
fi


echo "golangci-lint:"
time golangci-lint \
	run \
	./... 