package domain_navigation

import (
	"dns/internal/models"
	"dns/internal/validator"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handler interface {
	DataBankLocator(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	DataBankCalculator
}

func (dbc *Controller) DataBankLocator(w http.ResponseWriter, r *http.Request) {

	var req DataLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e := models.NewError(err.Error(), models.NOK, http.StatusInternalServerError)
		log.Errorf("error while unmarshalling the request body: %s", e.Reason)

		rs := models.NewResponseSetter(e, e.Code)
		rs.SendHttpResponse(w)
		return
	}

	v := validator.NewValidator()
	if vErr := v.Validate(req); vErr != nil {
		e := models.NewError(vErr.ErrorMessage, models.NOK, http.StatusBadRequest)
		rs := models.NewResponseSetter(e, e.Code)
		rs.SendHttpResponse(w)
		return
	}

	loc := dbc.DataBankCalculator.Calculate(req)

	res := Response{
		Location: loc,
		Status: models.OK,
	}

	rs := models.NewResponseSetter(res, http.StatusOK)
	rs.SendHttpResponse(w)

}
