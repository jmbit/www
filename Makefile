dev:
	hugo server -D
static:
	hugo
container:
	podman build -t docker.io/jmbitci/www-jmbit-de .	
publish:
	podman push docker.io/jmbitci/www-jmbit-de
nopub: static container
	podman run --rm docker.io/jmbitci/www-jmbit-de

rollout:
	kubectl rollout restart deployment www-jmbit-de -n jmbit-web

all: static container publish


