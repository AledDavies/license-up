package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("license-up", "A command-line tool to make licences.")
	mit        = app.Command("mit", "Create MIT license.")
	mitName    = mit.Arg("name", "Name of license holder.").Required().String()
	mitSurname = mit.Arg("surname", "Surname of license holder.").Required().String()
	mitWebsite = mit.Arg("website", "Website of license holder").String()

	// TODO: add cc0 
	// cc0 = app.Command("cc0", "Create CC0 license.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case mit.FullCommand():
		mitCreateWithSite(string(*mitName), string(*mitSurname), string(*mitWebsite))
	}
}
