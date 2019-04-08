package main

import (
	"fmt"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)


func main()  {
	// 直接使用 accesskey 和 secretkey
	//creds := credentials.NewStaticCredentialsFromCreds(credentials.Value{
	//	AccessKeyID:     *aws.String("youraccesskey"),
	//	SecretAccessKey: *aws.String("yourecretkey"),
	//})
	//sess,err:= session.NewSession(&aws.Config{
	//	Region: aws.String("us-east-1"),
	//	Credentials: creds,
	//})
	
	// 使用 Instance Role,不需要在服务器上配置 aws profile 和相关认证的环境变量
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	svc := secretsmanager.New(sess)
	input := &secretsmanager.ListSecretsInput{}
	
	result, err := svc.ListSecrets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidNextTokenException:
				fmt.Println(secretsmanager.ErrCodeInvalidNextTokenException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	
	fmt.Println(result)
}
