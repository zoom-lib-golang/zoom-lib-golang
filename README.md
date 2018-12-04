# Zoom.us Golang Client Library

[![Godoc](https://godoc.org/github.com/himalayan-institute/zoom-lib-golang?status.svg)](https://godoc.org/github.com/himalayan-institute/zoom-lib-golang)
[![Build Status](https://travis-ci.org/himalayan-institute/zoom-lib-golang.svg?branch=master)](https://travis-ci.org/himalayan-institute/zoom-lib-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/himalayan-institute/zoom-lib-golang)](https://goreportcard.com/report/github.com/himalayan-institute/zoom-lib-golang)
[![CodeClimate Maintainability](https://api.codeclimate.com/v1/badges/55b7484e20c0aaae35d7/maintainability)](https://codeclimate.com/github/himalayan-institute/zoom-lib-golang/maintainability)

Go (Golang) client library for the [Zoom.us REST API Version
2](https://zoom.github.io/api/). See
[here](https://gopkg.in/himalayan-institute/zoom-lib-golang.v1) for
Version 1 support.

For example use, see the Godoc documentation or the [examples
directory](_example/)

## v2 TODO

### Compatibility Notes

As with any major API update, there are some breaking changes:

* `GetUser` and `GetUserByEmail` have been combined into GetUser, which
  now has a signature more similar to GetUserByEmail
* User Login Type (`UserLoginType`) names have been upated to match Zoom
  documentation
* `ListRegistrationWebinars` has been combined with `ListWebinars`,
  which now gets everything

### Convert Endpoints

- [x] `GetUser`
- [x] `GetUserByEmail` *(deprecated)*
- [x] `GetWebinarInfo`
- [ ] `GetWebinarPanelist`
- [ ] `GetWebinarRegistrationInfo`
- [ ] `ListRegistrants`
- [x] `ListRegistrationWebinars` *(deprecated)*
- [x] `ListUsers`
- [x] `ListWebinars`
- [ ] `RegisterForWebinar`

## About

Built out of necessity, this repo will only support select endpoints at
first. Hopefully, it will eventually support all Zoom API endpoints.

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
