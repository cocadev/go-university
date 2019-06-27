package userPermission

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"../../../api/response"
	"../../../config"
	"../../../model"
	"../../../util/log"
	"../../../util/timeHelper"
	"../../userService"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GenerateToken returns token after generate with user
func GenerateToken(user *model.User) (string, error) {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["idx"] = user.ID
	claims["exp"] = timeHelper.FewDaysLater(config.AuthTokenExpirationDay)

	// generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AuthTokenKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// IDFromToken returns idx from token
func IDFromToken(c echo.Context) bson.ObjectId {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	idx := claims["idx"].(string)
	return bson.ObjectId(idx)
}

// AuthRequired run function when user logged in.
func AuthRequired(f func(c echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, exp, err := userService.CurrentUser(c)
		if err != nil {
			log.Error("Auth failed.")
			response.KnownErrorJSON(c, http.StatusUnauthorized, "error.auth.fail", errors.New("Auth failed"))
			return nil
		}

		if exp {
			log.Error("Token is expired.")
			response.KnownErrorJSON(c, http.StatusUnauthorized, "error.token.expire", errors.New("Token is expired"))
			return nil
		}

		f(c)
		return nil
	}
}
