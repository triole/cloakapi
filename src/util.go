package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
	"go.yaml.in/yaml/v3"
)

func deref(inp *string) (r string) {
	if inp != nil {
		r = *inp
	}
	return
}

func makeAbs(filename string) (r string, err error) {
	return filepath.Abs(filename)
}

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func pprintJSON(i any) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}

func pprintTOML(i any) {
	s, _ := toml.Marshal(i)
	fmt.Println(string(s))
}

func pprintYAML(i any) {
	s, _ := yaml.Marshal(i)
	fmt.Println(string(s))
}

func fmtYAML(i any) string {
	s, _ := yaml.Marshal(i)
	return string(s)
}
