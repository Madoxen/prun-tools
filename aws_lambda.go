package main

import (
	"context"
	"log"
	"os"
	"bytes"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/prun-tools/tools/recipe_calc"
	"github.com/prun-tools/tools/trade_finder"
)

type MyEvent struct {
	Status string `json:"status"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	// Run tools
	recipe_calc_out, err := recipe_calc.Calculate()
	if err != nil {
		log.Printf("ERROR - could not run recipe calc: %s", err)
		return "FAILURE", err
	}	
	trade_finder_out, err := trade_finder.Calculate()
	if err != nil {
		log.Printf("ERROR - could not run trade finder, %s", err)
		return "FAILURE", err
	}	

	keys_objects := map[string]*bytes.Buffer {
		"recipe_calc.csv" : recipe_calc_out,
		"trade_finder.csv": trade_finder_out,
	}

	aws_region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(aws_region),
		// Add more configurations if needed, such as credentials, endpoint, etc.
	})

	if err != nil {
		log.Printf("ERROR - could not create session: %s", err)
		return "FAILURE", err
	}

	svc := s3.New(sess)

	for key, object := range keys_objects {
		bucket := os.Getenv("OUTPUT_BUCKET")

		_, err = svc.PutObject(&s3.PutObjectInput{
			Body:   bytes.NewReader(object.Bytes()),
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			log.Printf("ERROR - could not PUT an object into S3: %s", err)
			return "FAILURE", err
		} else {
			log.Printf("File uploaded successfully.")
		}
	}
	return "SUCCESS", nil
}

func main() {
	lambda.Start(HandleRequest)
}
