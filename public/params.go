package public

import (
	"github.com/gin-gonic/gin"
)

// DefaultGetValidParams 参数绑定 && 参数校验
func DefaultGetValidParams(ginContext *gin.Context, params interface{}) error {
	if err := ginContext.ShouldBind(params); err != nil {
		return err
	}
	//
	// // 获取验证器
	// valid, err := GetValidator(ginContext)
	// if err != nil {
	// 	return err
	// }
	//
	// // 获取翻译器
	// trans, err := GetTranslation(ginContext)
	// if err != nil {
	// 	return err
	// }
	//
	// err = valid.Struct(params)
	// if err != nil {
	// 	errs := err.(validator.ValidationErrors)
	// 	var sliceErrs []string
	// 	for _, e := range errs {
	// 		sliceErrs = append(sliceErrs, e.Translate(trans))
	// 	}
	// 	return errors.New(strings.Join(sliceErrs, ","))
	// }
	return nil
}

// func GetValidator(c *gin.Context) (*validator.Validate, error) {
// 	val, ok := c.Get(ValidatorKey)
// 	if !ok {
// 		return nil, errors.New("未设置验证器")
// 	}
// 	validator, ok := val.(*validator.Validate)
// 	if !ok {
// 		return nil, errors.New("获取验证器失败")
// 	}
// 	return validator, nil
// }
//
// func GetTranslation(c *gin.Context) (ut.Translator, error) {
// 	trans, ok := c.Get(TranslatorKey)
// 	if !ok {
// 		return nil, errors.New("未设置翻译器")
// 	}
// 	translator, ok := trans.(ut.Translator)
// 	if !ok {
// 		return nil, errors.New("获取翻译器失败")
// 	}
// 	return translator, nil
// }
