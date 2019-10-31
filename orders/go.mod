module github.com/ademuanthony/bitenvoy/orders

go 1.13

require (
	github.com/ademuanthony/bitenvoy/airtime v0.0.0-00010101000000-000000000000
	github.com/ademuanthony/bitenvoy/utils v0.0.0-00010101000000-000000000000
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jessevdk/go-flags v1.4.0
	github.com/micro/go-micro v1.14.0
	github.com/prometheus/common v0.6.0
	github.com/volatiletech/sqlboiler v3.6.0+incompatible
)

replace (
	github.com/ademuanthony/bitenvoy/airtime => ../airtime
	github.com/ademuanthony/bitenvoy/utils => ../utils
)
