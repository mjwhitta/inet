# INet

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/inet?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/inet)
![License](https://img.shields.io/github/license/mjwhitta/inet?style=for-the-badge)

## What is this?

This module is a simple wrapper around `net/http`, [winhttp], and
[wininet]. It creates a minimal, cross-platform HTTP `Client`
interface and handles choosing an appropriate backend. On Windows you
can use any of the three supported backends (defaults to `wininet`).
On Linux the only supported backend is `net/http`.

## How to install

Open a terminal and run the following:

```
$ go get -u github.com/mjwhitta/inet
```

## Usage

Minimal example:

```
//nolint:errcheck,gosec,wrapcheck // This is merely an example
package main

import (
    "bytes"
    "crypto/tls"
    "flag"
    "io"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "strings"

    "github.com/mjwhitta/inet"
)

var (
    debug bool
    dst   string = "https://icanhazip.com"
)

func configuredRequestAndDedicatedClientExample() error {
    var b []byte
    var body io.Reader = bytes.NewReader([]byte("what is my ip?"))
    var c inet.Client
    var e error
    var jar http.CookieJar
    var req *http.Request
    var res *http.Response

    println("### Advanced usage ###")

    // Create cookiejar
    if jar, e = cookiejar.New(nil); e != nil {
        return e
    }

    // Create client
    if c, e = inet.NewClient("My custom UA"); e != nil {
        return e
    }

    // Enable or disable debug functionality
    c.Debug(debug) // Disabled by default

    // Set cookiejar
    c.SetJar(jar)

    if e = injectCookiesForTesting(c.Jar()); e != nil {
        return e
    }

    // Create request
    req, e = http.NewRequest(http.MethodPost, dst, body)
    if e != nil {
        return e
    }

    // Send request
    if res, e = c.Do(req); e != nil {
        return e
    }
    defer res.Body.Close()

    if b, e = io.ReadAll(res.Body); e != nil {
        return e
    }

    if !debug {
        println(strings.TrimSpace(string(b)))
    }

    return nil
}

func init() {
    flag.BoolVar(&debug, "debug", false, "Debug output")
    flag.Parse()

    // Disable TLS verification (WARNING: insecure)
    if t, ok := http.DefaultTransport.(*http.Transport); ok {
        if t.TLSClientConfig == nil {
            t.TLSClientConfig = &tls.Config{}
        }

        t.TLSClientConfig.InsecureSkipVerify = true
    }
}

func injectCookiesForTesting(jar http.CookieJar) error {
    var cookies []*http.Cookie = []*http.Cookie{
        {Name: "chocolatechip", Value: "yum"},
        {Name: "oatmealraisin", Value: "gross"},
        {Name: "snickerdoodle", Value: "best"},
    }
    var e error
    var uri *url.URL

    if uri, e = url.Parse(dst); e != nil {
        return e
    }

    jar.SetCookies(uri, cookies)

    return nil
}

func main() {
    if e := simpleGetExample(); e != nil {
        panic(e)
    }

    if e := configuredRequestAndDedicatedClientExample(); e != nil {
        panic(e)
    }
}

func simpleGetExample() error {
    var b []byte
    var e error
    var res *http.Response

    println("### Basic usage ###")

    // Use simple helper function
    if res, e = inet.Get(dst); e != nil {
        return e
    }
    defer res.Body.Close()

    if b, e = io.ReadAll(res.Body); e != nil {
        return e
    }

    println(strings.TrimSpace(string(b)))

    return nil
}
```

## Links

- [Source](https://github.com/mjwhitta/inet)

## TODO

- Unit tests

[winhttp]: https://github.com/mjwhitta/win/tree/main/winhttp
[wininet]: https://github.com/mjwhitta/win/tree/main/wininet
