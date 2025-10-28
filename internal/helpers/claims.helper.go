package helpers

import (
	"github.com/golang-jwt/jwt/v5"
)

func HasPermission(claims jwt.MapClaims, required int) bool {
	rawPerms, ok := claims["permissions"].([]any)
	if !ok {
		return false
	}
	for _, p := range rawPerms {
		if int(p.(float64)) == required {
			return true
		}
	}
	return false
}
