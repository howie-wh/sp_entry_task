package response

// CommentResponse 用户实体返回结构体
type CommentResponse struct {
	CommentId   uint64      `json:"commentId"`               //ID
	ActivityId  uint64      `json:"activityId"`              //活动ID
	UserId      uint64      `json:"userId"`                  //活动类型ID
	Content     string    	`json:"content"`                 //活动内容
}
