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

[winhttp]: https://github.com/mjwhitta/win/tree/main/winhttp
[wininet]: https://github.com/mjwhitta/win/tree/main/wininet

## How to install

Open a terminal and run the following:

```
$ go get --ldflags "-s -w" --trimpath -u github.com/mjwhitta/inet
```

## Usage

Minimal example:

```
package main

import (
    "bytes"
    "crypto/tls"
    "fmt"
    "io"
    "net/http"
    "net/http/cookiejar"

    "github.com/mjwhitta/inet"
)

var dst string = "https://icanhazip.com"

func configuredRequestAndDedicatedClientExample() error {
    var body io.Reader = bytes.NewReader([]byte("test"))
    var c inet.Client
    var e error
    var jar http.CookieJar
    var req *http.Request
    var res *http.Response

    // Create cookiejar
    if jar, e = cookiejar.New(nil); e != nil {
        panic(e)
    }

    // Create client
    if c, e = inet.NewClient(); e != nil {
        return e
    }

    // Set cookiejar
    c.Jar(jar)

    // Create request
    req, e = http.NewRequest(http.MethodPost, dst, body)
    if e != nil {
        return e
    }

    // Configure request
    req.AddCookie(&http.Cookie{Name: "chocolatechip", Value: "tasty"})
    req.AddCookie(&http.Cookie{Name: "oatmealraisin", Value: "gross"})
    req.AddCookie(&http.Cookie{Name: "snickerdoodle", Value: "yummy"})
    req.Header.Set("User-Agent", "testing, testing, 1, 2, 3...")

    // Send request
    if res, e = c.Do(req); e != nil {
        return e
    }
    defer res.Body.Close()

    return output(res)
}

func init() {
    // Disable TLS verification (WARNING: insecure)
    if t, ok := http.DefaultTransport.(*http.Transport); ok {
        if t.TLSClientConfig == nil {
            t.TLSClientConfig = &tls.Config{}
        }

        t.TLSClientConfig.InsecureSkipVerify = true
    }
}

func main() {
    if e := simpleGetExample(); e != nil {
        panic(e)
    }

    if e := configuredRequestAndDedicatedClientExample(); e != nil {
        panic(e)
    }
}

func output(res *http.Response) error {
    var b []byte
    var e error

    if b, e = io.ReadAll(res.Body); e != nil {
        return e
    }

    fmt.Println(res.Status)

    for k := range res.Header {
        fmt.Printf("%s: %s\n", k, res.Header.Get(k))
    }

    for _, cookie := range res.Cookies() {
        fmt.Printf("%s = %s\n", cookie.Name, cookie.Value)
    }

    if len(b) > 0 {
        fmt.Println(string(b))
    }

    fmt.Println()

    return nil
}

func simpleGetExample() error {
    var e error
    var res *http.Response

    // Use simple helper function
    if res, e = inet.Get(dst); e != nil {
        return e
    }
    defer res.Body.Close()

    return output(res)
}
```

## Links

- [Source](https://github.com/mjwhitta/inet)
