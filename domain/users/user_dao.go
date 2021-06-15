package users

import (
	"fmt"
	"main/utils/errors"
)

var(
	userDb = make(map[int64]*User)
)
//func something() {
//	user := User{}
//	if err := user.Get(); err!=nil{
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(user.FirstName)
//}
func (user *User)Get() *errors.RestError{
	result := userDb[user.Id]
	if result == nil{
		return  errors.NewNotFoundError(fmt.Sprintf("user %d Not found", user.Id))
	}
	user.Id=result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return  nil
}
func (user *User) Save() *errors.RestError{
	if userDb[user.Id]!=nil{
		if  userDb[user.Id].Email==user.Email {
			return errors.NewBadReqError(fmt.Sprintf("User Email %s already registered", user.Email))
			
		}
		return errors.NewBadReqError(fmt.Sprintf("User %d already exists", user.Id))
	}
	userDb[user.Id] = user
	return nil
}
