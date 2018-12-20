package project

import (
	pb "exo/project/pb"
)

// Project contains project metadata assigned to an ID
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FromGRPC get project from pb's Project
func FromGRPC(p *pb.Project) Project {
	return Project{ID: p.Id, Name: p.Name}
}

// ToGRPC convert project to pb's Project
func (p Project) ToGRPC() *pb.Project {
	return &pb.Project{Id: p.ID, Name: p.Name}
}
