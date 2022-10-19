package api

type User struct {
	ID      string
	Subject string
	Email   string
	Groups  []string
	Issuer  string
}
