package checklist

import (
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func unixDialer(target string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("unix", target, timeout)
}

func createGrpcClient(cliContext *cli.Context) (*grpc.ClientConn, error) {
	if cliContext.Bool("tls") {
		godotenv.Load()
		creds, _ := credentials.NewClientTLSFromFile(os.Getenv("CHECKLIST_CA"), os.Getenv("CHECKLIST_HOST"))
		return grpc.Dial(cliContext.String("host"),
			grpc.WithTransportCredentials(creds),
		)
	}
	return grpc.Dial("/tmp/checkupd.sock",
		grpc.WithDialer(unixDialer),
		grpc.WithInsecure())
}
