module github.com/ruandg/microservices/shipping

go 1.26.1

replace github.com/ruandg/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping

require (
	github.com/huseyinbabal/microservices/payment v0.0.0-20230110182123-6a0c8d9f8a8a
	github.com/ruandg/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
	github.com/ruandg/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.4
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.3.2
	google.golang.org/grpc v1.82.0
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.2
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.3.2 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.43.0 // indirect
	go.opentelemetry.io/otel/metric v1.43.0 // indirect
	go.opentelemetry.io/otel/trace v1.43.0 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260414002931-afd174a4e478 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/ruandg/microservices-proto/golang/payment => ../../microservices-proto/golang/payment
