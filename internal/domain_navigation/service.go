package domain_navigation

import (
	"math"
)

type DataBankCalculator interface {
	Calculate(request DataLocationRequest) Location
}

type DataBankService struct {
	SectorId int64
}

// Calculate calculates the Location using the coordinates and Velocity
func (dbs *DataBankService) Calculate(r DataLocationRequest) Location {
	sectorId := float64(dbs.SectorId)
	loc := Location((sectorId*float64(r.X) + sectorId*float64(r.Y) + sectorId*float64(r.Z)) + float64(r.Vel))
	return loc.Format()
}

// Format Format the value to 2 decimals
func (loc Location) Format() Location {
	return Location(math.Round(float64(loc)*100) / 100)
}
