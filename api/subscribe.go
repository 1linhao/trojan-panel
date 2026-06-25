package api

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"trojan-panel/model/constant"
	"trojan-panel/model/vo"
	"trojan-panel/service"
	"trojan-panel/util"
)

// ClashSubscribe 获取Clash订阅地址
func ClashSubscribe(c *gin.Context) {
	accountVo := service.GetCurrentAccount(c)
	password, err := service.SelectConnectPassword(&accountVo.Id, &accountVo.Username)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(fmt.Sprintf("/api/auth/subscribe/%s", base64.StdEncoding.EncodeToString([]byte(password))), c)
}

// ClashSubscribeForSb 获取sing-box订阅地址，传id时获取指定用户，不传id时获取当前用户
func ClashSubscribeForSb(c *gin.Context) {
	var (
		accountId *uint
		username  *string
	)
	accountVo := service.GetCurrentAccount(c)
	if accountVo == nil {
		return
	}
	if idStr := c.Query("id"); idStr != "" {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil || id == 0 {
			vo.Fail(constant.ValidateFailed, c)
			return
		}
		idUint := uint(id)
		if idUint != accountVo.Id && !util.IsAdmin(accountVo.Roles) {
			vo.Fail(constant.ForbiddenError, c)
			return
		}
		accountId = &idUint
		if idUint == accountVo.Id {
			username = &accountVo.Username
		}
	} else {
		accountId = &accountVo.Id
		username = &accountVo.Username
	}
	password, err := service.SelectConnectPassword(accountId, username)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(fmt.Sprintf("/api/auth/subscribe/%s?client=sing-box", base64.StdEncoding.EncodeToString([]byte(password))), c)
}

// Subscribe 订阅
func Subscribe(c *gin.Context) {
	token := c.Param("token")
	//userAgent := c.Request.Header.Get("User-Agent")
	tokenDecode, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		vo.Fail(constant.SysError, c)
		return
	}
	pass := string(tokenDecode)

	if c.Query("client") == "sing-box" {
		account, userInfo, singBoxConfigJson, err := service.SubscribeSingBox(pass)
		if err != nil {
			vo.Fail(err.Error(), c)
			return
		}
		c.Header("content-disposition", fmt.Sprintf("attachment; filename=%s-sing-box.json", *account.Username))
		c.Header("profile-update-interval", "12")
		c.Header("subscription-userinfo", userInfo)
		c.String(200, string(singBoxConfigJson))
		return
	}

	account, userInfo, clashConfigYaml, systemConfig, err := service.SubscribeClash(pass)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	result := fmt.Sprintf(`%s
%s`, string(clashConfigYaml), systemConfig.ClashRule)

	c.Header("content-disposition", fmt.Sprintf("attachment; filename=%s.yaml", *account.Username))
	c.Header("profile-update-interval", "12")
	c.Header("subscription-userinfo", userInfo)
	c.String(200, result)
	return
	//}
	//vo.Fail("This client is not supported", c)
}
