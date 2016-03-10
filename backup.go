package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	path, err := exec.LookPath("mongodump")
	if err != nil {
		log.Fatal("mongodump could not be found")
	}
	fmt.Printf("mongodump is available at %s\n", path)

	dumpCmd := exec.Command(
		"mongodump",
		"--host",
		"10.10.1.103",
		"--port",
		"27017",
		"--archive"
	)
	body, err := dumpCmd.StdoutPipe()
	if err != nil {
		// handle error
	}

	if err := dumpCmd.Start(); err != nil {
		// handle error
	}

	// Wrap the pipe to hide the seek methods from the uploader
	var bodyWrap = struct {
		io.Reader
	}{body}

	uploader := s3manager.NewUploader(
		session.New(&aws.Config{Region: aws.String("us-east-1")})
	)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   bodyWrap,
		Bucket: aws.String("net-openwhere-mongodb-snapshots-dev"),
		Key:    aws.String("myKey2"),
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	if err := dumpCmd.Wait(); err != nil {
		// handle error
	}

	log.Println("Successfully uploaded to", result.Location)
}
