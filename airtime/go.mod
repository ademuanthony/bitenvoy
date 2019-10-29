module github.com/ademuanthony/bitenvoy/airtime

go 1.13

require (
	github.com/ademuanthony/bitenvoy/utils v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jessevdk/go-flags v1.4.0
	github.com/lib/pq v1.2.0 // indirect
	github.com/micro/go-micro v1.14.0
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/sqlboiler v3.6.0+incompatible
)

replace github.com/ademuanthony/bitenvoy/utils => ../utils
