


build:
	mkdir -p ./bin
	cd ./collector && go build 
	mv ./collector/collector ./bin/collector