


THRIFT=thrift
GO=go

all: hmstool hmsbench hmsweb

gen: hmsclient/thrift/gen-go/hive_metastore/hive_metastore.go

hmsclient/thrift/gen-go/hive_metastore/hive_metastore.go: hmsclient/thrift/hive_metastore.thrift
	$(THRIFT) --gen go:package_prefix=github.com/SergeAlexandre/gometastore/hmsclient/thrift/gen-go/$(COMPILER_EXTRAFLAG) -out hmsclient/thrift/gen-go -r $<


hmstool: gen
	cd hmstool; go install

hmsbench: gen
	cd hmsbench; go install

hmsweb: gen
	cd hmsweb; go install
