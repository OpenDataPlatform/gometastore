# HMSTool - Hive Metastore CLI tool

## Building

    go get -u github.com/akolb1/gometastore/hmstool

## Linux build

    env GOOS=linux GOARCH=amd64 go build -o hmstool-linux-amd64


The tool uses documentation in the code to generate extarnal docs. So if you make any changes,
please update docs as well with

    hmstool doc ./doc
    
Then commit changes to `doc` directory and push to github.    

[Documentation][]

[Documentation]: https://github.com/akolb1/gometastore/blob/master/hmstool/doc/hmstool.md "Documentation"

