# yaml [![GoDoc](https://godoc.org/github.com/erizocosmico/flaggax/yaml?status.svg)](https://godoc.org/github.com/erizocosmico/flagga)

YAML Source and Extractor for [flagga](https://github.com/erizocosmico/flagga).

## Install

```
go get -v github.com/erizocosmico/flaggax/yaml
```

## Usage

```go

import (
    "github.com/erizocosmico/flagga"
    "github.com/erizocosmico/flaggax/yaml"
)

//...

var fs flagga.FlagSet

db := fs.String("db", defaultDBURI, "database connection string", yaml.Key("db_uri"))
users := fs.StringList("users", nil, "list of allowed users", yaml.Key("users"))

err := fs.Parse(os.Args[1:], yaml.Via("config.yaml"))
if err != nil {
    // handle err
}
```