package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io/ioutil"
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
