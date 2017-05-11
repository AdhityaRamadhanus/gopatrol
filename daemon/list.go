package daemon

//ListEndpoint is grpc service that returns all the endpoints in checkup server
func (handler *Daemon) ListEndpoint(q map[string]interface{}) ([]interface{}, error) {
	return handler.EndpointService.GetAllEndpoints(q)
}

//ListEndpoint is grpc service that returns all the endpoints in checkup server
func (handler *Daemon) GetEndpointBySlug(slug string) (interface{}, error) {
	return handler.EndpointService.GetEndpointBySlug(slug)
}
