package ext

import (
	"strings"

	"github.com/gouniverse/userstore"
	"github.com/samber/lo"
)

func DisplayNameFull(user userstore.UserInterface) string {
	if user == nil {
		return "n/a"
	}

	displayName := user.FirstName() + " " + user.LastName()

	if strings.TrimSpace(displayName) == "" {
		return user.Email()
	}

	return displayName
}

func IsClient(user userstore.UserInterface) bool {
	return user.Meta("is_client") == "yes"
}

func SetIsClient(user userstore.UserInterface, isClient bool) userstore.UserInterface {
	value := lo.Ternary(isClient, "yes", "no")
	user.SetMeta("is_client", value)
	return user
}
