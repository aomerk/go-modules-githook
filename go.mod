module github.com/aomerk/go-modules-githook

go 1.15

replace github.com/valyala/fasthttp => ./vendor/blah/fasthttp

require (
	github.com/valyala/fasthttp v1.20.0
	google.golang.org/grpc v1.35.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
