dev:
	hugo server -D
static:
	hugo
container:
	podman build -t docker.io/jmbitci/www-jmbfountain-de .	
publish:
	podman push docker.io/jmbitci/www-jmbfountain-de
nopub: static container
	podman run --rm docker.io/jmbitci/www-jmbfountain-de

all: static container publish


