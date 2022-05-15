package cmd

import (
	"dns/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

const (
	portEVar    = "PORT"
	portDefault = "3000"

	versionEvar    = "VERSION"
	versionDefault = "v1"

	sectorIDEvar   = "SECTOR_ID"
	sectorIDDefult = "1"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "The command is used for the starting the server",

	Run: serveRun,
}

func init() {
	RootCmd.AddCommand(serveCmd)
	bindEvar(portEVar, portDefault)
	bindEvar(versionEvar, versionDefault)
	bindEvar(sectorIDEvar, sectorIDDefult)

}

func serveRun(*cobra.Command, []string) {
	port, err := strconv.Atoi(viper.GetString(portEVar))
	if err != nil {
		log.Fatalln("invalid port")
	}

	c := &app.Config{
		Port:     port,
		Name:     "dns",
		Version:  viper.GetString(versionEvar),
		SectorId: viper.GetInt64(sectorIDEvar),
	}

	s := app.NewServer(c)
	s.Start()
}

func bindEvar(evar, defaultValue string) {
	viper.SetDefault(evar, defaultValue)
	viper.BindEnv(evar)
}
