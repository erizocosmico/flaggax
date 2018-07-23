package toml

import (
	"github.com/BurntSushi/toml"
	"github.com/erizocosmico/flagga"
)

// Key returns an Extractor that will match the given key in a provided
// TOML file to set as value for the flag.
func Key(key string) flagga.Extractor {
	return tomlExtractor(key)
}

type tomlExtractor string

func (e tomlExtractor) Get(sources []flagga.Source, dst flagga.Value) (bool, error) {
	for _, s := range sources {
		if _, ok := s.(*tomlSource); !ok {
			continue
		}

		ok, err := s.Get(string(e), dst)
		if err != nil {
			return false, err
		}

		if !ok {
			continue
		}

		return true, nil
	}

	return false, nil
}

type tomlSource struct {
	flagga.Source
}

// Via returns a Source that will use a TOML file as a provider of
// flag values.
func Via(file string) flagga.Source {
	return &tomlSource{flagga.NewFileSource(file, toml.Unmarshal)}
}
