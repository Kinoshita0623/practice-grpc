package main

import (
	"context"
	"fmt"
	"os"

	pb "com.github.Kinoshita0623/practice-grpc/pb/cat"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())

	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewCatClient(conn)
	message := &pb.GetMyCatMessage{TargetCat: "tama"}
	res, err := client.GetMyCat(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
