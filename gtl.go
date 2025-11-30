package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	"codeberg.org/fab/gtl/config"
	"codeberg.org/fab/gtl/core"
	"codeberg.org/fab/gtl/ui"
)

var Data = core.TlData{}
var Version string = "devel"

func main() {
	var configArg, modeArg, urlArg, titleArg string
	var helpArg, versionArg bool
	var cliLimitArg int

	flag.StringVar(&configArg, "config", "", "The path to gtl.toml config file.")
	flag.StringVar(&modeArg, "mode", "", "The mode for gtl, either 'cli', 'tui' or 'gemini'.")
	flag.StringVar(&urlArg, "url", "", "The tinylog url you want to (un)subscribe to.")
	flag.StringVar(&titleArg, "title", "", "The optional title of the tinylog you want to subscribe to.")
	flag.IntVar(&cliLimitArg, "limit", 0, "Limit number of items in CLI mode.")
	flag.BoolVar(&helpArg, "help", false, "Display help.")
	flag.BoolVar(&versionArg, "version", false, "Display gtl's current version.")
	flag.Parse()

	if helpArg == true {
		ui.Help()
		os.Exit(0)
	}

	if versionArg == true {
		fmt.Printf("gtl version: %s\n", Version)
		os.Exit(0)
	}

	// Init configuration and subscriptions.
	config.Init(configArg, &Data)

	if len(os.Args) > 1 {
		switch os.Args[0] {
		case "add":
			if urlArg == "" {
				ui.SubHelp()
				log.Fatalln("No url to add!")
			}
			e := config.AddTlSubscription(Data.Config.Subscribed_data, urlArg, titleArg)
			if e != nil {
				fmt.Println(e)
				log.Fatalln("Error adding entry.")
			} else {
				fmt.Println(urlArg, "has been added to the subscribed list.")
			}
			os.Exit(0)
		case "rm":
			if urlArg == "" {
				ui.SubHelp()
				log.Fatalln("No url to remove")
			}
			e := config.RemoveTlSubscription(Data.Config.Subscribed_data, urlArg)
			if e != nil {
				fmt.Println(e)
				log.Fatalln("Error removing entry:")
			} else {
				fmt.Println(urlArg, "has been removed from the subscribed list.")
			}
			os.Exit(0)
		}
	}

	// Display stream and quit.
	var gtlUI ui.TlUI
	if modeArg == "tui" || modeArg == "cli" || modeArg == "gemini" {
		gtlUI.Mode = modeArg
	} else if Data.Config.Mode == "tui" || Data.Config.Mode == "cli" || Data.Config.Mode == "gemini" {
		gtlUI.Mode = Data.Config.Mode
	} else {
		fmt.Printf("Unknown mode, fallback on TUI.")
		gtlUI.Mode = "tui"
	}

	if e := gtlUI.Run(&Data, cliLimitArg); e != nil {
		fmt.Println(e)
		log.Fatalln(e)
	}
}
