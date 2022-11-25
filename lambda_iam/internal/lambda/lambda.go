package lambda

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"os"
)

type Service interface {
	Connect() error
	Access() error
}

type LambdaEvent struct {
	sess        *session.Session
	svc         *iam.IAM
	KeyLastUsed *iam.GetAccessKeyLastUsedOutput
}

func (l *LambdaEvent) Connect() error {
	fmt.Println("Connected to Lambda !!")

	var err error
	l.sess, err = session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}
	l.svc = iam.New(l.sess)
	l.KeyLastUsed, err = l.svc.GetAccessKeyLastUsed(&iam.GetAccessKeyLastUsedInput{
		AccessKeyId: aws.String(os.Getenv("AccessKeyID")),
	})
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Success:", *l.KeyLastUsed.AccessKeyLastUsed)
	return nil
}

func (l *LambdaEvent) Access() error {
	return nil
}

func NewLambdaService() *LambdaEvent {
	return &LambdaEvent{}
}
