module meshop-product-service

go 1.16

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/micrease/meshop-protos v1.0.0
	github.com/micrease/micrease-core v1.0.5
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.8.0
	github.com/micro/go-plugins/wrapper/service/v2 v2.9.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
