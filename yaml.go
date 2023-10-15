package main

import "gopkg.in/yaml.v2"

func ReadTextsData(data []byte) ([]TextMeta, error) {
	texts := []TextMeta{}
	err := yaml.Unmarshal(data, &texts)

	if err != nil {
		return nil, err
	}

	return texts, nil
}
