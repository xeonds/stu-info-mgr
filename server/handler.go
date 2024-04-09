package main

import (
	"context"
	pb "stu-info-mgr/proto"
)

func (s *server) Query(ctx context.Context, req *pb.QueryRequest) (*pb.Student, error) {
	student := new(pb.Student)
	if err := s.db.First(student, req.Id).Error; err != nil {
		return nil, err
	}
	return student, nil
}
func (s *server) QueryByName(ctx context.Context, req *pb.QueryByNameRequest) (*pb.Student, error) {
	student := new(pb.Student)
	if err := s.db.Where("name = ?", req.Name).First(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	if s.db.Create(req.Student).Error != nil {
		return &pb.AddResponse{Success: false}, nil
	}
	return &pb.AddResponse{Success: true}, nil
}
func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if s.db.Delete(&pb.Student{Id: req.Id}).Error != nil {
		return &pb.DeleteResponse{Success: false}, nil
	}
	return &pb.DeleteResponse{Success: true}, nil
}
