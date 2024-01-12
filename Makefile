HEAD=$(shell git rev-parse --short HEAD)
CTNAME:=git.jmbit.de/jmb/www-jmbit-de

all: hugo container

dev:
	hugo server -D

hugo:
	hugo --minify

webserver:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

container:
	podman build -t $(CTNAME):latest -t $(CTNAME):$(HEAD) .	

run:
	podman run --rm -p8080:80 $(CTNAME)

clean:
	rm -rf public 

