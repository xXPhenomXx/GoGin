package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token Check Fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token Expired",
	ERROR_AUTH_TOKEN:                "Token Invalid",
	ERROR_AUTH:                      "Token Fail",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "Image Upload Fail",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "Image Check Fail",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "Image Invalid Format",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
