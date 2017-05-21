package cli

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"encoding/json"

	"bytes"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/urfave/cli"
)

func addHTTPEndpoint(c *cli.Context) error {
	endpointName := ""
	endpointURL := ""
	if c.NArg() > 1 {
		endpointName = c.Args().Get(0)
		endpointURL = c.Args().Get(1)
	} else {
		fmt.Println("Please provice name and url to add http checker")
		cli.ShowCommandHelp(c, "add-http")
		return cli.NewExitError("", 1)
	}
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/gopatrol.sock")
			},
		},
	}

	checkerReq := gopatrol.HTTPChecker{
		Type:           "http",
		Name:           endpointName,
		URL:            endpointURL,
		Attempts:       c.Int("attempts"),
		ThresholdRTT:   time.Duration(c.Int64("thresholdrtt")),
		MustContain:    c.String("mustcontain"),
		MustNotContain: c.String("mustnotcontain"),
		UpStatus:       c.Int("upstatus"),
	}

	jsonReq, _ := json.Marshal(checkerReq)

	var response *http.Response
	var err error
	response, err = client.Post("http://unix/api/v1/checkers/create", "application/json", bytes.NewReader(jsonReq))
	if err != nil {
		log.Println(err)
		return err
	}

	decodedResp := map[string]interface{}{}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&decodedResp); err != nil {
		log.Println(err)
	}
	switch response.StatusCode {
	case 200:
		log.Println(response.StatusCode, decodedResp["message"])
	default:
		log.Println(decodedResp["error"])
	}

	return nil
}

func addTCPEndpoint(c *cli.Context) error {
	endpointName := ""
	endpointURL := ""
	if c.NArg() > 1 {
		endpointName = c.Args().Get(0)
		endpointURL = c.Args().Get(1)
	} else {
		fmt.Println("Please provice name and url to add tcp endpoint")
		cli.ShowCommandHelp(c, "add-tcp")
		return cli.NewExitError("", 1)
	}
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/gopatrol.sock")
			},
		},
	}

	checkerReq := gopatrol.TCPChecker{
		Type:          "tcp",
		Name:          endpointName,
		URL:           endpointURL,
		Attempts:      c.Int("attempts"),
		ThresholdRTT:  time.Duration(c.Int64("thresholdrtt")),
		TLSEnabled:    c.Bool("tls-enabled"),
		TLSCAFile:     c.String("tls-ca"),
		TLSSkipVerify: c.Bool("tls-skip-verify"),
		Timeout:       time.Duration(c.Int64("timeout")),
	}

	jsonReq, _ := json.Marshal(checkerReq)

	var response *http.Response
	var err error
	response, err = client.Post("http://unix/api/v1/checkers/create", "application/json", bytes.NewReader(jsonReq))

	if err != nil {
		log.Println(err)
		return err
	}
	decodedResp := map[string]interface{}{}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&decodedResp); err != nil {
		log.Println(err)
	}
	log.Println(decodedResp["message"])

	return nil
}

func addDNSEndpoint(c *cli.Context) error {
	endpointName := ""
	endpointURL := ""
	endpointHost := ""

	if c.NArg() > 2 {
		endpointName = c.Args().Get(0)
		endpointURL = c.Args().Get(1)
		endpointHost = c.Args().Get(2)
	} else {
		fmt.Println("Please provice name, url (nameserver) and the hostname to add dns endpoint")
		cli.ShowCommandHelp(c, "add-dns")
		return cli.NewExitError("", 1)
	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/gopatrol.sock")
			},
		},
	}

	checkerReq := gopatrol.DNSChecker{
		Type:         "dns",
		Name:         endpointName,
		URL:          endpointURL,
		Attempts:     c.Int("attempts"),
		ThresholdRTT: time.Duration(c.Int64("thresholdrtt")),
		Timeout:      time.Duration(c.Int64("timeout")),
		Host:         endpointHost,
	}

	jsonReq, _ := json.Marshal(checkerReq)

	var response *http.Response
	var err error
	response, err = client.Post("http://unix/api/v1/checkers/create", "application/json", bytes.NewReader(jsonReq))

	if err != nil {
		log.Println(err)
		return err
	}
	decodedResp := map[string]interface{}{}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&decodedResp); err != nil {
		log.Println(err)
	}
	log.Println(decodedResp["message"])

	return nil
}
