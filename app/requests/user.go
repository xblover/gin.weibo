package requests

import (
	"gin.weibo/app/models"
)

//
type UserForm struct {
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

func (u *UserForm) emailUniqueValidator() validatorFunc {
	return func() (msg string) {
		m := &models.User{}
		if _, err := m.GetByEmail(u.Email); err != nil {
			return "邮箱已经被注册过了"
		}
		return ""
	}
}

// Validate : 验证函数
func (u *UserForm) Validate() (errors []string) {
	errors = RunValidators(
		validatorMap{
			"name": {
				RequiredValidator(u.Name),
				MaxLengthValidator(u.Name, 50),
			},
			"email": {
				RequiredValidator(u.Email),
				MaxLengthValidator(u.Email, 255),
				u.emailUniqueValidator(),
			},
			"password": {
				RequiredValidator(u.Password),
				MixLengthValidator(u.Password, 6),
				EqualValidator(u.Password, u.PasswordConfirmation),
			},
		},
		validatorMsgArr{
			"name": {
				"名称不能为空",
				"名称长度不能大于 50个字符",
			},
			"email": {
				"邮箱不能为空",
				"邮箱长度不能大于 255 个字符",
				"邮箱已经被注册过了",
			},
			"password": {
				"密码不能为空",
				"密码长度不能小于 6 个字符",
				"两次输入的密码不一致",
			},
		},
	)
	return errors
}
