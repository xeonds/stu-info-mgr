package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "stu-info-mgr/proto"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (s *server) Query(ctx context.Context, req *pb.QueryRequest) (*pb.Student, error) {
	student := &pb.Student{
		Id:   req.Id,
		Name: "John Doe",
	}
	return student, nil
}
func (s *server) QueryByName(ctx context.Context, req *pb.QueryByNameRequest) (*pb.Student, error) {
	student := &pb.Student{
		Name: req.Name,
	}
	return student, nil
}
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Success: true}, nil
}
func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{Success: true}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
