//go:build !skiptest

package mgox

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Using the driver "github.com/mongodb/mongo-go-driver"
func TestGithubDBMongoDriver(t *testing.T) {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//client, err := mongo.Connect(ctx, "mongodb://192.168.3.105:27017")
	//if nil != err {
	//	logx.Warnln(err)
	//}
	//println(client)
}

func TestMgo(t *testing.T) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://192.168.3.105:27017"))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

	ctxVal, _ := context.WithTimeout(context.Background(), 20*time.Second)
	res, err := client.Database("TestDB").Collection("TestC").InsertOne(ctxVal, bson.M{"name": "go.mongodb.org", "value": "InsertOne"})
	if nil != err {
		t.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(id)

	list, err := client.ListDatabaseNames(context.Background(), bson.D{})
	if nil != err {
		t.Fatal(err)
	}
	fmt.Println(list)
}
