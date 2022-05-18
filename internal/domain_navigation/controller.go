package domain_navigation

import (
	"dns/internal/models"
	"dns/internal/validator"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Handler Interface for the DataBank
type Handler interface {
	DataBankLocator(w http.ResponseWriter, r *http.Request)
}

// Controller for the DataBank
type Controller struct {
	DataBankCalculator
}

// DataBankCalculator calculate the location for the drone
// request contains the coordinates for the position of the drone
// Response contains the location calculated by the API.
func (dbc *Controller) DataBankLocator(w http.ResponseWriter, r *http.Request) {

	// Decode the data to the req struct.
	var req DataLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// error handling
		e := models.NewError(err.Error(), models.NOK, http.StatusInternalServerError)
		log.Errorf("error while unmarshalling the request body: %s", e.Reason)

		rs := models.NewResponseSetter(e, e.Code)
		rs.SendHttpResponse(w)
		return
	}

	// validates the request
	v := validator.NewValidator()
	if vErr := v.Validate(req); vErr != nil {
		e := models.NewError(vErr.ErrorMessage, models.NOK, http.StatusBadRequest)
		rs := models.NewResponseSetter(e, e.Code)
		rs.SendHttpResponse(w)
		return
	}

	// Calculate the location
	loc := dbc.DataBankCalculator.Calculate(req)

	res := Response{
		Location: loc,
		Status: models.OK,
	}

	// return the response.
	rs := models.NewResponseSetter(res, http.StatusOK)
	rs.SendHttpResponse(w)

}
