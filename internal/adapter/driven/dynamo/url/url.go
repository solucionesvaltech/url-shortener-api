package url

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"url-shortener/internal/adapter/driven/dynamo"
	"url-shortener/internal/core/domain"
	"url-shortener/pkg/config"
	"url-shortener/pkg/log"
)

// Repository implements URLRepository
type Repository struct {
	db        dynamo.ClientInterface
	tableName string
}

// NewURLRepository initialize the repository with the DynamoDB client and the table name
func NewURLRepository(db dynamo.ClientInterface, conf *config.Config) (*Repository, error) {
	return &Repository{
		db:        db,
		tableName: conf.DatabasesConfig.DynamoDB.TableName,
	}, nil
}

// Save saves a new record in the `urls` table
func (r *Repository) Save(url domain.URL) error {
	item, err := dynamodbattribute.MarshalMap(url)
	if err != nil {
		return fmt.Errorf("error serializing URL: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	}

	log.Log.Debugf("going to save: %v", item)
	_, err = r.db.PutItem(input)
	if err != nil {
		return fmt.Errorf("error saving URL: %w", err)
	}

	return nil
}

// Find lookup the url in the `urls` table using `shortID`
func (r *Repository) Find(shortID string) (*domain.URL, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"short": {
				S: aws.String(shortID),
			},
		},
	}

	result, err := r.db.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("error when searching for URL: %w", err)
	}
	if result.Item == nil {
		return nil, fmt.Errorf("URL with shortID: %s not found", shortID)
	}

	var url domain.URL
	err = dynamodbattribute.UnmarshalMap(result.Item, &url)
	if err != nil {
		return nil, fmt.Errorf("error deserializing URL: %w", err)
	}

	return &url, nil
}

// Update updates an existing URL in the `urls` table
func (r *Repository) Update(url domain.URL) error {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"short": {
				S: aws.String(url.Short),
			},
		},
		UpdateExpression: aws.String("SET original = :original, enabled = :enabled"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":original": {
				S: aws.String(url.Original),
			},
			":enabled": {
				BOOL: aws.Bool(url.Enabled),
			},
		},
	}

	_, err := r.db.UpdateItem(input)
	if err != nil {
		return fmt.Errorf("error updating URL: %w", err)
	}

	return nil
}
