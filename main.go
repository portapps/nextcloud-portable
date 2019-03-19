//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("nextcloud-portable", "Nextcloud"); err != nil {
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

	app.Launch(os.Args[1:])
}
