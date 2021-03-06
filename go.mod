module github.com/benny1213/follow-golang-example

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	// github.com/EDDYCJY/go-gin-example v0.0.0-20201228125222-28f372bf41f9
	github.com/astaxie/beego v1.12.3
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/spec v0.20.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.3 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20210112080510-489259a85091 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.0.0-20210111221946-d33bae441459 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/benny1213/follow-golang-example/conf => ./project/follow-golang-example/pkg/conf
	github.com/benny1213/follow-golang-example/docs => ./project/follow-golang-example/docs
	github.com/benny1213/follow-golang-example/middleware => ./project/follow-golang-example/middleware
	github.com/benny1213/follow-golang-example/middleware/jwt => ./project/follow-golang-example/middlewarejwt
	github.com/benny1213/follow-golang-example/models => ./project/follow-golang-example/models
	github.com/benny1213/follow-golang-example/pkg/e => ./project/follow-golang-example/pkg/e
	github.com/benny1213/follow-golang-example/pkg/logging => ./project/follow-golang-example/pkg/logging
	github.com/benny1213/follow-golang-example/pkg/setting => ./project/follow-golang-example/pkg/setting
	github.com/benny1213/follow-golang-example/pkg/util => ./project/follow-golang-example/pkg/util
	github.com/benny1213/follow-golang-example/routers => ./project/follow-golang-example/routers
	github.com/benny1213/follow-golang-example/routers/api => ./project/follow-golang-example/routers/api
)
