package main

import (
	"fmt"

	pb "github.com/cloudtv/backend/pb/auth"
)

// Import generated proto package

type server struct {
	pb.UnimplementedAuthServiceServer
}

func main() {
	fmt.Println("cloudTV :)")

}
