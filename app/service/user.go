package service

import (
	"myapp/app/dao"
	"myapp/app/model"
)

// 中间件管理服务
var User = userService{}

type userService struct{}

func (s *userService) SignUp(r *model.UserServiceSignUpReq) (error, *model.Account) {
	// 昵称为非必需参数，默认使用账号名称
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}
	// 账号唯一性数据检查
	//if !s.CheckPassport(r.Passport) {
	//	return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Passport)),nil
	//}
	//// 昵称唯一性数据检查
	//if !s.CheckNickName(r.Nickname) {
	//	return errors.New(fmt.Sprintf("昵称 %s 已经存在", r.Nickname))
	//}

	//if _, err := dao.User.Save(r); err != nil {
	//	return err
	//}
	//return nil

	//account, _ := dao.Account.FindOne("account","zhangsan")
	var account model.Account
	var err error
	if err = dao.Account.Where("account", "zhangsan").Scan(&account);err != nil{
		return err,nil
	}
	return nil, &account




}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckPassport(passport string) bool {
	if i, err := dao.User.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}