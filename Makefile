TAG = `git rev-parse --short HEAD`

rpm:
	docker build -f Dockerfile.rpm -t validation-service-rpm:${TAG} .
	- @docker rm -f validation-service-rpm-${TAG} 2>/dev/null || exit 0
	docker create --name=validation-service-rpm-${TAG} validation-service-rpm:${TAG}
	docker cp validation-service-rpm-${TAG}:/home/builder/rpm/x86_64/validation-service-${TAG}-1.el7.x86_64.rpm .
	docker rm -f validation-service-rpm-${TAG}
	docker rmi validation-service-rpm:${TAG}

clean:
	- @docker rm -f validation-service-rpm-${TAG} 2>/dev/null || exit 0
	- @docker rmi validation-service-rpm:${TAG}

.PHONY: rpm