package config

import (
	"os"
	"path"
)

var (
	// Setup file
	DefaultCaddyDir      string
	DefaultStatuspageDir string
	DefaultCheckupDir    string
)

func SetDefaultConfig() {
	var workingDir = "."
	workingDir, _ = os.Getwd()
	DefaultCaddyDir = path.Join(workingDir, "caddy_config")
	DefaultStatuspageDir = path.Join(workingDir, "statuspage")
	DefaultCheckupDir = path.Join(workingDir, "checkup_config")
}
