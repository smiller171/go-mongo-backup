package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func dumpStart(d dumpTarget) result {
	t := time.Now()

	var r result
	path, err := exec.LookPath("mongodump")
	if err != nil {
		log.Fatal("mongodump could not be found")
		r.Result = "failed"
		return r
	}
	fmt.Printf("mongodump is available at %s\n", path)

	dumpCmd := exec.Command("mongodump", "--host", "10.10.1.103", "--port", "27017", "--archive")
	body, err := dumpCmd.StdoutPipe()
	if err != nil {
		r.Result = "failed"
		return r
	}

	if err := dumpCmd.Start(); err != nil {
		r.Result = "failed"
		return r
	}

	// Wrap the pipe to hide the seek methods from the uploader
	var bodyWrap = struct {
		io.Reader
	}{body}

	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   bodyWrap,
		Bucket: aws.String(d.Bucket),
		Key:    aws.String(d.Path + t.Format("2006-01-02T15:04")),
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
		r.Result = "failed"
		return r
	}

	if err := dumpCmd.Wait(); err != nil {
		log.Fatalln("Failed to dump", err)
		r.Result = "failed"
		return r
	}

	log.Println("Successfully uploaded to", result.Location)
	r.Result = "success"
	return r
}
