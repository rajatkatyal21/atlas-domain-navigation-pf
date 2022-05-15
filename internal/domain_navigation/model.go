package domain_navigation

import "dns/internal/models"

type Location float64
type coordinate float64

type DataLocationRequest struct {
	X   coordinate `validate:"required"`
	Y   coordinate `validate:"required"`
	Z   coordinate `validate:"required"`
	Vel coordinate `validate:"required"`
}

type Response struct {
	Location Location `json:"loc"`
	Status   models.Status
}
