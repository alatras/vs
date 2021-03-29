FROM golang:1.13.14-buster AS build

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOLANG_CI_LINT_VERSION=1.18.0

RUN apt-get update && apt-get install -y --no-install-recommends g++ gcc libc6-dev wget && \
    rm -rf /var/lib/apt/lists/*

# Install the linter
RUN wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v$GOLANG_CI_LINT_VERSION
RUN golangci-lint --version

# Install Cobertura coverage report converter
RUN go get github.com/t-yuki/gocover-cobertura

RUN mkdir -p /build_artifacts

WORKDIR /go/src/validation-service

COPY . .

# Run linter
RUN golangci-lint run

# Run tests
RUN go test -v -timeout 30s -coverprofile=/build_artifacts/coverage.out ./... && \
    gocover-cobertura < /build_artifacts/coverage.out > /build_artifacts/coverage.xml && \
    rm /build_artifacts/coverage.out

# Build the app
RUN go version && \
    commit=$(git rev-parse --short HEAD) && \
    branch=$(git rev-parse --abbrev-ref HEAD) &&\
    version="$branch:$commit" && \
    echo "Building version $version..." && \
    go build -o /go/bin/validation-service -ldflags "-X main.version=${version} -s -w"

FROM golang:1.13.14-buster

WORKDIR /go/bin

ENV LD_LIBRARY_PATH=/go/lib

COPY --from=build /go/src/validation-service/appdynamics/lib /go/lib
COPY --from=build /go/bin/validation-service /go/bin/validation-service
COPY --from=build /go/src/validation-service/config.yml /go/bin/

EXPOSE 8080

CMD ["validation-service"]
