// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"id-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/idgo",
		beego.NSNamespace("/v1",
			beego.NSRouter("/getuid", &controllers.UidController{}, "get:GetUid"),
			beego.NSRouter("/album_id", &controllers.AlbumController{}, "get:GetAlbumId"),
			beego.NSRouter("/block_id", &controllers.BlockController{}, "get:GetBlockId"),
			beego.NSRouter("/like_id", &controllers.LikeController{}, "get:GetLikeId"),
			beego.NSRouter("/visitor_id", &controllers.VisitorController{}, "get:GetVisitorId"),
			beego.NSRouter("/user_gift_id", &controllers.UserGiftController{}, "get:GetUserGiftId"),
			beego.NSRouter("/send_gift_id", &controllers.SendGiftController{}, "get:GetSendGiftId"),
			beego.NSRouter("/account_opt_id", &controllers.AccountOptController{}, "get:GetAccountOptId"),
			beego.NSRouter("/order_id", &controllers.OrderController{}, "get:GetOrderId"),
			beego.NSRouter("/point_id", &controllers.PointController{}, "get:GetPointId"),
			beego.NSRouter("/report_id", &controllers.ReportController{}, "get:GetReportId"),
		),
	)
	beego.AddNamespace(ns)
}
