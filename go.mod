module github.com/benny1213/etLab_BE

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.3 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210110051926-789bb1bd4061 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/benny1213/etLab_BE/conf => ./project/etLab_BE/pkg/conf
	github.com/benny1213/etLab_BE/middleware => ./project/etLab_BE/middleware
	github.com/benny1213/etLab_BE/models => ./project/etLab_BE/models
	github.com/benny1213/etLab_BE/pkg/e => ./project/etLab_BE/pkg/e
	github.com/benny1213/etLab_BE/pkg/setting => ./project/etLab_BE/pkg/setting
	github.com/benny1213/etLab_BE/pkg/util => ./project/etLab_BE/pkg/util
	github.com/benny1213/etLab_BE/routers => ./project/etLab_BE/routers
	github.com/benny1213/etLab_BE/routers/api => ./project/etLab_BE/routers/api
)
