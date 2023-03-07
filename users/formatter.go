package users

type UserFormatter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Passion string `json:"passion"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

func FormatUser(users User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:      users.ID,
		Name:    users.Name,
		Passion: users.Passion,
		Email:   users.Email,
		Token:   token,
	}
	return formatter
}
