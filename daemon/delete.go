package daemon

import (
	"fmt"
	"log"

	checkup "github.com/AdhityaRamadhanus/gopatrol"
	"github.com/pkg/errors"
)

//DeleteEndpoint is grpc service to delete some endpoint
func (handler *Daemon) DeleteEndpoint(slug string) error {
	if err := handler.EndpointService.DeleteEndpointBySlug(slug); err != nil {
		log.Println("error deleting endpoint", err)
		return err
	}
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	var deletedEndpoint checkup.Checker
	for i, endpoint := range handler.Checkers {
		if endpoint.GetSlug() == slug {
			deletedEndpoint = handler.Checkers[i]
			handler.Checkers = append(handler.Checkers[:i], handler.Checkers[i+1:]...)
			break
		}
	}
	message := "Cannot find endpoint to delete"
	if deletedEndpoint != nil {
		message = fmt.Sprintf("Endpoint %s->%s Deleted", deletedEndpoint.GetName(), deletedEndpoint.GetURL())
		log.Printf(message)
	} else {
		return errors.New(message)
	}
	return nil
}
