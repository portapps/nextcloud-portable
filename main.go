package main

//go:generate go tool goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest

import (
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/files"
	"github.com/portapps/portapps/v3/pkg/log"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *portapps.App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = portapps.NewWithCfg("nextcloud-portable", "Nextcloud", cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	confPath := filepath.Join(app.DataPath, "conf")
	for _, dir := range []string{
		confPath,
		filepath.Join(app.DataPath, "storage"),
		app.DataPath,
	} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			log.Fatal().Err(err).Msgf("Cannot create directory %s", dir)
		}
	}

	app.Process = filepath.Join(app.AppPath, "nextcloud.exe")
	app.Args = []string{
		"--confdir",
		confPath,
	}

	confFilePath := filepath.Join(confPath, "nextcloud.cfg")
	if _, err := os.Stat(confFilePath); err == nil {
		ini.PrettyFormat = false
		log.Info().Msg("Update configuration...")
		conf, err := ini.LoadSources(ini.LoadOptions{
			IgnoreInlineComment:         true,
			SkipUnrecognizableLines:     false,
			UnescapeValueDoubleQuotes:   true,
			UnescapeValueCommentSymbols: true,
			PreserveSurroundedQuote:     true,
			SpaceBeforeInlineComment:    true,
		}, confFilePath)
		if err == nil {
			conf.Section("General").Key("skipUpdateCheck").SetValue("true")
			if err := conf.SaveTo(confFilePath); err != nil {
				log.Error().Err(err).Msg("Write configuration")
			}
		} else {
			log.Error().Err(err).Msg("Load nextcloud.cfg file")
		}
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			files.Cleanup(filepath.Join(os.Getenv("LOCALAPPDATA"), "Nextcloud"))
		}()
	}

	defer app.Close()
	app.Launch(os.Args[1:])
}
