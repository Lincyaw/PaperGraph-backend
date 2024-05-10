package handlers

type Response struct {
	Code              int         `json:"code,omitempty"`
	ErrMessage        string      `json:"err_message,omitempty"`
	UserFriendMessage string      `json:"user_friend_message,omitempty"`
	Detail            interface{} `json:"detail,omitempty"`
}

func NewResponse(detail interface{}) Response {
	return Response{
		Detail: detail,
	}
}

func NewErrResponse(err error, explanation string, detail interface{}) Response {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return Response{
		ErrMessage:        errMsg,
		Detail:            detail,
		UserFriendMessage: explanation,
	}
}
