package api

import (
	"net/url"
	"strconv"
)

type GetAllUsersOptions struct {
	Email      string
	Take       int
	Skip       int
	CountryIDs []string
}

func ParseGetAllUserOptions(v url.Values) (GetAllUsersOptions, error) {
	ret := GetAllUsersOptions{}
	ret.Email = v.Get("email")
	takeStr := v.Get("take")
	if len(takeStr) != 0 {
		take, err := strconv.Atoi(takeStr)
		if err != nil {
			return ret, err
		}
		ret.Take = take
	}
	skipStr := v.Get("skip")
	if len(skipStr) != 0 {
		skip, err := strconv.Atoi(skipStr)
		if err != nil {
			return ret, err
		}
		ret.Skip = skip
	}
	ret.CountryIDs = v["country"]
	return ret, nil
}

type User struct {
	ID      string `db:"id"`
	Subject string `db:"subject"`
	Email   string `db:"email"`
}
