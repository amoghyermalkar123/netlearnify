package Student

// "gopkg.in/mgo.v2/bson"
// "netlui-go-server/conn"

// MODEL

type RegisterModel struct {
	UserId   int    `bson:"user_id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Year     int    `bson:"year"`
	Branch   string `bson:"branch"`
}
