package config

import (
	"os"
	"path"
)

var (
	// Setup file
	DefaultCheckupJSON string
	DefaultConfigJS    string
	DefaultCaddyFile   string
	DefaultIndexHtml   string

	// Template file
	DefaultTplFSJS      string
	DefaultTplS3JS      string
	DefaultTplIndexFS   string
	DefaultTplIndexS3   string
	DefaultTplCaddyfile string
)

func SetDefaultConfig() {
	var workingDir = "."
	workingDir, _ = os.Getwd()
	DefaultConfigJS = path.Join(workingDir, "statuspage/js/config.js")
	DefaultCheckupJSON = path.Join(workingDir, "checkup_config/checkup.json")
	DefaultCaddyFile = path.Join(workingDir, "caddy_config/Caddyfile")
	DefaultIndexHtml = path.Join(workingDir, "statuspage/index.html")

	// setup template path default
	DefaultTplFSJS = path.Join(workingDir, "templates/config_fs.js")
	DefaultTplS3JS = path.Join(workingDir, "templates/config_s3.js")
	DefaultTplIndexFS = path.Join(workingDir, "templates/index_fs.html")
	DefaultTplIndexS3 = path.Join(workingDir, "templates/index_s3.html")
	DefaultTplCaddyfile = path.Join(workingDir, "templates/Caddyfile")
}
