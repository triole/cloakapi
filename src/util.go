package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	toml "github.com/pelletier/go-toml/v2"
	"go.yaml.in/yaml/v3"
)

// keep-sorted start block=yes newline_separated=yes
func (kc tKC) readFile(pth string) (by []byte) {
	by, err := os.ReadFile(pth)
	kc.Lg.IfErrFatal("can not read file: %q, %q", pth, err)
	return
}

func deref(inp *string) (r string) {
	if inp != nil {
		r = *inp
	}
	return
}

func find(basedir string, rxFilter string) []string {
	_, err := os.Stat(basedir)
	// if err != nil {
	// 	fmt.Printf("can not access folder %q\n", err)
	// 	os.Exit(1)
	// }
	filelist := []string{}
	rxf, _ := regexp.Compile(rxFilter)

	err = filepath.Walk(basedir, func(path string, f os.FileInfo, err error) error {
		if rxf.MatchString(path) {
			inf, err := os.Stat(path)
			if err == nil {
				if !inf.IsDir() {
					filelist = append(filelist, path)
				}
			} else {
				print("stat file failed %q", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("unable to detect files %q\n", err)
		os.Exit(1)
	}
	return filelist
}

func fmtYAML(i any) string {
	s, _ := yaml.Marshal(i)
	return string(s)
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

// keep-sorted end
