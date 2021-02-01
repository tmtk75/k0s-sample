package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	RunE: RunE,
}

func main() {
	RootCmd.Execute()
}

func RunE(cmd *cobra.Command, args []string) error {
	sess, err := session.NewSession(&aws.Config{
		//Region:      aws.String("us-west-2"),
		//Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	conf := &aws.Config{
		//Endpoint:         aws.String("http://10.96.10.1:4566"),
		//S3ForcePathStyle: aws.Bool(true),
	}
	svc := s3.New(sess, conf)
	r, err := svc.GetObject(&s3.GetObjectInput{Bucket: aws.String("my1st-bucket"), Key: aws.String("foo")})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%v", r)
	return nil
}

func init() {
	pf := RootCmd.PersistentFlags()
	pf.String("endpoint", "")
}
