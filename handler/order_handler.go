package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderHandler interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindOne(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
