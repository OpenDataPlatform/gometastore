module github.com/SergeAlexandre/gometastore/hmsbench

go 1.15

replace github.com/SergeAlexandre/gometastore/hmsclient v0.1.0 => ../hmsclient

replace github.com/SergeAlexandre/gometastore/microbench v0.1.0 => ../microbench

require (
	github.com/SergeAlexandre/gometastore/hmsclient v0.1.0
	github.com/SergeAlexandre/gometastore/microbench v0.1.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.8.1
)
