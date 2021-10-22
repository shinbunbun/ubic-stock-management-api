package s3

import (
	"bytes"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewS3Session() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApNortheast1RegionID),
	})
}

// Up先URLとエラーを返却
func UploadToS3(sess *session.Session, wb *bytes.Buffer, id string) (string, error) {
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(config.GetS3BucketName()),
		Key:         aws.String(id + ".jpeg"),
		Body:        bytes.NewReader(wb.Bytes()),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}
	return res.Location, nil
}
