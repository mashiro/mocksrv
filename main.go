package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
)

var version = "master"

// Options is a server options
type Options struct {
	Version bool   `long:"version" description:"Show version"`
	Address string `short:"a" long:"address" description:"Server address" default:":3000"`
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

// ModuleFunc is module callback function
type ModuleFunc func(*gin.Engine)

var modules []ModuleFunc

// RegisterModule register new module
func RegisterModule(desc string, data interface{}, mod ModuleFunc) {
	parser.AddGroup(desc, "", data)
	modules = append(modules, mod)
}

func main() {
	_, err := parser.Parse()
	if err != nil {
		ferr, ok := err.(*flags.Error)
		if ok && ferr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if options.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	gin.SetMode(gin.ReleaseMode)
	if len(options.Verbose) >= 2 {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()

	if len(options.Verbose) >= 1 {
		engine.Use(gin.Logger())
	}
	engine.Use(gin.Recovery())

	for _, mod := range modules {
		mod(engine)
	}

	engine.Run(options.Address)
}
