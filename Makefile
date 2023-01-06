HEAD=$(shell git rev-parse --short HEAD)
dev:
	hugo server -D
static:
	hugo
container:
	podman build -t docker.io/jmbitci/www-jmbit-de:latest .	
	podman build -t docker.io/jmbitci/www-jmbit-de:$(HEAD) .	
publish:
	podman push docker.io/jmbitci/www-jmbit-de:latest
	podman push docker.io/jmbitci/www-jmbit-de:$(HEAD)
nopub: static container
	podman run --rm -p8080:80 docker.io/jmbitci/www-jmbit-de

rollout:
	kubectl rollout restart deployment www-jmbit-de -n jmbit-web

clean:
	rm 

all: static container publish


