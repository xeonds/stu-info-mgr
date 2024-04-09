package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"stu-info-mgr/config"
	"stu-info-mgr/lib"
	pb "stu-info-mgr/proto"
)

type server struct {
	db *gorm.DB
	pb.UnimplementedStudentServiceServer
}

func main() {
	flag.Parse()
	config := lib.LoadConfig[config.Config]()
	db := lib.NewDB(&config.DatabaseConfig, func(db *gorm.DB) error {
		return db.AutoMigrate(&pb.Student{})
	})
	port := flag.Int("port", config.Server.Port, "The server port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, &server{
		db: db,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
