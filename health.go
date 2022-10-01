package tel

import (
	"encoding/json"
	"net/http"

	health "github.com/tel-io/tel/v2/monitoring/heallth"
)

type HealthHandler struct {
	health.CompositeChecker
}

// NewHealthHandler returns a new Handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// ServeHTTP returns a json encoded health
// set the status to http.StatusServiceUnavailable if the check is down
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	checker := h.CompositeChecker.Check()

	body, err := json.Marshal(checker)
	if err != nil {
		FromCtx(r.Context()).Error("health check encode failed", Error(err))
	}

	if checker.Is(health.Down) {
		w.WriteHeader(http.StatusServiceUnavailable)
		FromCtx(r.Context()).Error("health", String("body", string(body)))
	}

	if _, err = w.Write(body); err != nil {
		FromCtx(r.Context()).Error("health check encode failed", Error(err))
	}
}
