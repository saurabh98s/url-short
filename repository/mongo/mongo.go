package mongo

import (
	"context"
	"time"
	"url-short/shortener"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}


func newMongoClient(mongoURL string,mongoTimeout int) (*mongo.Client,error) {
	ctx,cancel:=context.WithTimeout(context.Background(),time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client,err:= mongo.Connect(ctx,options.Client().ApplyURI(mongoURL))
	if err!=nil {
		return nil,err
	}
	err=client.Ping(ctx,readpref.Primary())
	if err!=nil {
		return nil,err
	}
	return client,nil
}

func NewMongoRepository(mongoURL,mongoDB string, mongoTimeout int)(shortener.RedirectRepository,error){
	repo:= &mongoRepository{
		timeout: time.Duration(mongoTimeout)*time.Second,
		database: mongoDB,
	}
	client,err:=newMongoClient(mongoURL,mongoTimeout)
	if err != nil {
		return nil,errors.Wrap(err,"[ERROR] at repository.NewMongoRepository")
	}
	repo.client=client
	return repo,nil //wont work unless you attach the interface aka the two methods required by the interface.
}



