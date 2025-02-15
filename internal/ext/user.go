package ext

import (
	"github.com/gouniverse/userstore"
	"github.com/samber/lo"
)

func IsClient(user userstore.UserInterface) bool {
	return user.Meta("is_client") == "yes"
}

func SetIsClient(user userstore.UserInterface, isClient bool) userstore.UserInterface {
	value := lo.Ternary(isClient, "yes", "no")
	user.SetMeta("is_client", value)
	return user
}
