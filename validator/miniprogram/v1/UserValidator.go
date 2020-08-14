// @Description	小程序接口验证器
// @Author    2020/8/14 10:53
package validatorMiniprogramV1

type UserValidator struct {
	Code string `validate:"required"`
	Tode string `validate:"required"`
}
