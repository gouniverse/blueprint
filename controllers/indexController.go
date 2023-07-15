package controllers

import "net/http"

type indexController struct{}

func NewIndexController() *indexController {
	return &indexController{}
}

func (controller *indexController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	return "Hello world"
}
