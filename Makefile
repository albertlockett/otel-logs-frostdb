
clean:
	rm -rf ./bin/collector
	rm ./query/*.pb.go

.PHONY: install-protoc
install-protoc:
	./dev/scripts/install-protoc.sh
	GOBIN=$(shell pwd)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(shell pwd)/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: install-tools
install-tools: install-protoc

.PHONY: gen
gen:
	PATH=$(shell pwd)/bin:$PATH protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
    ./proto/query.proto
	mv ./proto/query*.pb.go ./query


.PHONY: build
build:
	mkdir -p ./bin
	cd ./collector && go build 
	mv ./collector/collector ./bin/collector

.PHONY: start
start:
	./bin/collector --config ./collector-config.yaml

# TODO -- need to download otel-proto or something here
# e.g. clean this up cause it's effectivly hard-coded to how you downloaded things
.PHONY: send-test-log
send-test-log:
	cat log_messages.json | \
	ENV_TIMESTAMP=$(shell date +%s%N)  envsubst | \
	~/go/bin/grpcurl \
		-plaintext \
		-import-path ../opentelemetry-proto/ \
		-proto ../opentelemetry-proto/opentelemetry/proto/collector/logs/v1/logs_service.proto \
		-d @ \
		localhost:4317 opentelemetry.proto.collector.logs.v1.LogsService/Export
