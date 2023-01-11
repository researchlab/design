package proxy

import (
	"log"
	"time"
)

// 静态代理在不改变原有接口的前提下， 做一些额外的工作比如监控，统计, 逻辑过滤等操作
// 代理类一般需要实现被代理类的部分/所有方法

// IUser IUser
type IUser interface {
	Login(username, password string) error
}

// User 用户
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// do something
	return nil
}

// UserProxy 代理类
type UserProxy struct {
	user *User
}

// NewUserProxy NewUserProxy
func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

// Login same with user Login
func (p *UserProxy) Login(username, password string) error {
	// before 这里可能会有一些统计逻辑
	start := time.Now()

	// 这里是原有的业务逻辑
	if err := p.user.Login(username, password); err != nil {
		return err
	}

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s\n", time.Now().Sub(start))
	return nil
}
