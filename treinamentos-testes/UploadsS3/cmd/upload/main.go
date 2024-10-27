package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

// init initializes the S3 client
func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIA23456789012345678",
				"1234567890123456789012345678901234567890",
				"",
			),
		},
	)
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "test-bucket-123"
}

func main() {
	// Open the directory containing the files to upload
	dir, err := os.Open("../../tmp")
	if err != nil {
		log.Fatalf("failed to open directory, %v", err)
	}

	defer dir.Close()

}

func uploadFiles(dir *os.File) error {
	return nil
}
