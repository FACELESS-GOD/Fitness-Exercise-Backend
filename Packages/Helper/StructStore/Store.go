package StructStore

type UserData struct {
	UserName           string
	Password           string
	Email              string
	FirstName          string
	MiddleName         string
	LastName           string
	Designation        int64
	AuthorizationId    int64
	IsValid            bool
	Last_Modified_Date string
}

type UserAuth struct {
	UserName string
	Password string
}

type ValidateUserResponse struct {
	UserName string
}
