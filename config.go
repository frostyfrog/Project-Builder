package main

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"os/user"
	"path/filepath"
	//	"github.com/jonathankarsh/logstreamer"
)

const (
	git byte = iota
	svn byte = iota
)

type ProjectConfig struct {
	Project string
	Type    byte
	Scripts struct {
		Build   string
		Package string
	}
}

type SystemConfig struct {
	Map map[string][]interface{}
}

func (self *SystemConfig) Load() {
	usr, err := user.Current()
	checkErr(err)

	self.ReadFile(filepath.Join(usr.HomeDir, ".go-git-builder"))
}

// Read Configuration from a file
func (self *SystemConfig) ReadFile(path string) {
	data, err := ioutil.ReadFile(path)
	checkErr(err)

	err = yaml.Unmarshal(data, &self.Map)
	checkErr(err)
}
