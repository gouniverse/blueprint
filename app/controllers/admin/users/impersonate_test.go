package admin

import (
	"net/http"
	"net/http/httptest"
	"project/config"
	"testing"
)

func TestImpersonate(t *testing.T) {
	config.TestsConfigureAndInitialize()

	// Setup
	userID := "test_user"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	// Act
	err := Impersonate(w, req, userID)

	// Assert
	if err != nil {
		config.LogStore.ErrorWithContext("At Impersonate", err.Error())
		t.Fail()
	}
}

// func TestUserStoreAuthToken(t *testing.T) {
// 	config.TestsConfigureAndInitialize()

// 	// Setup
// 	userID := uid.HumanUid()
// 	token := strutils.Random(32)
// 	err := UserStoreAuthToken(token, userID)

// 	if err != nil {
// 		config.LogStore.ErrorWithContext("At Impersonate", err.Error())
// 		t.Fail()
// 	}
// }
