package aws

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
)
type IHCommunityAWS struct {
	sess *session.Session
}

// NewSession 创建一个新的 AWS 会话
func Run() *session.Session {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

    // 创建 AWS 会话
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2"),
        Credentials: credentials.NewStaticCredentials("accessKey", "secretKey", ""),
    })
    if err != nil {
        return nil, err
    }
    return sess, nil
}