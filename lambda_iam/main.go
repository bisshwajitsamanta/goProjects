package main

import (
	"fmt"
	"lambda_iam/internal/lambda"
)

type App struct {
	Keycheck *lambda.LambdaEvent
}

func Run() error {
	fmt.Println("Starting the Lambda Application")
	iamKeyCheck := lambda.NewLambdaService()
	app := App{
		Keycheck: iamKeyCheck,
	}
	err := app.Keycheck.Connect()
	if err != nil {
		return err
	}

	err = app.Keycheck.Access()
	if err != nil {
		return err
	}

	err = app.Keycheck.ListUser()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up the Lambda Application")
		fmt.Println(err)
	}
}
