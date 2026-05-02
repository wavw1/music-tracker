package handler

type HTTPHandlers struct {
}

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHttpHandlers() *HTTPHandlers {
	return &HTTPHandlers{}
}

func NewHTTPServer(
	httpHandlers *HTTPHandlers,
) HTTPServer {
	return HTTPServer{
		httpHandlers: httpHandlers,
	}
}
