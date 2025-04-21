package app

type LoginInfo struct {
	Provider string
	Email    string
	Pass     string
}

func Login() LoginInfo {
	provider := Input("provider [apple|gmail]: ")
	email := Input("email: ")
	pass := Input("pass: ")

	return LoginInfo{
		provider,
		email,
		pass,
	}
}
