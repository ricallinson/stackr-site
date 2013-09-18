# Stackr [![Build Status](https://secure.travis-ci.org/ricallinson/stackr.png?branch=master)](http://travis-ci.org/ricallinson/stackr)

__WARNING: UNSTABLE API__

Stackr is an extensible HTTP server framework for Go, shipping with over 2 bundled middleware and a poor selection of 3rd-party middleware.

* [GoDoc](http://godoc.org/github.com/ricallinson/stackr)
* [GitHub](https://github.com/ricallinson/stackr)

## Install

    go get github.com/ricallinson/stackr

## Write Code

    package main

    import "github.com/ricallinson/stackr"

    func main() {
        app := stackr.CreateServer()
        app.Use("/", stackr.Logger())
        app.Use("/", stackr.Static())
        app.Use("/", func(req *stackr.Request, res *stackr.Response, next func()) {
            res.End("hello world\n")
        })
        app.Listen(3000)
    }

## Use Middleware

* `Logger` request logger with custom format support
* `Favicon` efficient favicon server
* `Static` static file server currently based on http.FileServer

## Make Middleware

    func MyMiddleware() (func(req *Request, res *Response, next func())) {
        // Prep
        return func(req *Request, res *Response, next func()) {
            // Process
        })
    })