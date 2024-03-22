package canonical

type Login struct {
	UserName     string
	Registration string
	Password     string
}

type User struct {
	Id           string
	UserName     string
	Registration string
	Email        string
	Password     string
}
