module github.com/wcygan/yap/yap-api

go 1.22.0

require (
	github.com/nats-io/nats-server/v2 v2.10.17
	github.com/nats-io/nats.go v1.36.0
	github.com/wcygan/yap/generated/go v0.0.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.5.7 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
)

replace github.com/wcygan/yap/generated/go => ../generated/go
