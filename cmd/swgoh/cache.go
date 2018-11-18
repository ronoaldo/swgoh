package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
)

func cacheFileName(kind string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	file := fmt.Sprintf("swgoh.%s.%s.json", allyCode, kind)
	return path.Join(usr.HomeDir, file), nil
}

func loadCache(kind string, dst interface{}) error {
	fname, err := cacheFileName(kind)
	if err != nil {
		return err
	}

	fd, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer fd.Close()

	return json.NewDecoder(fd).Decode(dst)
}

func saveCache(kind string, src interface{}) error {
	fname, err := cacheFileName(kind)
	if err != nil {
		return err
	}

	fd, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		return err
	}
	defer fd.Close()
	return json.NewEncoder(fd).Encode(src)
}
