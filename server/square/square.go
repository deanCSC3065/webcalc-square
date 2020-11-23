package square

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	xQueryParam = "x"
	yQueryParam = "y"
)

// Handle provides logic for handling queries to the `/modulo` endpoint, with error handling and
// success responses
func Handle(w http.ResponseWriter, r *http.Request) {
	xParam := r.URL.Query().Get(xQueryParam)
	// Check params provided
	if xParam == "" {
		server.Error(w, false, "The x field is required.", http.StatusBadRequest)
		return
	}
	yParam := r.URL.Query().Get(yQueryParam)
	if yParam == "" {
		server.Error(w, false, "The y field is required.", http.StatusBadRequest)
		return
	}

	// Convert to integers
	xValue, err := strconv.Atoi(xParam)
	if err != nil {
		server.Error(w, false, fmt.Sprintf("The value '%s' is not valid.", xParam), http.StatusBadRequest)
		return
	}
	yValue, err := strconv.Atoi(yParam)
	if err != nil {
		server.Error(w, false, fmt.Sprintf("The value '%s' is not valid.", yParam), http.StatusBadRequest)
		return
	}

	// Check for division by zero
	if yValue == 0 {
		server.Error(w, false, "Division by zero is not allowed.", http.StatusBadRequest)
		return
	}

	// Perform calculation
	answer := maths.Modulo(xValue, yValue)
	server.Success(w, true, fmt.Sprintf("%d%%%d=%d", xValue, yValue, answer), answer)
}
