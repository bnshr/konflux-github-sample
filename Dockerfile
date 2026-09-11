# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/server .

FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

LABEL name="konflux-github-sample"
LABEL summary="Minimal HTTP service for learning Konflux GitHub onboarding"
LABEL io.k8s.display-name="konflux-github-sample"
LABEL io.k8s.description="Sample Docker app for Konflux component onboarding"
LABEL maintainer="bmandal@redhat.com"

COPY --from=builder /out/server /server

EXPOSE 8080
USER 65532:65532

ENTRYPOINT ["/server"]
