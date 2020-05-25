package internal

func (s *Server) Route() {
	s.Router.POST("/trade", s.trade())
}
