package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/hieunm6/say-grpc/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file where the output will written")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("usage:\n\t%s \"text to speak\"", os.Args[0])
		os.Exit(1)
	}

	connect, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}

	defer connect.Close()
	client := pb.NewTexToSpeechClient(connect)
	text := &pb.Text{Text: os.Args[1]}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not write file to %s: %v", *output, err)
	}
}
