package entity

import (
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/security"
)

//// User 用户表
//type User struct {
//	UID        uint32      `gorm:"primary_key;column:uid;type:int unsigned auto_increment;comment:主键"`
//	Username   string      `gorm:"unique;column:username;type:varchar(100);not null;comment:用户名(唯一)"`
//	Password   string      `gorm:"column:password;type:varchar(100);not null;comment:登录密码"`
//	Email      string      `gorm:"column:email;type:varchar(50);not null;default:'';comment:邮箱号"`
//	Salt       string      `gorm:"column:salt;type:varchar(20);not null;default:'';comment:盐"`
//	Nickname   string      `gorm:"column:nickname;type:varchar(30);not null;default:'';comment:昵称"`
//	Avatar     string      `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:头像地址"`
//	Intro      string      `gorm:"column:intro;type:varchar(255);not null;default:'';comment:用户简介"`
//	WebSite    string      `gorm:"column:web_site;type:varchar(255);not null;default:'';comment:个人网站"`
//	UserType   pb.UserType `gorm:"column:user_type;type:enum('ADMIN','USER');not null;default:'USER';comment:用户类型 ADMIN-管理员 USER-普通用户"`
//	LoginType  uint8       `gorm:"column:login_type;type:tinyint unsigned;not null;default:0;comment:登录方式 1-用户名 2-GitHub 3-码云 4-QQ 5-微博"`
//	Status     uint8       `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:状态 0-正常 1-禁用"`
//	IsDeleted  BoolBit     `gorm:"column:is_deleted;type:bit(1);not null;default:b'0';comment:是否注销"`
//	CreateTime Timestamp   `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
//	UpdateTime Timestamp   `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
//	UnionID    string      `gorm:"column:union_id;type:varchar(50);not null;default:'';comment:用户唯一标识（第三方网站）"`
//}
//
//// TableName 设置 User 结构体对应的数据库表名
//func (User) TableName() string {
//	return "blog_user"
//}

type User struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string     `gorm:"size:100;not null;" json:"first_name"`
	LastName  string     `gorm:"size:100;not null;" json:"last_name"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Password  string     `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type PublicUser struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:100;not null;" json:"first_name"`
	LastName  string `gorm:"size:100;not null;" json:"last_name"`
}

// BeforeSave is a gorm hook
func (u *User) BeforeSave() error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

type Users []User

// PublicUsers So that we dont expose the user's email address and password to the world
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

// PublicUser So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func (u *User) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "email email"
			}
		}

	case "login":
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	default:
		if u.FirstName == "" {
			errorMessages["firstname_required"] = "first name is required"
		}
		if u.LastName == "" {
			errorMessages["lastname_required"] = "last name is required"
		}
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			errorMessages["invalid_password"] = "password should be at least 6 characters"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	}
	return errorMessages
}
