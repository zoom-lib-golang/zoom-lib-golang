# Zoom.us Golang Client Library (Zoom API v1)

**💥DEPRECATION NOTICE💥**

As of November 1, 2018, Zoom stopped supporting [Version
1](https://zoom.github.io/api-v1/) of their API. This branch is for
Version 1, and while you are welcome to use it for as long as Zoom
continues to support v1, we will no longer be adding new features. We
are still open to pull requests containing bug fixes, but we highly
recommend you update to v2 on the `master` branch as soon as possible.
We will do our best to maintain backwards compatibility.

To use the v1 version of the Zoom API, import as follows:

```golang
import "gopkg.in/himalayan-institute/zoom-lib-golang.v1"
```

----

[![Godoc](https://godoc.org/github.com/himalayan-institute/zoom-lib-golang?status.svg)](https://godoc.org/github.com/himalayan-institute/zoom-lib-golang)
[![Build Status](https://travis-ci.org/himalayan-institute/zoom-lib-golang.svg?branch=master)](https://travis-ci.org/himalayan-institute/zoom-lib-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/himalayan-institute/zoom-lib-golang)](https://goreportcard.com/report/github.com/himalayan-institute/zoom-lib-golang)

Go (Golang) client library for the [Zoom.us REST
API](git@github.com:himalayan-institute/zoom-lib-golang.git).

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
