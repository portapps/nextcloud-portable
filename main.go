//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"path"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = NewWithCfg("nextcloud-portable", "Nextcloud", cfg); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	confPath := utl.CreateFolder(app.DataPath, "conf")
	utl.CreateFolder(app.DataPath, "storage")

	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "nextcloud.exe")
	app.Args = []string{
		"--confdir",
		confPath,
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			utl.Cleanup([]string{
				path.Join(os.Getenv("LOCALAPPDATA"), "Nextcloud"),
			})
		}()
	}

	app.Launch(os.Args[1:])
}
