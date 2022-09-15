package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "wairimuian.com/usermanagement_grpc/usermanagement"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := pb.NewUserManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var newUsers = make(map[string]int32)
	newUsers["Alice"] = 43
	newUsers["Bob"] = 30
	for name, age := range newUsers {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("couldn not create user: %v", err)
		}
		log.Printf(`User Details: 
Name: %s
Age: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
