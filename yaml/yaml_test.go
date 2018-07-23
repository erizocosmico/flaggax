package yaml

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/erizocosmico/flagga"
	yaml "gopkg.in/yaml.v2"
)

func TestKey(t *testing.T) {
	testCases := []struct {
		key      string
		err      bool
		ok       bool
		value    interface{}
		expected interface{}
	}{
		{"foo", false, false, nil, nil},
		{"bar", false, true, new(int64), int64(42)},
		{"baz", true, false, new(bool), nil},
	}

	sources := []flagga.Source{
		&yamlSource{&flagga.FileSource{Value: map[string]interface{}{
			"bar": int64(42),
			"baz": float64(3.14),
		}}},
	}

	for _, tt := range testCases {
		t.Run(tt.key, func(t *testing.T) {
			ok, err := Key(tt.key).Get(sources, flagga.NewValue(tt.value))
			if tt.err && err == nil {
				t.Errorf("expecting error, got nil instead")
			} else if !tt.err && err != nil {
				t.Errorf("got unexpected error: %s", err)
			}

			if tt.ok != ok {
				t.Errorf("expected ok to be: %v, got: %v", tt.ok, ok)
			}

			if tt.ok {
				val := reflect.ValueOf(tt.value).Elem().Interface()
				if !reflect.DeepEqual(val, tt.expected) {
					t.Errorf("expecting value to be: %v, got: %v", tt.expected, val)
				}
			}
		})
	}
}

func TestVia(t *testing.T) {
	data, err := yaml.Marshal(map[string]interface{}{
		"foo": "bar",
		"bar": 1,
		"baz": []interface{}{3, 1, "5"},
	})
	if err != nil {
		t.Fatalf("unexpected error encoding yaml: %s", err)
	}

	f, err := ioutil.TempFile(os.TempDir(), "yaml-test-flagga")
	if err != nil {
		t.Fatalf("unexpected error saving yaml file: %s", err)
	}

	defer func() {
		if err := os.Remove(f.Name()); err != nil {
			t.Errorf("error removing file: %s", err)
		}
	}()

	if _, err := io.Copy(f, bytes.NewBuffer(data)); err != nil {
		t.Fatalf("unexpected error copying yaml: %s", err)
	}

	if err := f.Close(); err != nil {
		t.Errorf("error closing file: %s", err)
	}

	source := Via(f.Name())
	if err := source.Open(); err != nil {
		t.Fatalf("unable to open json file: %s", err)
	}

	testCases := []struct {
		dst      interface{}
		key      string
		expected interface{}
		err      bool
		ok       bool
	}{
		{new(string), "qux", nil, false, false},
		{new(string), "foo", "bar", false, true},
		{new(int), "foo", nil, true, false},
		{new([]int), "baz", []int{3, 1, 5}, false, true},
	}

	for _, tt := range testCases {
		t.Run(tt.key, func(t *testing.T) {
			ok, err := source.Get(tt.key, flagga.NewValue(tt.dst))
			if tt.err && err == nil {
				t.Errorf("expecting error, got nil instead")
			} else if !tt.err && err != nil {
				t.Errorf("got unexpected error: %s", err)
			}

			if tt.ok != ok {
				t.Errorf("expected ok to be: %v, got: %v", tt.ok, ok)
			}

			if tt.ok {
				val := reflect.ValueOf(tt.dst).Elem().Interface()
				if !reflect.DeepEqual(val, tt.expected) {
					t.Errorf("expecting value to be: %v, got: %v", tt.expected, val)
				}
			}
		})
	}
}
