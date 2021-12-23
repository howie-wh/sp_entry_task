package constant

// GetWhiteList 获取url白名单
func GetWhiteList() []string {
	return []string{
		"/api/entry_task/v1/admin/login",
	}
}

// GetBlackList 获取url黑名单
func GetBlackList() []string {
	return []string{}
}
