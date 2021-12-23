package page

import (
	"entry_task/pkg/logger"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

// Page 分页结构体
type Page struct {
	Total int64       `json:"total"` //总条数
	List  interface{} `json:"list"`  //数据
}

// GetTotal 获取总条数
func GetTotal(engine *xorm.Session, args ...interface{}) (int64, error) {
	if args != nil {
		engine.Table(args)
	}
	count, err := engine.Count()
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return 0, err
	}
	return count, nil
}
