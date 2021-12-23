package req

// ActivityJoinBody 用于活动报名或者活动退出
type ActivityJoinBody struct {
	JoinId      uint64     	`json:"joinId"` 	    				//报名ID
	ActivityId  uint64     	`json:"activityId" binding:"required"` 	//活动ID
	UserId  	uint64     	`json:"userId"` 						//用户ID
	DelFlag     string      `json:"delFlag"`        				//删除标志
}
