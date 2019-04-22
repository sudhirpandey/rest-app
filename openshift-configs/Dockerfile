FROM registry.access.redhat.com/devtools/go-toolset-rhel7 AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux /opt/rh/go-toolset-1.11/root/usr/bin/go build -a -ldflags '-extldflags "-static"' .
FROM scratch
WORKDIR /
COPY --from=builder /opt/app-root/src/rest-api /bin/rest-api
EXPOSE 8080
CMD ["./rest-app"]
