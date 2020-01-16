package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/lycanware/toolkit-go/filesys/copy"

	"github.com/urfave/cli"
)

var cliApp = cli.NewApp()

func main() {
	registerInfo()
	registerCommands()

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func registerInfo() {
	cliApp.Version = "1.0.0"
	cliApp.Name = "Lycanware Ant"
	cliApp.Usage = "A CLI to package software"
	cliApp.Authors = []cli.Author{{Name: "Lycanware", Email: "https://github.com/lycanware/ant"}}
}

func registerCommands() {
	cliApp.Commands = []cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build <settings_file>",
			Action:  cmdBuild,
		},
	}
}

func cmdBuild(c *cli.Context) error {
	var conf *Config
	var err error

	if !c.Args().Present() {
		printLog("Settings file missing")
		return nil
	}
	filePath := c.Args().First()

	// Resolve environment variables
	// This is needed if running the program from some IDEs
	r := regexp.MustCompile("$[A-Z_]+[A-Z0-9_]*")
	matches := r.FindAllStringSubmatch(filePath, -1)

	for _, match := range matches {
		filePath = strings.Replace(filePath, match[0], os.Getenv(match[0][1:]), -1)
	}

	// Load and unmarshall YML file
	if conf, err = NewConfig(filePath); err != nil {
		log.Fatal(err)
	}

	for _, i := range conf.Actions {
		if i.Action == "copy" {
			actionCopy(i)
		}
	}

	return nil
}

func actionCopy(i actionItem) {
	var err error
	var errList []error

	if i.DstClearFirst {
		if err = os.RemoveAll(i.Dst); err != nil {
			printLog(err)
			printLog(fmt.Printf("Skipped %#v copy because the destination %#v couldn't be cleared", i.Src, i.Dst))
			return
		}
	}

	if errList, err = copy.Dir(i.Src, i.Dst); err != nil {
		printLog(err)
		return
	}

	if len(errList) > 0 {
		printLog("Some files were not copied:")
		printLog(errList)
		return
	}

	printLog(fmt.Printf("Copied %#v - to - %#v", i.Src, i.Dst))
}

func printLog(a ...interface{}) {
	fmt.Println(a...)
}
