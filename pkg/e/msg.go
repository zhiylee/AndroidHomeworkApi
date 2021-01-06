package e

var MsgFlags = map[int]string {
	SUCCESS : "success",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",

	TIP_NO_ACTICLE: "没有文章",

	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",

	ERROR_USER_EXISTS: "用户名已存在",
	ERROR_LOGIN_FAIL: "用户名或密码错误",
	ERROR_USER_NO_EXIST: "用户不存在",

	ERROR_UPDATE_USER_INFO_FAIL: "更新用户信息失败",

	ERROR_NO_FAVORITE: "没有收藏的文章",
	ERROR_FAVORITE_EXSIT: "收藏的文章已存在",
	ERROR_FAVORITE_NO_EXIT: "该文章不在收藏夹中",

	ERROR_NO_COMMENT: "没有评论",

	ERROR_NO_FRIEND: "没有好友",
	ERROR_ADD_FRIEND_NO_EXIST: "所添加的用户不存在",
	ERROR_ALREADY_FRIEND: "已经是好友",
	ERROR_ADD_FRIEND_SELF: "不能添加自己为好友",

	ERROR_NO_TEST_PAPER: "没有试卷",
	ERROR_NO_EXIST_TEST_PAPER: "试卷不存在",

	ERROR_NO_TEST_RECORDS: "没有测试记录",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}