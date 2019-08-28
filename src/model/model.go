package model

// User type
type User struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string `json:"username" bson:"username"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Password  string `json:"password" bson:"password"`
	Token     string `json:"token" bson:"token"`
}

// ResponseResult for model
type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
