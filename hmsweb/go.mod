module github.com/SergeAlexandre/gometastore/hmsweb

go 1.15

replace github.com/SergeAlexandre/gometastore/hmsclient v0.1.0 => ../hmsclient

require (
	github.com/SergeAlexandre/gometastore/hmsclient v0.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/oklog/ulid v1.3.1
)
