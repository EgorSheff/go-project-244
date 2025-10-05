package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	ErrUsupportedFormat = errors.New("unsupported format")
)

func ParseFile(path string) (map[string]any, error) {
	ext := filepath.Ext(path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error read file %s: %v", path, err)
	}

	var res map[string]any

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &res); err != nil {
			return nil, fmt.Errorf("error parse json file: %v", err)
		}
	case ".yaml":
		if err = yaml.Unmarshal(data, &res); err != nil {
			return nil, fmt.Errorf("error parse yaml file: %v", err)
		}
	default:
		return nil, ErrUsupportedFormat
	}

	return res, nil
}
