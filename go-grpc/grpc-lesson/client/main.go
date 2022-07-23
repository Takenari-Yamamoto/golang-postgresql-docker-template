package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main () {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewFilesServiceClient(conn)
	// callListFiles(client)
	callDownload(client)
}

func callListFiles(client pb.FilesServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetFilenames())

}

func callDownload (client pb.FilesServiceClient) {
	req := &pb.DownloadRequest{ Filename: "name.txt" }
	stream, err := client.Dowload(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Response from Downloaded(bytes): %v", res.GetData())
		log.Printf("Response from Downloaded(string): %v", string(res.GetData()))
	}
}
