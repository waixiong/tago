package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	"rj/tago/auth"
	authPB "rj/tago/auth/proto"
)

var (
	tls      = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("../key/certs/mycert.pem", "../key/certs/mycert.pem", "The TLS cert file")
	keyFile  = flag.String("../key/private/mykey.pem", "../key/private/mykey.pem", "The TLS key file")
	//jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port = flag.Int("port", 8080, "The server port")
	//svc  = &dynamodb.DynamoDB{}
)

func main() {
	print("Start\n")
	flag.Parse()
	print("1\n")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	print("2\n")
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			println("CERT not found use Alt")
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		print(*keyFile)
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	print("3\n")
	grpcServer := grpc.NewServer(opts...)
	print("4\n")
	authPB.RegisterAuthenticationServiceServer(grpcServer, auth.NewServer())
	print("5\n")

	grpcServer.Serve(lis)
	print("Running...\n")
}
