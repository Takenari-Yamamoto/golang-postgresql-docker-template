package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFilesServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {

	fmt.Println("ListFiles was invoked")
	dir := "/Users/takenariyamamoto/dev/go-practice/go-grpc/grpc-lesson/storage"

	paths, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil,  err
	}

	var filenames []string

	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := &pb.ListFilesResponse {
		Filenames: filenames,
	}

	return res, nil

}

func (*server) Dowload(req *pb.DownloadRequest, stream pb.FilesService_DowloadServer) error {
	fmt.Println("Download was invoked")

	filename := req.GetFilename()
	path := "/Users/takenariyamamoto/dev/go-practice/go-grpc/grpc-lesson/storage/" + filename

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		res := &pb.DownloadResponse{Data: buf[:n]}
		sendErr := stream.Send(res)
		if sendErr != nil {
			return sendErr
		}
		time.Sleep(1 * time.Second)

	}
	return nil

}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFilesServiceServer(s, &server{})

	fmt.Println("server is running")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}


}
