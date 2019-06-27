package userService

import (
	"time"

	"../../model"
	"../../util/log"
	"../../util/timeHelper"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/now"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// LoginByInfo returns user with email and password
func LoginByInfo(user *model.User) (*model.User, error) {
	userCollection, session := userCollection()
	defer session.Close()

	err := userCollection.Find(bson.M{"email": user.Email, "password": user.Password}).One(user)
	return user, err
}

// ReadUserByEmail returns user
func ReadUserByEmail(email string) (*model.User, error) {
	userCollection, session := userCollection()
	defer session.Close()

	user := &model.User{}
	err := userCollection.Find(bson.M{"email": email}).Select(bson.M{}).One(user)
	return user, err
}

// CurrentUser get a current user.
func CurrentUser(c echo.Context) (*model.User, bool, error) {
	userCollection, session := userCollection()
	defer session.Close()

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	idx := claims["idx"].(string)
	exp, _ := time.Parse(time.RFC3339, claims["exp"].(string))

	if len(idx) == 0 {
		log.Debug("header token not exist.")
		return nil, false, nil
	}

	isExpired := timeHelper.IsExpired(exp)
	log.Debugf("expired : %t", isExpired)
	if isExpired {
		return nil, true, nil
	}

	objid := bson.ObjectIdHex(idx)
	result, err := userCollection.Find(bson.M{"_id": objid, "token": token.Raw}).Count()
	if err != nil || result == 0 {
		return nil, true, err
	}

	user := model.User{}
	return &user, false, nil
}

// UpdateVerifyCode update verify code for forgot password
func UpdateVerifyCode(objid bson.ObjectId, verifyCode string) error {
	userCollection, session := userCollection()
	defer session.Close()

	// update verify code with object id
	return userCollection.UpdateId(objid, bson.M{
		"$set": bson.M{
			"verify.code":      verifyCode,
			"verify.createdAt": timeHelper.GetCurrentTime(),
			"updatedAt":        timeHelper.GetCurrentTime(),
		},
	})
}

// CheckVerifyCode checks that exists email and verifyCode
func CheckVerifyCode(email string, verifyCode string) (*model.User, error) {
	userCollection, session := userCollection()
	defer session.Close()

	// calculate over time
	timeup := timeHelper.GetCurrentTime()
	timeup = timeup.Add(-time.Minute)
	timeup = now.New(timeup).BeginningOfMinute()

	// check verify code with email
	user := &model.User{}
	err := userCollection.Find(bson.M{
		"email":            email,
		"verify.code":      verifyCode,
		"verify.createdAt": bson.M{"$gte": timeup},
	}).One(user)

	return user, err
}
