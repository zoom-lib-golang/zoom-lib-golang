# Zoom.us Golang Client Library

[![Godoc](https://godoc.org/github.com/zoom-lib-golang/zoom-lib-golang?status.svg)](https://godoc.org/github.com/zoom-lib-golang/zoom-lib-golang)
[![Build Status](https://travis-ci.org/zoom-lib-golang/zoom-lib-golang.svg?branch=master)](https://travis-ci.org/zoom-lib-golang/zoom-lib-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/zoom-lib-golang/zoom-lib-golang)](https://goreportcard.com/report/github.com/zoom-lib-golang/zoom-lib-golang)
[![CodeClimate Maintainability](https://api.codeclimate.com/v1/badges/55b7484e20c0aaae35d7/maintainability)](https://codeclimate.com/github/zoom-lib-golang/zoom-lib-golang/maintainability)

Go (Golang) client library for the [Zoom.us REST API Version
2](https://zoom.github.io/api/). See
[here](https://gopkg.in/zoom-lib-golang/zoom-lib-golang.v1) for
Version 1 support.

## About

Built out of necessity, this repo will only support select endpoints at
first. Hopefully, it will eventually support all Zoom API endpoints.

### Examples

For example use, see the Godoc documentation or the [examples
directory](_example/)

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

## Known Issues

- [ ] Calls to `/webinar/get` will return webinar occurrences that have
  been deleted with no indication of status (per [this
forum post](https://support.zoom.us/hc/en-us/community/posts/115010565986--webinar-get-returns-deleted-occurrence))
- [ ] Behavior of the `occurrence_ids` field in `/webinar/register` is
  unclear - see [this
forum post](https://support.zoom.us/hc/en-us/community/posts/115019165043-Behavior-of-occurrence-ids-in-webinar-register-?page=1#community_comment_115004843466)
for more details

## Contributing

Contributions welcome! Please see the [contributing
guidelines](CONTRIBUTING.md) for more details.

## Contact

For any questions regarding this library, please contact
[@rafecolton](https://github.com/rafecolton) or the Himalayan Institute
webteam at webteam@himalayaninstitute.org

Code inspired by
[mattbaird/gochimp](https://github.com/mattbaird/gochimp)
