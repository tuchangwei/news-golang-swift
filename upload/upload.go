package upload

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"mime/multipart"
	"time"
)
var awsSession = func() *session.Session {
	s, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		log.Fatal("aws error:", err)
	}
	return s
}()

func Upload(file multipart.File) (string, error) {
	defer file.Close()
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("images/%d.png", timestamp)
	uploader := s3manager.NewUploader(awsSession)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("go-swift-news"),
		Key: aws.String(fileName),
		Body: file,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		fmt.Println("error: ", err)
		return "", err
	} else {
		fmt.Printf("output:+%v\n", output)

	}
	return output.Location, err

}