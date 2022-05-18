package health_check

import (
	"dns/internal/models"
	"net/http"
)

// Health check hanler
type Handler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

// Controller to implement the Handler function.
type Controller struct {}

// HealthCheck API is used by the LB to check if the container is UP or down.
func (s *Controller) HealthCheck(w http.ResponseWriter, r *http.Request) {

	res := &Response{
		Status: "OK",
	}

	rs := models.NewResponseSetter(res, 200)
	rs.SendHttpResponse(w)
}

