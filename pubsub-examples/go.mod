module pbsample

go 1.12

require (
	cloud.google.com/go v0.37.0
	gocloud.dev v0.12.0
	google.golang.org/genproto v0.0.0-20190306203927-b5d61aea6440
	google.golang.org/grpc v1.19.0
)

replace gocloud.dev => ../../go-cloud
