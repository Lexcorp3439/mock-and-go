package users

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"heisenbug/identification/internal/pb/users"
	"os"
)

func NewClient() (users.UsersClient, *grpc.ClientConn, error) {
	usersAPI := os.Getenv("USERS_API")
	fmt.Println(usersAPI)
	usersConn, err := grpc.Dial(usersAPI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return users.NewUsersClient(usersConn), usersConn, nil
}
