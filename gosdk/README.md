Couchbase Go SDK: https://github.com/couchbase/gocb

Quickstart: https://docs.couchbase.com/go-sdk/current/hello-world/start-using-sdk.html

Connect to Capella

M2 Pro - ARM

Connect to Server

docker pull couchbase:enterprise-8.0.0

linux/arm64/v8 

https://hub.docker.com/layers/library/couchbase/enterprise-8.0.0/images/sha256-f3ada4cdd0b2c43d1687f9c992f271e450e30dcc355a9eab483f37fb24fd2ce0 

sha256:5161e61c4758c73dbe106870f7a3409848138890a738ff644c47677d8d85aa31

sha256:f3ada4cdd0b2c43d1687f9c992f271e450e30dcc355a9eab483f37fb24fd2ce0

Specify platform: 

docker pull --platform=linux/arm64/v8 couchbase:enterprise-8.0.0

docker pull --platform=linux/arm64/v8 couchbase@sha256:f3ada4cdd0b2c43d1687f9c992f271e450e30dcc355a9eab483f37fb24fd2ce0
