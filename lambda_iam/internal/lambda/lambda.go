package lambda

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type Service interface {
	Connect() error
	Access() error
}

type LambdaEvent struct {
	sess *session.Session
	svc  *iam.IAM
}

func (l *LambdaEvent) Connect() error {
	fmt.Println("Connected to Lambda !!")
	var err error
	l.sess, err = session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return err
	}
	l.svc = iam.New(l.sess)
	return nil
}

func (l *LambdaEvent) Access() error {
	return nil
}

func NewLambdaService() *LambdaEvent {
	return &LambdaEvent{}
}
