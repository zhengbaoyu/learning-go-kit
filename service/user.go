package service

type IUserService interface {
	GetUserName(uid int) string
}

type UserSevice struct {
}

func (p *UserSevice) GetUserName(uid int) string {
	if uid == 100 {
		return "hello"
	}
	return "work"
}
