package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func dumpStart(d dumpTarget) result {
	log.Println("Starting dump")
	var region string
	if d.Region != "" {
		region = d.Region
	} else {
		region = "us-east-1"
	}

	t := time.Now()
	mongoHost := os.Getenv("MONGOHOST")
	mongoPort := os.Getenv("MONGOPORT")
	if mongoHost == "" {
		mongoHost = "mongo"
	}
	if mongoPort == "" {
		mongoPort = "27017"
	}
	var r result

	path, err := exec.LookPath("mongodump")
	if err != nil {
		log.Println("Could not find Mongodump", err)
		r.Result = "Could not find Mongodump " + err.Error()
		return r
	}
	fmt.Printf("mongodump is available at %s ", path)

	dumpCmd := exec.Command("mongodump", "--host", mongoHost, "--port", mongoPort, "--archive")
	body, err := dumpCmd.StdoutPipe()
	if err != nil {
		log.Println("failed executing mongodump", err)
		r.Result = "failed executing mongodump " + err.Error()
		return r
	}

	if err := dumpCmd.Start(); err != nil {
		log.Println("failed starting mongodump", err)
		r.Result = "failed starting mongodump " + err.Error()
		return r
	}

	// Wrap the pipe to hide the seek methods from the uploader
	var bodyWrap = struct {
		io.Reader
	}{body}

	if d.Bucket != "" {
		uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String(region)}))
		result, err := uploader.Upload(&s3manager.UploadInput{
			Body:   bodyWrap,
			Bucket: aws.String(d.Bucket),
			Key:    aws.String(d.Path + t.Format("2006-01-02T15:04")),
		})

		if err != nil {
			log.Println("Failed to upload", err)
			r.Result = "Failed to upload " + err.Error()
			return r
		}

		if err := dumpCmd.Wait(); err != nil {
			log.Println("Failed to dump", err)
			r.Result = "Failed to dump " + err.Error()
			return r
		}

		log.Println("Successfully uploaded to", result.Location)
		r.Result = "success"
		return r
	}
	if err := dumpCmd.Wait(); err != nil {
		log.Println("Failed to dump", err)
		r.Result = "Failed to dump " + err.Error()
		return r
	}

	log.Println("Successfully dumped to null")
	r.Result = "success"
	return r
}
