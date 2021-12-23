package response

// ActivityResponse 实体返回结构体
type ActivityResponse struct {
	ActivityId  uint64      `json:"activityId"`              //活动ID
	TypeId      uint64      `json:"typeId"`                  //活动类型ID
	Title    	string    	`json:"title"`                   //活动标题
	Location    string    	`json:"location"`                //活动地点
	Content     string    	`json:"content"`                 //活动内容
	StartTime  	uint64 	    `json:"startTime"` 				 //开始时间
	EndTime  	uint64 	    `json:"endTime"`                 //结束时间
	IsJoin      bool        `json:"isJoin"`                  //是否参加，false-未参加 true-已参加
}

// ActivityTempResponse 临时表结构，用于联表查询获取返回结果
type ActivityTempResponse struct {
	ActivityId  uint64      `json:"activityId"`              //活动ID
	TypeId      uint64      `json:"typeId"`                  //活动类型ID
	Title    	string    	`json:"title"`                   //活动标题
	Location    string    	`json:"location"`                //活动地点
	Content     string    	`json:"content"`                 //活动内容
	StartTime  	uint64 	    `json:"startTime"` 				 //开始时间
	EndTime  	uint64 	    `json:"endTime"`                 //结束时间
	JoinId      uint64      `json:"joinId"`                  //报名id
}
