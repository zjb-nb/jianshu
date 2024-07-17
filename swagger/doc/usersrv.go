package doc

// swagger:parameters userInfo
type User struct {
	//username of user info
	//
	// Required: true
	// in: path
	User string
}

// swagger:response usersrvResp
type Respon struct {
	Code int
	Data string
}

//swagger:route GET /users/{User}  user userInfo
// 获取用户信息.
// responses:
//   200: usersrvResp
