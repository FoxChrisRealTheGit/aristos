package models
import(
	"time"
	"ARISTOS/dblayer/common/utility"
)

type User struct{
	UUID string `json:"uuid" bson:"uuid"`
	Username string `json:"username" bson:"username"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"firstName"`
	Email string `json:"email" bson:"bson"`
	PasswordHash string `json:"passwordHash" bson:"timestampCreated"`
	TimestampCreated int64 `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64 `json:"timestampModified" bson:"timestampModified"`
}

func NewUser(username string, firstName string, lastName, string, email string, password string) *User{
	passwordHash := utility.SHA256OfString(password)
	now := time.Now()
	unixTimestamp := now.Unix()
	u := User{UUID: utility.GenerateUUID(), Username: username, FirstName: firstName, LastName: lastName, Email: email, PasswordHash: passwordHash, TimestampCreated: unixTimestamp}
	return &u
}
