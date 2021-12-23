package base

// GlobalPageQuery 全局Query通用条件
type GlobalPageQuery struct {
	Start   int    `form:"start"`   //当前页码
	Offset  int    `form:"offset"`  //显示条数
}
