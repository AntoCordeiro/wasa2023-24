package api

import (
	"errors"
	"strconv"
	"strings"
)

func GetUserID(authString string) (int, error) {
	// Check the Authorization header
	authStringComponents := strings.Split(authString, " ")
	if authStringComponents[0] != "Bearer" || len(authStringComponents) != 2 {
		return 0, errors.New("Invalid Authentication header")
	}

	// Convert and return the user id
	userID, err := strconv.Atoi(authStringComponents[1])
	if err != nil {
		return 0, err
	}

	return userID, nil
}
