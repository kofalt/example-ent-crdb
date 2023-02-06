package server

import ()

func (s *Server) Routes() {
	s.GET("/users", s.GetUsers)

	s.GET("/user/:id", s.GetUser)
	s.POST("/user/move", s.SetUserAddress)

	s.GET("/top", s.TopRides)

	s.GET("/version", s.Version)
}
