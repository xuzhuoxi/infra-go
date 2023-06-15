package mgox

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"testing"
	"time"
)

// Using the driver "github.com/mongodb/mongo-go-driver"
func TestGitHudbMongoDriver(t *testing.T) {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//client, err := mongo.Connect(ctx, "mongodb://192.168.3.105:27017")
	//if nil != err {
	//	logx.Warnln(err)
	//}
	//println(client)
}

func TestMgo(t *testing.T) {
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.3.105:27017"))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer client.Disconnect(ctx)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ctxVal, _ := context.WithTimeout(context.Background(), 20*time.Second)
	res, err := client.Database("TestDB").Collection("TestC").InsertOne(ctxVal, bson.M{"name": "go.mongodb.org", "value": "InsertOne"})
	if nil != err {
		t.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(id)

	list, err := client.ListDatabaseNames(context.Background(), bsonx.Doc{})
	if nil != err {
		t.Fatal(err)
	}
	fmt.Println(list)

	mongo.NewClient()

}
