package config

import "github.com/gobuffalo/envy"

const HTTPBIN = "https://httpbin.org"

var ANOTHERURL = envy.Get("HTTPBIN_URL", "https://httpbin.org")

var TestDataMode = true