module github.com/tel-io/tel/example/demo/client/v2

go 1.17

require (
	github.com/joho/godotenv v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/tel-io/instrumentation/middleware/grpc v1.1.3-0.20221112184934-38eaca0ccf95
	github.com/tel-io/instrumentation/middleware/http v1.2.2-0.20221112184934-38eaca0ccf95
	github.com/tel-io/instrumentation/module/otelgrpc v1.0.2
	github.com/tel-io/instrumentation/module/otelgrpc/example v0.0.0-20221112184934-38eaca0ccf95
	github.com/tel-io/tel/v2 v2.1.2-0.20221111152654-9e5e734f01e2
	github.com/urfave/cli/v2 v2.10.3
	go.opentelemetry.io/otel v1.11.2-0.20221111171059-308d0362e6c5
	go.opentelemetry.io/otel/metric v0.33.1-0.20221111171059-308d0362e6c5
	go.opentelemetry.io/otel/trace v1.11.1
	go.uber.org/zap v1.21.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.50.1
)

require (
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/caarlos0/env/v9 v9.0.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shirou/gopsutil/v3 v3.22.9 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.36.4 // indirect
	go.opentelemetry.io/contrib/instrumentation/host v0.36.4 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.36.4 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.36.4 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.11.2-0.20221111171059-308d0362e6c5 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.33.1-0.20221111171059-308d0362e6c5 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.33.1-0.20221111171059-308d0362e6c5 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.11.2-0.20221111171059-308d0362e6c5 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.11.2-0.20221111171059-308d0362e6c5 // indirect
	go.opentelemetry.io/otel/sdk v1.11.1 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.33.1-0.20221111171059-308d0362e6c5 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220602131408-e326c6e8e9c8 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/tel-io/tel/v2 => ../../..
