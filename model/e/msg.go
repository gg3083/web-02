package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	LOGIN_FAILED: "用户名或密码错误",

	ERROR_EXIST_TAG:       "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:  "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:   "该标签不存在",
	ERROR_GET_TAGS_FAIL:   "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:  "统计标签失败",
	ERROR_ADD_TAG_FAIL:    "新增标签失败",
	ERROR_EDIT_TAG_FAIL:   "修改标签失败",
	ERROR_DELETE_TAG_FAIL: "删除标签失败",
	ERROR_EXPORT_TAG_FAIL: "导出标签失败",
	ERROR_IMPORT_TAG_FAIL: "导入标签失败",

	ERROR_NOT_EXIST_ACCEPTOR:        "该承兑人不存在",
	ERROR_ADD_ACCEPTOR_FAIL:         "新增承兑人失败",
	ERROR_DELETE_ACCEPTOR_FAIL:      "删除承兑人失败",
	ERROR_CHECK_EXIST_ACCEPTOR_FAIL: "检查承兑人是否存在失败",
	ERROR_EDIT_ACCEPTOR_FAIL:        "修改承兑人失败",
	ERROR_COUNT_ACCEPTOR_FAIL:       "统计承兑人失败",
	ERROR_GET_ACCEPTORS_FAIL:        "获取多个承兑人失败",
	ERROR_GET_ACCEPTOR_FAIL:         "获取单个承兑人失败",
	ERROR_GEN_ACCEPTOR_POSTER_FAIL:  "生成承兑人海报失败",

	ERROR_NOT_EXIST_ACCEPTOR_CARD:        "该承兑人卡信息不存在",
	ERROR_ADD_ACCEPTOR_CARD_FAIL:         "新增承兑人卡信息失败",
	ERROR_DELETE_ACCEPTOR_CARD_FAIL:      "删除承兑人卡信息失败",
	ERROR_CHECK_EXIST_ACCEPTOR_CARD_FAIL: "检查承兑人卡信息是否存在失败",
	ERROR_EDIT_ACCEPTOR_CARD_FAIL:        "修改承兑人卡信息失败",
	ERROR_COUNT_ACCEPTOR_CARD_FAIL:       "统计承兑人卡信息失败",
	ERROR_GET_ACCEPTOR_CARDS_FAIL:        "获取多个承兑人卡信息失败",
	ERROR_GET_ACCEPTOR_CARD_FAIL:         "获取单个承兑人卡信息失败",
	ERROR_GEN_ACCEPTOR_CARD_POSTER_FAIL:  "生成承兑人卡信息海报失败",

	ERROR_GET_USER_FAIL:    "获取单个用户失败",
	ERROR_NOT_EXIST_USER:   "用户不存在",
	ERROR_UPDATE_USER_FAIL: "更新用户信息失败",

	ERROR_VALIDATE_PASSWORD:       "密码错误",
	ERROR_UPDATE_PASSWORD:         "更新密码失败",
	ERROR_VALIDATE_TRADE_PASSWORD: "交易密码错误",
	ERROR_UPDATE_TRADE_PASSWORD:   "更新交易密码失败",

	ERROR_COUNT_BILL_FAIL: "统计账单失败",
	ERROR_GET_BILLS_FAIL:  "获取多个账单失败",

	ERROR_NOT_EXIST_CHANNEL:        "该通道信息不存在",
	ERROR_CHECK_EXIST_CHANNEL_FAIL: "检查通道信息是否存在失败",
	ERROR_ADD_CHANNEL_FAIL:         "新增通道信息失败",
	ERROR_DELETE_CHANNEL_FAIL:      "删除通道信息失败",
	ERROR_EDIT_CHANNEL_FAIL:        "更新通道信息失败",
	ERROR_COUNT_CHANNEL_FAIL:       "统计通道失败",
	ERROR_GET_CHANNELS_FAIL:        "获取多个通道失败",
	ERROR_GET_CHANNEL_FAIL:         "获取单个通道失败",

	ERROR_NOT_EXIST_ORDER:        "该订单信息不存在",
	ERROR_CHECK_EXIST_ORDER_FAIL: "检查订单信息是否存在失败",
	ERROR_ADD_ORDER_FAIL:         "新增订单信息失败",
	ERROR_DELETE_ORDER_FAIL:      "删除订单信息失败",
	ERROR_EDIT_ORDER_FAIL:        "更新订单信息失败",
	ERROR_COUNT_ORDER_FAIL:       "统计订单失败",
	ERROR_GET_ORDERS_FAIL:        "获取多个订单失败",
	ERROR_GET_ORDER_FAIL:         "获取单个订单失败",

	ERROR_NOT_EXIST_MATCH_CACHE:        "该匹配缓存信息不存在",
	ERROR_CHECK_EXIST_MATCH_CACHE_FAIL: "检查匹配缓存信息是否存在失败",
	ERROR_ADD_MATCH_CACHE_FAIL:         "新增匹配缓存信息失败",
	ERROR_DELETE_MATCH_CACHE_FAIL:      "删除匹配缓存信息失败",
	ERROR_EDIT_MATCH_CACHE_FAIL:        "更新匹配缓存信息失败",
	ERROR_COUNT_MATCH_CACHE_FAIL:       "统计匹配缓存失败",
	ERROR_GET_MATCH_CACHES_FAIL:        "获取多个匹配缓存失败",
	ERROR_GET_MATCH_CACHE_FAIL:         "获取单个匹配缓存失败",

	ERROR_NOT_EXIST_FUND:        "该资金信息不存在",
	ERROR_CHECK_EXIST_FUND_FAIL: "检查资金信息是否存在失败",
	ERROR_ADD_FUND_FAIL:         "新增资金信息失败",
	ERROR_DELETE_FUND_FAIL:      "删除资金信息失败",
	ERROR_EDIT_FUND_FAIL:        "更新资金信息失败",
	ERROR_COUNT_FUND_FAIL:       "统计资金失败",
	ERROR_GET_FUNDS_FAIL:        "获取多个资金失败",
	ERROR_GET_FUND_FAIL:         "获取单个资金失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
