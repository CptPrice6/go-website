package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id          int       `orm:"pk;auto"`
	Email       string    `orm:"unique;size(255)"`
	Password    string    `orm:"size(255)"`
	Name        string    `orm:"size(100)"`
	Surname     string    `orm:"size(100)"`
	Description string    `orm:"type(text);null"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	Role        string    `orm:"size(20)"`
	Ban         bool      `orm:"default(false)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

func CreateUser(email, password, role string, ban bool) error {
	o := orm.NewOrm()
	user := User{
		Email:    email,
		Password: password,
		Role:     role,
		Ban:      ban,
	}

	_, err := o.Insert(&user)
	return err
}

func GetUserByEmail(email string) (*User, error) {
	o := orm.NewOrm()
	user := User{Email: email}
	err := o.Read(&user, "Email")
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id int) (*User, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user, "Id")
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsUserBanned(userID int) (bool, error) {
	o := orm.NewOrm()
	user := User{Id: userID}

	err := o.Read(&user, "Id")
	if err == orm.ErrNoRows {
		return false, nil // User not found, so not banned
	} else if err != nil {
		return false, err // Some other error occurred
	}

	return user.Ban, nil // Return the ban status
}

func BanUserByID(userID int) error {
	o := orm.NewOrm()
	user := User{Id: userID}

	if err := o.Read(&user); err != nil {
		return err
	}

	// Update the ban status
	user.Ban = true
	_, err := o.Update(&user, "Ban")
	return err
}

func UpdateUser(user *User) error {
	o := orm.NewOrm()

	// Update user
	_, err := o.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserByID(id int) error {
	o := orm.NewOrm()

	user := User{Id: id}

	// Delete the user by ID
	_, err := o.Delete(&user)
	if err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]User, error) {
	o := orm.NewOrm()

	var users []User

	_, err := o.QueryTable(new(User)).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
