package setting

type password struct {
	AdminPwd string
	ROnlyPwd string
}

var pw *password

func GetPassword() *password {
	return pw
}

func init() {
	pw = &password{}
}
