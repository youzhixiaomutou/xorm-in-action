package main

type xUser struct {
	Id       int64
	Name     string `xorm:"varchar(255) notnull unique 'user_name' comment('姓名')"`
	NickName string `xorm:"varchar(255) 'nick_name' comment('昵称')"`
	Email    string `xorm:"varchar(255) 'email' comment('邮件')"`
}
