HEAD=$(shell git rev-parse --short HEAD)
dev:
	hugo server -D
hugo:
	hugo
container:
	podman build -t docker.io/jmbitci/www-jmbit-de:latest -t docker.io/jmbitci/www-jmbit-de:$(HEAD) .	
publish:
	podman push docker.io/jmbitci/www-jmbit-de:latest --all-tags

nopub: hugo container
	podman run --rm -p8080:80 docker.io/jmbitci/www-jmbit-de

#rollout:
#	kubectl --context=jmbit-prod rollout restart deployment www-jmbit-de -n jmbit-web

clean:
	rm -rf public 

all: hugo container publish


