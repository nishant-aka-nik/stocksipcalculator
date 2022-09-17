package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type RuleRepository struct {
	client *mongo.Database
	ctx    context.Context
}

func NewRuleRepository(ctx context.Context, client *mongo.Database) *RuleRepository {
	return &RuleRepository{
		client: client,
		ctx:    ctx,
	}
}

func (repository *RuleRepository) AddRule(rules []interface{}) error {
	coll := repository.client.Collection("rules")
	// var DBRules  //sending it to db
	_, err := coll.InsertMany(repository.ctx, rules)
	if err != nil {
		panic(err)
	}
	return nil
}
