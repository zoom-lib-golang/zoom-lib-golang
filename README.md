# Zoom.us Golang Client Library

[![Build Status](https://travis-ci.org/himalayan-institute/zoom-lib-golang.svg?branch=master)](https://travis-ci.org/himalayan-institute/zoom-lib-golang)
[![Godoc](https://godoc.org/github.com/himalayan-institute/zoom-lib-golang?status.svg)](https://godoc.org/github.com/himalayan-institute/zoom-lib-golang)

Go (Golang) client library for the [Zoom.us REST
API](git@github.com:himalayan-institute/zoom-lib-golang.git).

## About

Built out of necessity, this repo will only support select endpoints at
first. Hopefully, it will eventually support all Zoom API endpoints.

## Known Issues

- [ ] Calls to `/webinar/get` will return webinar occurrences that have
  been deleted with no indication of status (per [this
ticket](https://support.zoom.us/hc/en-us/community/posts/115010565986--webinar-get-returns-deleted-occurrence)

## Contributing

Contributions welcome! Please see the [contributing
guidelines](CONTRIBUTING.md) for more details.

## Contact

For any questions regarding this library, please contact
[@rafecolton](https://github.com/rafecolton) or the Himalayan Institute
webteam at webteam@himalayaninstitute.org

Code inspired by
[mattbaird/gochimp](https://github.com/mattbaird/gochimp)
