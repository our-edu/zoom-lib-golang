# Zoom.us Golang Client Library


## This package is forked originally from [zoom-lib-golang](https://github.com/zoom-lib-golang/zoom-lib-golang) because of problem of upgrading to v2 in their tagging
Go (Golang) client library for the [Zoom.us REST API Version
2](https://zoom.github.io/api/). See
[here](https://gopkg.in/zoom-lib-golang/zoom-lib-golang.v1) for
Version 1 support.

## About

Built out of necessity, this repo will only support select endpoints at
first. Hopefully, it will eventually support all Zoom API endpoints.

### Tests

To run unit tests and the linter:

```bash
./fmtpolice
go test -v ./...
```

To run the integration tests:

```bash
# first, define the required environment variables
export ZOOM_API_KEY="<key>"
export ZOOM_API_SECRET="<secret>"
export ZOOM_EXAMPLE_EMAIL="<account email>"

# then run the tests with the integration build tag
go test -tags integration -v ./...
```