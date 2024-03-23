package repositories

import (
	"authentication-service/internal/canonical"
	"authentication-service/internal/config"
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	GetUser(ctx context.Context, login canonical.Login) (*canonical.User, error)
	CreateUser(context.Context, canonical.User) error
}

const (
	tableName         = "users"
	indexRegistration = ""
	indexUserName     = ""
)

type repository struct {
	database  *dynamodb.Client
	tableName string
}

func New() Repository {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.Get().AWS.AccessKeyId, config.Get().AWS.SecretAccessKey, config.Get().AWS.SessionToken)), awsconfig.WithRegion(config.Get().AWS.Region),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("an error occurred when connect to the database")
	}
	return &repository{
		database:  dynamodb.NewFromConfig(cfg),
		tableName: tableName,
	}
}

func (r *repository) CreateUser(ctx context.Context, user canonical.User) error {
	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	_, err = r.database.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      av,
		TableName: &r.tableName,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUser(ctx context.Context, login canonical.Login) (*canonical.User, error) {
	var queryItem, index, valueToFind string

	if login.Registration != "" {
		queryItem = "registration"
		index = ""
		valueToFind = login.Registration
	}

	if login.UserName != "" {
		queryItem = "user_name"
		index = ""
		valueToFind = login.UserName
	}

	result, err := r.database.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		IndexName:              aws.String(index),
		KeyConditionExpression: aws.String(fmt.Sprintf("#%s = :%s", queryItem, queryItem)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			fmt.Sprintf(":%s", queryItem): &types.AttributeValueMemberS{Value: valueToFind},
		},
		ExpressionAttributeNames: map[string]string{
			fmt.Sprintf("#%s", queryItem): queryItem,
		},
	})
	if err != nil {
		return nil, err
	}

	var user canonical.User
	for _, item := range result.Items {
		if err := attributevalue.UnmarshalMap(item, &user); err != nil {
			return nil, err
		}
	}

	return &user, nil
}
