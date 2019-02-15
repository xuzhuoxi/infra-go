package mongox

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/xuzhuoxi/infra-go/logx"
	"testing"
	"time"
)

//Using the driver "gopkg.in/mgo.v2"
func TestMongoDriver(t *testing.T) {
	//driver := &MongoDriver{}
	//err := driver.Open(&mgo.DialInfo{Addrs: []string{"192.168.3.105"}}, mgo.Monotonic, 50)
	//fmt.Println(err)
	//s1, err := driver.NewSession(true)
	//fmt.Println(s1.CurrentSession().DatabaseNames())
	//fmt.Println(driver.NumSessions())
	//fmt.Println(driver.OriginSession().CurrentSession().DatabaseNames())
	//fmt.Println(driver.NumSessions())
}

//Using the driver "github.com/mongodb/mongo-go-driver"
func TestGitHudbMongoDriver(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://192.168.3.105:27017")
	if nil != err {
		logx.Warnln(err)
	}
	println(client)
}
