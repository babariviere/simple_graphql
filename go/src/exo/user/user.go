package user

import "exo/user/pb"

// User contains project metadata assigned to an ID
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// FromGRPC converts pb's user to user
func FromGRPC(u *pb.User) User {
	return User{ID: u.Id, Email: u.Email}
}

// ToGRPC converts user to pb's user
func (u User) ToGRPC() *pb.User {
	return &pb.User{Id: u.ID, Email: u.Email}
}
