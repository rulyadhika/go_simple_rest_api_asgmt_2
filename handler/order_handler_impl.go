package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderHandlerImpl struct{}

func NewOrderHandlerImpl() *OrderHandlerImpl {
	return &OrderHandlerImpl{}
}

func (o *OrderHandlerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderHandlerImpl) FindOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderHandlerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderHandlerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
