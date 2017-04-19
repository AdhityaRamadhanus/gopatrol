package grpc

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	checkup "github.com/AdhityaRamadhanus/checkupd"
	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"golang.org/x/net/context"
)

//AddTCPEndpoint is grpc service to add tcp endpoint to checkup server
func (handler *ServiceHandler) AddTCPEndpoint(ctx context.Context, request *checkupservice.AddTcpEndpointRequest) (*checkupservice.EndpointResponse, error) {
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	log.Println(request)
	tcpChecker := checkup.TCPChecker{
		Name:          request.Endpoint.Name,
		URL:           request.Endpoint.Url,
		ThresholdRTT:  time.Duration(request.Endpoint.Thresholdrtt),
		Attempts:      int(request.Endpoint.Attempts),
		TLSSkipVerify: request.TlsSkipVerify,
		TLSEnabled:    request.TlsEnabled,
		TLSCAFile:     request.TlsCaFile,
		Timeout:       time.Duration(request.Timeout),
	}
	handler.CheckupServer.Checkers = append(handler.CheckupServer.Checkers, tcpChecker)
	message := fmt.Sprintf("Tcp Endpoint %s->%s Added", tcpChecker.Name, tcpChecker.URL)
	log.Printf(message)
	return &checkupservice.EndpointResponse{Message: message}, nil
}

//AddHTTPEndpoint is grpc service to add http endpoint to checkup server
func (handler *ServiceHandler) AddHTTPEndpoint(ctx context.Context, request *checkupservice.AddHttpEndpointRequest) (*checkupservice.EndpointResponse, error) {
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	log.Println(request)
	httpChecker := checkup.HTTPChecker{
		Name:           request.Endpoint.Name,
		URL:            request.Endpoint.Url,
		ThresholdRTT:   time.Duration(request.Endpoint.Thresholdrtt),
		Attempts:       int(request.Endpoint.Attempts),
		MustContain:    request.MustContain,
		MustNotContain: request.MustNotContain,
		UpStatus:       int(request.Upstatus),
	}
	// Parsing Http Headers
	reader := bufio.NewReader(strings.NewReader(request.Headers + "\r\n"))
	tp := textproto.NewReader(reader)
	httpHeaders, err := tp.ReadMIMEHeader()
	if err == nil {
		httpChecker.Headers = http.Header(httpHeaders)
	}

	handler.CheckupServer.Checkers = append(handler.CheckupServer.Checkers, httpChecker)
	message := fmt.Sprintf("Http Endpoint %s->%s Added", httpChecker.Name, httpChecker.URL)
	log.Printf(message)
	return &checkupservice.EndpointResponse{Message: message}, nil
}

//AddDNSEndpoint is grpc service to add dns endpoint to checkup server
func (handler *ServiceHandler) AddDNSEndpoint(ctx context.Context, request *checkupservice.AddDNSEndpointRequest) (*checkupservice.EndpointResponse, error) {
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	log.Println(request)
	dnsChecker := checkup.DNSChecker{
		Name:         request.Endpoint.Name,
		URL:          request.Endpoint.Url,
		ThresholdRTT: time.Duration(request.Endpoint.Thresholdrtt),
		Attempts:     int(request.Endpoint.Attempts),
		Timeout:      time.Duration(request.Timeout),
		Host:         request.Hostname,
	}

	handler.CheckupServer.Checkers = append(handler.CheckupServer.Checkers, dnsChecker)
	message := fmt.Sprintf("DNS Endpoint %s->%s %s Added", dnsChecker.Name, dnsChecker.URL, request.Hostname)
	log.Printf(message)
	return &checkupservice.EndpointResponse{Message: message}, nil
}
