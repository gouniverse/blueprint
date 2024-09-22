package helpers

import (
	"project/config"

	"github.com/samber/lo"
)

// Untokenize accepts a map of key token pairs and returns a map of key value pairs
//
// Example:
//
//	keyTokenMap := map[string]string{
//	  "key1": "token1",
//	  "key2": "token2",
//	}
//
//	untokenizedMap, err := Untokenize(keyTokenMap)
//	if err != nil {
//	  return
//	}
//
//	fmt.Println(untokenizedMap)
//	// map[key1:value1 key2:value2]
//
// Parameters:
// - keyTokenMap (map[string]string): A map of key token pairs
//
// Returns:
// - untokenizedMap (map[string]string): A map of key value pairs
// - err (error): An error if one occurred
func Untokenize(keyTokenMap map[string]string) (map[string]string, error) {
	tokens := lo.Values(keyTokenMap)
	values, err := config.VaultStore.TokensRead(tokens, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading tokens", err.Error())
		return map[string]string{}, err
	}

	untokenized := lo.MapValues(keyTokenMap, func(token string, key string) string {
		return values[token]
	})

	return untokenized, nil
}
