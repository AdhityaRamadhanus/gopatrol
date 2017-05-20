package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"bytes"

	"github.com/urfave/cli"
)

func deleteEndpoint(c *cli.Context) error {
	endpointSlug := ""

	if c.NArg() > 0 {
		endpointSlug = c.Args().Get(0)
	} else {
		fmt.Println("Please provide endpoint slug")
		cli.ShowCommandHelp(c, "delete")
		return cli.NewExitError("", 1)
	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/gopatrol.sock")
			},
		},
	}

	// jsonReq, _ := json.Marshal(checkerReq)

	var response *http.Response
	var err error
	url := "http://unix/api/v1/checkers/" + endpointSlug + "/delete"
	log.Println(url)
	response, err = client.Post(url, "application/json", bytes.NewReader(nil))

	if err != nil {
		log.Println(err)
		return err
	}
	decodedResp := map[string]interface{}{}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&decodedResp); err != nil {
		log.Println(err)
		return err
	}
	log.Println(decodedResp)
	return nil
}
