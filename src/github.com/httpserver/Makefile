export tag=1.0
root:
	export ROOT=github.com\k8s-study
build:
	echo "building httpserver binary"
	mkdir -p "bin/amd64"
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .
release: build
	echo "building httpserver container"
	docker build -t pingzhu/httpserver:${tag} .
push: release
	echo "pushing pingzhu/httpserver:`${tag}`"
	docker push pingzhu/httpserver:${tag}