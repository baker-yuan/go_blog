package persistence

import (
	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/application/user/domain/repository"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// UserRepo 强制UserRepo实现repository.UserRepository接口
var _ repository.UserRepository = &UserRepo{}

// SaveUser 保存用户
func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	//dbErr := map[string]string{}
	//err := r.db.Debug().Create(&user).Error
	//if err != nil {
	//	// If the email is already taken
	//	if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
	//		dbErr["email_taken"] = "email already taken"
	//		return nil, dbErr
	//	}
	//	// any other db error
	//	dbErr["db_error"] = "database error"
	//	return nil, dbErr
	//}
	//return user, nil
	return nil, nil
}

// GetUser 通过用户ID查询用户
func (r *UserRepo) GetUser(id uint64) (*entity.User, error) {
	//var user entity.User
	//err := r.db.Debug().Where("id = ?", id).Take(&user).Error
	//if err != nil {
	//	return nil, err
	//}
	//if gorm.IsRecordNotFoundError(err) {
	//	return nil, errors.New("user not found")
	//}
	//return &user, nil

	return nil, nil
}

// GetUsers 获取所有用户
func (r *UserRepo) GetUsers() ([]entity.User, error) {
	//var users []entity.User
	//err := r.db.Debug().Find(&users).Error
	//if err != nil {
	//	return nil, err
	//}
	//if gorm.IsRecordNotFoundError(err) {
	//	return nil, errors.New("user not found")
	//}
	//return users, nil

	return nil, nil
}

// GetUserByEmailAndPassword 通过邮箱+用户吗查找用户，并且验证密码
func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, map[string]string) {
	//var user entity.User
	//dbErr := map[string]string{}
	//err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	//if gorm.IsRecordNotFoundError(err) {
	//	dbErr["no_user"] = "user not found"
	//	return nil, dbErr
	//}
	//if err != nil {
	//	dbErr["db_error"] = "database error"
	//	return nil, dbErr
	//}
	//// 校验密码
	//err = security.VerifyPassword(user.Password, u.Password)
	//if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	//	dbErr["incorrect_password"] = "incorrect password"
	//	return nil, dbErr
	//}
	//return &user, nil

	return nil, nil
}
