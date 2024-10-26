package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"url-shortener/pkg/config"
)

//go:generate mockgen -source=dynamo.go -destination=../../../../mocks/dynamo.go -package=mock

// DynamoDBClient es la interfaz para interactuar con DynamoDB
type ClientInterface interface {
	DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
	CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
}

// Client es una implementación de DynamoDBClient
type Client struct {
	*dynamodb.DynamoDB
}

// NewDynamoDBClient configura y devuelve una conexión a DynamoDB
func NewDynamoDBClient(conf *config.Config) (*Client, error) {
	var awsConfig *aws.Config

	if conf.Stage == "production" {
		awsConfig = &aws.Config{
			Region: aws.String(conf.DatabasesConfig.DynamoDB.Region),
		}
	} else {
		awsConfig = &aws.Config{
			Region:   aws.String(conf.DatabasesConfig.DynamoDB.Region),
			Endpoint: aws.String(conf.DatabasesConfig.DynamoDB.Endpoint),
		}
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		DynamoDB: dynamodb.New(sess),
	}, nil
}
