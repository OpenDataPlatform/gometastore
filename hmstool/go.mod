module github.com/SergeAlexandre/gometastore/hmstool

go 1.15

replace github.com/SergeAlexandre/gometastore/hmsclient v0.1.0 => ../hmsclient

require (
	git.apache.org/thrift.git v0.14.2 // indirect
	github.com/SergeAlexandre/gometastore/hmsclient v0.1.0
	github.com/colinmarc/hdfs v1.1.3
	github.com/gobwas/glob v0.2.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.8.1
)
