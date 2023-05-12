package controller

import (
	"crypto/rand"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/email"

	"github.com/gin-gonic/gin"
)

type ValidationCodeController struct {
}

func NewValidationCodeController() Controller {
	return &ValidationCodeController{}
}

func (c *ValidationCodeController) RegisterRoutes(gr *gin.RouterGroup) {
	group := gr.Group("/v1/validation_codes")
	group.POST("", c.Create)
}

func (c *ValidationCodeController) GetPaged(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *ValidationCodeController) Get(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (vcc *ValidationCodeController) Create(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(400, "参数错误")
		return
	}
	str, err := generateDigits()
	if err != nil {
		log.Println("[generateDigits fail]", err)
		c.String(500, "生成验证码失败")
		return
	}
	q := database.NewQuery()
	vc, err := q.CreateValidationCode(c, queries.CreateValidationCodeParams{
		Email: body.Email,
		Code:  str,
	})
	if err != nil {
		// TODO 没有做校验
		c.Status(400)
		return
	}

	if err := email.SendValidationCode(vc.Email, vc.Code); err != nil {
		log.Println("[SendValidationCode fail]", err)
		c.String(500, "发送失败")
		return
	}
	c.Status(200)
}

func (c *ValidationCodeController) Update(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *ValidationCodeController) Destroy(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

// helpers

func generateDigits() (string, error) {
	len := 4
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	digits := make([]byte, len)
	for i := range b {
		digits[i] = b[i]%10 + 48
	}
	return string(digits), nil
}
