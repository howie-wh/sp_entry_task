package constant

// GetWhiteList 获取url白名单
func GetWhiteList() []string {
	return []string{
		"/api/entry_task/v1/activity/user_register",
		"/api/entry_task/v1/activity/login",
		"/api/entry_task/v1/activity/user_list",
		"/api/entry_task/v1/activity/user_info",
		"/api/entry_task/v1/activity/activity_list",
		"/api/entry_task/v1/activity/activity_info",
		"/api/entry_task/v1/activity/activity_type_list",
		"/api/entry_task/v1/activity/comment_list",
	}
}

// GetBlackList 获取url黑名单
func GetBlackList() []string {
	return []string{}
}
