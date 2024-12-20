HEAD=$(shell git rev-parse --short HEAD)
CTNAME:=git.jmbit.de/jmb/www-jmbit-de

all: hugo webserver

dev:
	cd hugo && hugo server -D

hugo:
	cd hugo && hugo

webserver:
	templ generate .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o www .

container:
	podman build -t $(CTNAME):latest -t $(CTNAME):$(HEAD) .	

run:
	podman run --rm -p8080:80 $(CTNAME)

clean:
	rm -rf hugo/public
	rm -f www

.PHONY: all dev hugo webserver container run clean
