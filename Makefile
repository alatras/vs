rpm:
	docker build -f Dockerfile.rpm -t validation-service-rpm .
	- @docker rm -f validation-service-rpm 2>/dev/null || exit 0
	docker run -d --name=validation-service-rpm validation-service-rpm
	docker cp validation-service-rpm:/home/builder/rpm/x86_64/validation-service-`git rev-parse --short HEAD`-1.el7.x86_64.rpm .
	docker rm -f validation-service-rpm

clean:
	- @docker rm -f validation-service-rpm 2>/dev/null || exit 0

.PHONY: rpm