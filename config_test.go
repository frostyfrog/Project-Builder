package main

import (
	"testing"
	"os"
	. "github.com/franela/goblin"
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
		g.It("Should load custom config file", func() {
			fo, err := os.Create("output.txt")
			if err != nil { g.Fail("Unable to open config file for writing") }
			defer fo.Close()
			fo.Write([]byte(`Projects
			- mice
			- golang
			- testing
			- gopher
			- goblin`))
			config := SystemConfig{}
			config.ReadFile("output.txt")
		})
		g.It("config file should have array of projects")
		g.It("should be able to load project config")
	})
}
