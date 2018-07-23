package yaml

import (
	"github.com/erizocosmico/flagga"
	yaml "gopkg.in/yaml.v2"
)

// Key returns an Extractor that will match the given key in a provided
// YAML file to set as value for the flag.
func Key(key string) flagga.Extractor {
	return yamlExtractor(key)
}

type yamlExtractor string

func (e yamlExtractor) Get(sources []flagga.Source, dst flagga.Value) (bool, error) {
	for _, s := range sources {
		if _, ok := s.(*yamlSource); !ok {
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

// Via returns a Source that will use a YAML file as a provider of
// flag values.
func Via(file string) flagga.Source {
	return &yamlSource{flagga.NewFileSource(file, yaml.Unmarshal)}
}

type yamlSource struct {
	flagga.Source
}
