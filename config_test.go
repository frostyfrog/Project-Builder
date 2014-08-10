package main

import (
	"testing"
	"os"
	. "github.com/franela/goblin"
	//"fmt"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)
	setupLoggers()
	g.Describe("Config", func() {
		g.It("Should fail when file doesn't exit", func() {
			defer func() {
				if r := recover(); r == nil {
					g.Fail("Config Loader failded to panic")
				}
			}()
			config := SystemConfig{}
			config.Load()
			g.Assert(1 + 1).Equal(2)
		})
		g.Before(func(){
			fo, err := os.Create("test_config.conf")
			if err != nil { g.Fail("Unable to open config file for writing") }
			defer fo.Close()
			fo.Write([]byte(`Projects:
 - mice
 - golang
 - testing
 - gopher
 - goblin`))
			fo, err = os.Create("test_project_config.conf")
			if err != nil { g.Fail("Unable to open config file for writing") }
			defer fo.Close()
			fo.Write([]byte(`Project: TestProj
Type: git
URL: http://github.com/frostyfrog/Project-Builder
Scripts:
 Build:
  - pkgbuild
 Package:
  - goblin.sh`))
		})
		g.After(func(){
			os.Remove("test_config.conf")
		})
		g.It("Should load custom config file", func() {
			config := SystemConfig{}
			config.ReadFile("test_config.conf")
		})
		g.It("config file should have array of projects", func() {
			config := SystemConfig{}
			config.ReadFile("test_config.conf")
			//for k, v := range config.Map["Projects"] {
			//	fmt.Println(k, v)
			//}
			g.Assert(len(config.Map["Projects"]) == 5).IsTrue()
		})
		g.It("should be able to load project config")
	})
}
/*
This is "great" testing debug code :3
f, err := os.Create("tmp")
if err != nil { g.Fail("Unable to open config file for writing") }
defer f.Close()
f.WriteString(fmt.Sprintf("%v", config))
*/
