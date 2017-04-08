package commands

import (
	"html/template"
	"log"
	"os"

	"io/ioutil"

	"github.com/AdhityaRamadhanus/checkup"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// Main Function

func setupStatusPage(cliContext *cli.Context) {
	switch cliContext.String("storage") {
	case "fs":
		if err := setupFS(cliContext.String("url")); err != nil {
			log.Println(err)
		}
	}
}

// Helper
func setupFS(url string) error {
	// Create directory for logs
	// ignore the error
	log.Println("Setting up directory")
	os.Mkdir("./logs", 0777)
	os.Mkdir("./caddy-logs", 0777)
	os.Mkdir("./caddy-errors", 0777)
	// setup checkup.json
	log.Println("Creating checkup.json")
	checkup := checkup.Checkup{
		Checkers: []checkup.Checker{},
		Storage: checkup.FS{
			Dir: "/checkup/logs",
		},
	}
	jsonBytes, err := checkup.MarshalJSON()
	if err != nil {
		return err
	}
	file, err := os.OpenFile("./checkup.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		return errors.Wrap(err, "Failed opening checkup.json file")
	}
	if _, err = file.Write(jsonBytes); err != nil {
		return errors.Wrap(err, "Failed to write checkup.json")
	}
	// setup Caddyfile
	// Read the template
	log.Println("Creating Caddyfile")
	caddyTemplate, err := template.ParseFiles("./templates/Caddyfile")
	if err != nil {
		return errors.Wrap(err, "Error parsing caddy template config")
	}
	caddyFile, err := os.OpenFile("./Caddyfile", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	defer caddyFile.Close()
	if err != nil {
		return errors.Wrap(err, "Error opening Caddyfile")
	}
	// Execute the template
	if err := caddyTemplate.Execute(caddyFile, struct{ URL string }{URL: url}); err != nil {
		return errors.Wrap(err, "Failed writing to Caddyfile")
	}

	// setup config status page
	log.Println("Creating config.js for status page")
	srcConfigBytes, err := ioutil.ReadFile("./templates/config_fs.js")
	if err != nil {
		return errors.Wrap(err, "Failed opening config.js")
	}

	dstConfigFile, err := os.OpenFile("./statuspage/js/config.js", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	defer dstConfigFile.Close()
	if _, err = dstConfigFile.Write(srcConfigBytes); err != nil {
		return errors.Wrap(err, "Failed to write config.js")
	}

	// setup config status page
	log.Println("Creating index.html for status page")

	srcIndexHTML, err := ioutil.ReadFile("./templates/index_fs.html")
	if err != nil {
		return errors.Wrap(err, "Failed opening index_fs.html")
	}

	dstIndexHTML, err := os.OpenFile("./statuspage/index.html", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	defer dstConfigFile.Close()
	if _, err = dstIndexHTML.Write(srcIndexHTML); err != nil {
		return errors.Wrap(err, "Failed to write index.html")
	}
	log.Println("Success!")
	return nil
}
