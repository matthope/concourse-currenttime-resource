FROM --platform=$BUILDPLATFORM golang:1.21-alpine3.17 AS base
WORKDIR /opt/resource

FROM base AS build
WORKDIR /src

ENV GOOS=$TARGETOS GOARCH=$TARGETARCH

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -ldflags="-w -s" -o /opt/resource/resource .
RUN for n in check in out ; do ln -s resource /opt/resource/$n ; done


FROM scratch
LABEL org.opencontainers.image.source https://github.com/matthope/concourse-currenttime-resource

COPY --from=build /opt/resource /opt/resource

CMD [ "/opt/resource/check" ]
