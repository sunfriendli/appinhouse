// filters
package filters

import (
	. "appinhouse/server/constants"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var SecretKeyFilter = func(ctx *context.Context) {
	secretKey := ctx.Input.Header("X-SecretKey")
	if secretKey == "" || len(secretKey) == 0 || secretKey != Secret_Key {
		msg := "Missing or invalid X-SecretKey HTTP header"
		beego.Info(msg)
		ctx.Output.SetStatus(400)
		ctx.Output.Body([]byte(msg))
	}
}

var SecurityGroupFilter = func(ctx *context.Context) {
	ip := ctx.Input.IP()
	if !AllowIP(ip) {
		msg := "not allow this ip access."
		beego.Info(msg, "ip:", ip)
		ctx.Output.SetStatus(400)
		ctx.Output.Body([]byte(msg))
	}
}
