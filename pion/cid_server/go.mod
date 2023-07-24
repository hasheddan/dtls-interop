module github.com/hasheddan/dtls-interop/pion/cid_server

go 1.20

require github.com/pion/dtls/v2 v2.2.7

replace github.com/pion/dtls/v2 => github.com/hasheddan/dtls/v2 v2.0.0-20230719021958-aad3d97ee307

require (
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/transport/v2 v2.2.2-0.20230711104634-a789100cc553 // indirect
	golang.org/x/crypto v0.10.0 // indirect
)
