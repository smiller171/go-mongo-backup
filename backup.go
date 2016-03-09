package main

import (
	"fmt"
	"log"
	"os"
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

	dumpCmd := exec.Command("mongodump", "--host", "10.10.1.103", "--port", "27017", "--archive=file.txt")

	dumpOut, err := dumpCmd.Output()
	fmt.Println(string(dumpOut))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dumpOut))

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   file,
		Bucket: aws.String("net-openwhere-mongodb-snapshots-dev"),
		Key:    aws.String("myKey"),
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	log.Println("Successfully uploaded to", result.Location)
}
