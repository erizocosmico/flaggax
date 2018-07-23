# toml [![GoDoc](https://godoc.org/github.com/erizocosmico/flaggax/toml?status.svg)](https://godoc.org/github.com/erizocosmico/flagga)

TOML Source and Extractor for [flagga](https://github.com/erizocosmico/flagga).

## Install

```
go get -v github.com/erizocosmico/flaggax/toml
```

## Usage

```go

import (
    "github.com/erizocosmico/flagga"
    "github.com/erizocosmico/flaggax/toml"
)

//...

var fs flagga.FlagSet

db := fs.String("db", defaultDBURI, "database connection string", toml.Key("db_uri"))
users := fs.StringList("users", nil, "list of allowed users", toml.Key("users"))

err := fs.Parse(os.Args[1:], toml.Via("config.toml"))
if err != nil {
    // handle err
}
```