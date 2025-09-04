@echo off

pushd ..

if not exist bin mkdir bin

go build -o bin ./cmd/pauvm
go build -o bin ./cmd/pauven
go build -o bin ./cmd/paudiss

popd
