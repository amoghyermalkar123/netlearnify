package Student

// "gopkg.in/mgo.v2/bson"
// "netlui-go-server/conn"

// MODEL

type User struct {
	UserId int    `bson:"user_id"`
	Email  string `bson:"email"`
	Year   int    `bson:"year"`
	Branch string `bson:"branch"`
}
