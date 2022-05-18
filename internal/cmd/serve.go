package cmd

import (
	"dns/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

// Environment variables and their default values
const (
	portEVar    = "PORT"
	portDefault = "3000"

	versionEvar    = "VERSION"
	versionDefault = "v1"

	sectorIDEvar   = "SECTOR_ID"
	sectorIDDefult = "1"
)

// Serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: `The command is used for the starting the server
			to start the server
			dns server`,

	Run: serveRun,
}

func init() {
	RootCmd.AddCommand(serveCmd)
	// binding environment variable PORT to default value
	bindEvar(portEVar, portDefault)
	// binding environment variable VERSION to default value
	bindEvar(versionEvar, versionDefault)
	// binding environment variable SECTOR_ID to default value
	bindEvar(sectorIDEvar, sectorIDDefult)

}

//
func serveRun(*cobra.Command, []string) {
	port, err := strconv.Atoi(viper.GetString(portEVar))
	if err != nil {
		log.Fatalln("invalid port")
	}

	// initialize the config
	c := &app.Config{
		Port:     port,
		Name:     "dns",
		Version:  viper.GetString(versionEvar),
		SectorId: viper.GetInt64(sectorIDEvar),
	}

	// create Server using config
	s := app.NewServer(c)

	// start the server
	s.Start()
}

// Bind Environment variable to default value
func bindEvar(evar, defaultValue string) {
	viper.SetDefault(evar, defaultValue)
	viper.BindEnv(evar)
}
