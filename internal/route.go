package internal

func (s *Server) Route() {
	s.Router.POST("/upload-file",  s.test())
}
