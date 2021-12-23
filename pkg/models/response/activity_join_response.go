package response

// ActivityJoinResponse 实体返回结构体
type ActivityJoinResponse struct {
	JoinId  	uint64      `json:"joinId"`              	 //活动报名ID
	ActivityId  uint64      `json:"activityId"`              //活动ID
	UserId      uint64      `json:"userId"`                  //用户ID
	DelFlag     string      `json:"delFlag"`				 //删除标志
}
