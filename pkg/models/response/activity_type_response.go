package response

// ActivityTypeResponse 用户实体返回结构体
type ActivityTypeResponse struct {
	TypeId      uint64         `json:"typeId"`           	//活动类型ID
	TypeName    string         `json:"typeName"`            //活动类型名称
}
