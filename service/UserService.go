package service

import (
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

//通过登录名和密码获取userinfo
func UserInfoByLoginNameAndPassword(user *entity.User)( userInfo entity.User){
	userInfo = UserInfoByLoginName(user.LoginName)
	//将明文密码+查询出来的盐值md5，匹配查询出来的密码
	if restgo.MD5(user.Password + userInfo.Salt) == userInfo.Password {

		setNewSaltAndPassword(&userInfo)
		restgo.Db.Save(&userInfo)
		return
	}
	userInfo = entity.User{}
	return
}

//通过登录名获取用户信息
func UserInfoByLoginName(loginName string)(user entity.User){
	restgo.Db.Where("login_name = ?",loginName).Find(&user)
	return
}

//注册用户
func UserInfoOfCreate(user *entity.User)(affectedCount int64){
	userInfo := UserInfoByLoginName(user.LoginName)
	if userInfo.Id == 0 {
		setNewSaltAndPassword(user)
		affectedCount = restgo.Db.Create(&user).RowsAffected
	}

	return
}

//更新盐值和密码
func setNewSaltAndPassword(user *entity.User){
	user.Salt = restgo.GetRandomSalt()
	user.Password = restgo.MD5(user.Password + user.Salt)
}
