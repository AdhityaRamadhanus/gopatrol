package checklist

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func createGrpcClient(cliContext *cli.Context) (*grpc.ClientConn, error) {
	if cliContext.Bool("tls") {
		godotenv.Load()
		creds, _ := credentials.NewClientTLSFromFile(os.Getenv("CHECKLIST_CA"), os.Getenv("CHECKLIST_HOST"))
		return grpc.Dial(cliContext.String("host"),
			grpc.WithTransportCredentials(creds),
		)
	} else {
		return grpc.Dial(cliContext.String("host"), grpc.WithInsecure())
	}
}
