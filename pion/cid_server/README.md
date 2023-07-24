# Connection ID Server

Build image:

```
docker build . -t hasheddan/dtls-interop-pion-cid-server
```

Run image:

```
docker run --rm --net host hasheddan/dtls-interop-pion-cid-server
```

Override settings using environment variables:

> Default: `5684`
```
-e DTLS_INTEROP_SERVER_PORT=4444
```

> Default: `secretPSK`
```
-e DTLS_INTEROP_CLIENT_PSK=supersecret
```

> Default: `8`
```
-e DTLS_INTEROP_SERVER_CID_LENGTH=4
```
