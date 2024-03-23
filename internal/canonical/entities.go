package canonical

type Login struct {
	UserName     string
	Registration string
	Password     string
}

type User struct {
	Id           int    `dynamodbav:"id"`
	UserName     string `dynamodbav:"user_name"`
	Registration string `dynamodbav:"registration"`
	Email        string `dynamodbav:"email"`
	Password     string `dynamodbav:"password"`
}
