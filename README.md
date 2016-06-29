gluahttpscrape
==============
[![Build Status](https://travis-ci.org/felipejfc/gluahttpscrape.svg?branch=master)](https://travis-ci.org/felipejfc/gluahttpscrape)
[![Coverage Status](https://coveralls.io/repos/github/felipejfc/gluahttpscrape/badge.svg?branch=master)](https://coveralls.io/github/felipejfc/gluahttpscrape?branch=master)

A html parser module for [gopher-lua](https://github.com/yuin/gopher-lua)

### Usage

```
L.PreloadModule("scrape", gluahttpscrape.NewHttpScrapeModule().Loader)
```

See tests for usage of module inside lua scripts.
