package internal

// TODO: server authentication
func (s *Server) Route() {
	s.Router.POST("/trade", s.checkAuthentication(s.trade()))
}
