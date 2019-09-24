FROM golang:1.13.0-alpine3.10 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOLANG_CI_LINT_VERSION=1.18.0

RUN apk add --no-cache --update tzdata git bash && \
    rm -rf /var/cache/apk/*

# Install the linter
RUN wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v$GOLANG_CI_LINT_VERSION
RUN golangci-lint --version

# Install Cobertura coverage report converter
RUN go get github.com/t-yuki/gocover-cobertura

RUN mkdir -p /build_artifacts

WORKDIR /go/src/bitbucket.verifone.com/validation-service

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
    go build -o /go/bin/validation-service -ldflags "-X main.version=${version} -s -w -extldflags -static"

FROM scratch

COPY --from=build /go/bin/validation-service /go/bin/validation-service

EXPOSE 8080

CMD ["/go/bin/validation-service", "server"]
