package service

import (
	"context"
	"go_project16/proto"
)

//
type Server struct {
}

func (s *Server) GetPerson(ctx context.Context, request *proto.PersonRequest) (*proto.PersonResponse, error) {

	return &proto.PersonResponse{
		Person: &proto.Person{
			Name: "张三",
			Sex:  "男",
		},
	}, nil
}

func (s *Server) ListPerson(ctx context.Context, request *proto.PersonRequest) (*proto.ListPersonResponse, error) {
	var listPerson []*proto.Person
	listPerson = append(listPerson, &proto.Person{
		Name: "张三",
		Sex:  "男",
	})
	listPerson = append(listPerson, &proto.Person{
		Name: "李四",
		Sex:  "男",
	})
	return &proto.ListPersonResponse{
		Person: listPerson,
	}, nil
}
