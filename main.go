//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "nextcloud-portable"
	Papp.Name = "Nextcloud"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	confPath := CreateFolder(PathJoin(Papp.DataPath, "conf"))
	CreateFolder(PathJoin(Papp.DataPath, "storage"))

	Papp.Process = PathJoin(Papp.AppPath, "nextcloud.exe")
	Papp.Args = []string{
		"--confdir",
		confPath,
	}
	Papp.WorkingDir = Papp.AppPath

	// Update nextcloud settings
	/*nextcloudSettingsPath := PathJoin(confPath, "nextcloud.cfg")
	if _, err := os.Stat(nextcloudSettingsPath); err == nil {
		Log.Info("Update Nextcloud settings...")
		cfg, err := ini.Load(nextcloudSettingsPath)
		if err != nil {
			Log.Error("Fail to read file:", err)
		}
		keys := cfg.Section("Accounts").Keys()
		for _, key := range keys {
			if strings.Contains(key.Name(), "localPath") {
				Log.Info("Update path of %s : %s", key.Name(), key.Value())
			}
		}
	} else {
		Log.Warningf("Nextcloud settings not found in %s", nextcloudSettingsPath)
	}*/

	Launch(os.Args[1:])
}
