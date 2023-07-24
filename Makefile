
clean:
	rm -rf ./bin

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
	cat log_messages.json | ~/go/bin/grpcurl \
		-plaintext \
		-import-path ../opentelemetry-proto/ \
		-proto ../opentelemetry-proto/opentelemetry/proto/collector/logs/v1/logs_service.proto \
		-d @ \
		localhost:4317 opentelemetry.proto.collector.logs.v1.LogsService/Export
