package domain_navigation

import (
	"math"
)

// DataBank Calculator
type DataBankCalculator interface {
	Calculate(request DataLocationRequest) Location
}

//DatBank Service
type DataBankService struct {
	SectorId int64
}

// Calculate calculates the Location using the coordinates and Velocity
func (dbs *DataBankService) Calculate(r DataLocationRequest) Location {
	// get the sectorId from the configuration.
	sectorId := float64(dbs.SectorId)

	// calculate the location
	// location := sectorId * X +  sectorId * Y + sectorId * Z + Velocity
	loc := Location((sectorId*float64(r.X) + sectorId*float64(r.Y) + sectorId*float64(r.Z)) + float64(r.Vel))

	// format the location to 2 decimal
	return loc.Format()
}

// Format Format the value to 2 decimals
func (loc Location) Format() Location {
	//rounding off to the 2 digits.
	return Location(math.Round(float64(loc)*100) / 100)
}
