package lambda

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"os"
	"time"
)

type Service interface {
	Connect() error
	Access() error
	ListUser() error
	RotateKeys() error
}

type LambdaEvent struct {
	sess        *session.Session
	svc         *iam.IAM
	KeyLastUsed *iam.GetAccessKeyLastUsedOutput
	KeysList    []string
}

//Connect - This method connects to IAM and get access keys
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
	return nil
}

// Next Goal is if the access key is last used more than say 200 days then rotate the keys

// Access - This method defines the logic if the key is not used for say
func (l *LambdaEvent) Access() error {
	date := time.Now().UTC()
	diff := date.Sub(*l.KeyLastUsed.AccessKeyLastUsed.LastUsedDate)
	fmt.Println("Number of Days spend after Last Used:", int(diff.Hours()/24), "days")
	return nil
}

//ListUser - This method defines how many access ID's does an user have and return to a list
func (l *LambdaEvent) ListUser() error {
	var (
		err error
	)
	l.sess, err = session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}
	l.svc = iam.New(l.sess)

	result, err := l.svc.ListAccessKeys(&iam.ListAccessKeysInput{
		MaxItems: aws.Int64(5),
		UserName: aws.String("s3admin"),
	})
	if err != nil {
		return err
	}
	// Goal is to have all the access keys in a list
	for _, v := range result.AccessKeyMetadata {
		l.KeysList = append(l.KeysList, *v.AccessKeyId)
	}
	fmt.Println(l.KeysList)
	return nil
}

func (l *LambdaEvent) RotateKeys() error {

	return nil
}

func NewLambdaService() *LambdaEvent {
	return &LambdaEvent{}
}
