package cli

import (
	"context"
	"log"
	"net"
	"net/http"

	"encoding/json"

	"github.com/urfave/cli"
)

type checker struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Slug string `json:"slug"`
}

func listEndpoint(c *cli.Context) error {
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/gopatrol.sock")
			},
		},
	}

	var response *http.Response
	var err error
	response, err = client.Get("http://unix/api/v1/checkers/all?pagination=false")

	if err != nil {
		log.Println(err)
		return err
	}

	decodedResp := struct {
		Checkers []checker `json:"checkers"`
	}{}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&decodedResp); err != nil {
		log.Println(err)
	}
	log.Println(decodedResp.Checkers)
	return nil
}
