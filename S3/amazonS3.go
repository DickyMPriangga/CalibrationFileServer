package S3

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sessionAWS *session.Session
var bucketName string
var awsRegion string

func init() {
	//err := godotenv.Load("dev.env")
	//if err != nil {
	//	log.Fatal("Error loading awsS3.env file", err)
	//}

	sessionAWS = connectAWS()
}

func connectAWS() *session.Session {
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion = os.Getenv("AWS_REGION")
	bucketName = os.Getenv("BUCKET_NAME")

	sess, err := session.NewSession(
		&aws.Config{
			Region: &awsRegion,
			Credentials: credentials.NewStaticCredentials(
				awsAccessKeyID,
				awsSecretKey,
				"",
			),
		},
	)

	if err != nil {
		log.Fatal("Error when creating session to AWS", err)
	}

	return sess
}

func UploadFile(file multipart.File, header *multipart.FileHeader) (bool, string) {
	uploader := s3manager.NewUploader(sessionAWS)

	filename := header.Filename

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		log.Fatal("Error when uploading file : ", err)
	}

	filepath := "https://" + bucketName + "." + "s3-" + awsRegion + ".amazonaws.com/" + filename

	return true, filepath
}
