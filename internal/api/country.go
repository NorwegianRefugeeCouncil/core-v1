package api

type Country struct {
	ID   string `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}

type UserProfile struct {
	UserID  string `db:"user_id"`
	Issuer  string `db:"issuer"`
	Subject string `db:"subject"`
	Email   string `db:"email"`
}
