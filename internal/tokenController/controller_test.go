package tokenController

import (
	"auth/internal/model"
	"testing"
)

var testController = JWTController{}

func TestCreateToken(t *testing.T) {
	testUsers := []model.User{
		{ID: "asdf-wewfs-234s-as", Login: "diPhantxm", Password: ""},
		{ID: "someIdToCheck", Login: "testCase", Password: ""},
		{ID: "thisIsIdForTest", Login: "admin", Password: ""},
	}

	controller := testController

	for _, user := range testUsers {
		_, err := controller.Create(user)

		if err != nil {
			t.Errorf("Token was not created. Error: %s\n", err.Error())
		}
	}
}

func TestVerifyToken(t *testing.T) {
	testCases := []struct {
		token    string
		expected bool
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ5NDI2OTYsImlkIjoiYXNkZi13ZXdmcy0yMzRzLWFzIiwibG9naW4iOiJkaVBoYW50eG0ifQ.yhUG-o6yKIKwOEMOhuV3bUTvYTYwLsAlIu_y4Y2bU50", true},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ5NDI0MjUsImlkIjoic29tZUlkVG9DaGVjayIsImxvZ2luIjoidGVzdENhc2UifQ.sPTzthofvRDtjEfaV7SCCb4DVEeyi6bDJl-6_pg7bOA", true},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTU3NjY1MzQsImlkIjoic29tZUlkVG9DaGVjayIsImxvZ2luIjoidGVzdENhc2UifQ.WoPoWJsBouQDZm2kqTnOZNTgCsYewNjP2qEeCcR_gx4", false},
	}

	controller := testController

	for _, testCase := range testCases {
		actual, _ := controller.Verify(testCase.token)
		if actual != testCase.expected {
			t.Errorf("Got wrong result on token '%s'. Expected: %v\n", testCase.token, testCase.expected)
		}
	}
}
