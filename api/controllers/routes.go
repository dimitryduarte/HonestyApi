package controllers

import "github.com/dimitryduarte/honestyapi/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//GET
	s.Router.HandleFunc("/product", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/product/{id}", middlewares.SetMiddlewareJSON(s.GetProductId)).Methods("GET")

	//POST
	//	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/product", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")

	//PUT
	s.Router.HandleFunc("/product", middlewares.SetMiddlewareJSON(s.UpdateProduct)).Methods("PUT")

	//DELETE
	s.Router.HandleFunc("/product/{id}", middlewares.SetMiddlewareJSON(s.DeleteProduct)).Methods("DELETE")

}
