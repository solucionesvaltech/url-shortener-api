package url

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"time"
	"url-shortener/internal/adapter/driven/dynamo"
	"url-shortener/pkg/log"
)

// ensureURLTableExists checks if the table exists; If not, create it
func (r *Repository) ensureURLTableExists() error {
	_, err := r.db.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(r.tableName),
	})

	if err == nil {
		log.Log.Debugf("table: %s already exists", r.tableName)
		return nil
	}

	if isResourceNotFoundException(err) {
		_, err = r.db.CreateTable(&dynamodb.CreateTableInput{
			TableName: aws.String(r.tableName),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("short"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("short"),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})
		if err != nil {
			return fmt.Errorf("error creating table %s: %w", r.tableName, err)
		}

		fmt.Printf("table %s successfully created, waiting for it to be active...", r.tableName)
		return waitForTableActive(r.db, r.tableName)
	}

	return fmt.Errorf("error checking table %s: %w", r.tableName, err)
}

// isResourceNotFoundException check if the error is due to a missing table
func isResourceNotFoundException(err error) bool {
	var awsError awserr.Error
	if errors.As(err, &awsError) && awsError.Code() == dynamodb.ErrCodeResourceNotFoundException {
		return true
	}
	return false
}

// waitForTableActive wait for the table to be active
func waitForTableActive(db dynamo.ClientInterface, tableName string) error {
	for {
		resp, err := db.DescribeTable(&dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})
		if err != nil {
			return fmt.Errorf("error checking table status: %w", err)
		}
		if *resp.Table.TableStatus == dynamodb.TableStatusActive {
			log.Log.Debugf("table: %s is active", tableName)
			return nil
		}
		time.Sleep(2 * time.Second)
	}
}
