package v1

import (
	"errors"

	"../../config"
	"../../model"
	"../../service/userService"
	"../../service/userService/userPermission"
	"../../util/crypto"
	"../../util/email"
	"../../util/random"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitAuth inits user authorization apis
// @Title Auth
// @Description Auth's router group.
func InitAuth(parentRoute *echo.Group) {
	parentRoute.POST("/login", loginUser)
	parentRoute.POST("/register", registerUser)

	parentRoute.POST("/forgotPassword", forgotPassword)
	parentRoute.PUT("/verifyCode", verifyCode)

	parentRoute.Use(middleware.JWT([]byte(config.AuthTokenKey)))
	parentRoute.GET("/me", readMe)
}

// UserForm struct.
type UserForm struct {
	Token string            `json:"token"`
	User  *model.PublicUser `json:"user"`
}

// @Title loginUser
// @Description Login a user.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"User Email."
// @Param   password	form   string 	true	"User Password."
// @Success 200 {object} UserForm 				"Returns login user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.incorrect"
// @Failure 400 {object} response.BasicResponse "err.user.token"
// @Resource /login
// @Router /login [post]
func loginUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(&user); err != nil {
		return response.KnownErrJSON(c, "err.user.bind", err)
	}

	// generate hash password
	user.Password = crypto.GenerateHash(user.Password)
	// check user crediential
	user, err := userService.LoginByInfo(user)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.incorrect", errors.New("Incorrect email or password"))
	}

	// generate encoded token and send it as response.
	t, err := userPermission.GenerateToken(user)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public user
	publicUser := &model.PublicUser{User: user}
	return response.SuccessInterface(c, UserForm{t, publicUser})
}

// @Title registerUser
// @Description Register a user.
// @Accept  json
// @Produce	json
// @Param   firstname	form   string   true	"User Firstname."
// @Param   lastname   	form   string   true	"User Lastname."
// @Param   email       form   string   true	"User Email."
// @Param   password	form   string 	true	"User Password."
// @Success 200 {object} UserForm				"Returns registered user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.exist"
// @Failure 400 {object} response.BasicResponse "err.user.create"
// @Failure 400 {object} response.BasicResponse "err.user.token"
// @Resource /register
// @Router /register [post]
func registerUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return response.KnownErrJSON(c, "err.user.bind", err)
	}

	// check existed email
	if _, err := userService.ReadUserByEmail(user.Email); err == nil {
		return response.KnownErrJSON(c, "err.user.exist",
			errors.New("Same email is existed. Please input other email"))
	}

	// create user with registered info
	user, err := userService.CreateUser(user)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.create", err)
	}

	// generate encoded token and send it as response.
	t, err := userPermission.GenerateToken(user)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public user
	publicUser := &model.PublicUser{User: user}
	return response.SuccessInterface(c, UserForm{t, publicUser})
}

// @Title forgotPassword
// @Description Forgot Password.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"User Email."
// @Success 200 {object} string					"Returns created user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.read"
// @Resource /forgotPassword
// @Router /forgotPassword [post]
func forgotPassword(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return response.KnownErrJSON(c, "err.user.bind", err)
	}

	// check user crediential
	user, err := userService.ReadUserByEmail(user.Email)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.read", err)
	}

	// generate verify code to reset password
	verifyCode := random.GenerateRandomDigitString(6)
	userService.UpdateVerifyCode(user.ID, verifyCode)

	// send forgot email to user email
	go email.SendForgotEmail(user, verifyCode)

	return response.SuccessJSON(c, "Server has sent email to you. Please check your email and reset password.")
}

// @Title verifyCode
// @Description Verify code.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"User Email."
// @Success 200 {object} string					"Returns token to reset password"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.read"
// @Resource /verifyCode
// @Router /verifyCode [post]
func verifyCode(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return response.KnownErrJSON(c, "err.user.bind", err)
	}

	// check email with verify code
	user, err := userService.CheckVerifyCode(user.Email, user.Verify.Code)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.verify", err)
	}

	// Generate encoded token and send it as response.
	t, err := userPermission.GenerateToken(user)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.token", err)
	}

	return response.SuccessInterface(c, t)
}

// @Title readMe
// @Description Read self profile information.
// @Accept  json
// @Produce	json
// @Success 200 {object} model.PublicUser		"Returns self profile information"
// @Failure 400 {object} response.BasicResponse "err.user.read"
// @Resource /me
// @Router /me [get]
func readMe(c echo.Context) error {
	objid := userPermission.IDFromToken(c)
	// read user by objid from token
	user, err := userService.ReadUser(objid)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.read", err)
	}
	// retrieve to public
	publicUser := &model.PublicUser{User: user}
	return response.SuccessInterface(c, publicUser)
}
