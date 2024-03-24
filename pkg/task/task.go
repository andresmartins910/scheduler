package task

import (
	"scheduler/pkg/model"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/hibiken/asynq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TypeReport = "report"
)

// client
func NewReportTask(ctx context.Context, report model.Report) (*asynq.Task, error) {
	payload, err := json.Marshal(report)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeReport, payload), nil
}

// worker
func ProcessReport(ctx context.Context, t *asynq.Task) error {
	key := "DO00AV9FDC63L7MZLGML"
	secret := "9hBvzmSJEUsAJv+PV+wpBvdaAWuDMg1CGpRvAQGcJPw"

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String("https://nyc3.digitaloceanspaces.com"),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String("nyc3"),
		HTTPClient:       &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.New(newSession)

	dsn := "root:220422@ndrE@tcp(127.0.0.1:3306)/scheduler?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var reports []model.Report
	result := db.Where("DATE(created_at) = CURDATE()").Find(&reports)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	for _, report := range reports {
		reportJSON, err := json.Marshal(report)
		if err != nil {
			log.Fatal(err)
		}

		object := s3.PutObjectInput{
			Bucket: aws.String("reportsss"),
			Key:    aws.String(fmt.Sprintf("reports/%s.json", report.Name)),
			Body:   strings.NewReader(string(reportJSON)),
			ACL:    aws.String("private"),
			Metadata: map[string]*string{
				"x-amz-meta-my-key": aws.String("your-value"),
			},
		}

		_, err = s3Client.PutObject(&object)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
