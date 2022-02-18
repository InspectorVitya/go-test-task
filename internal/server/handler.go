package httpserver

import (
	"net/http"
)

func (s *Server) reversProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method !=  http.MethodGet{
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	data, err := s.App.Get(r.URL.String())
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(data.([]byte))
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
