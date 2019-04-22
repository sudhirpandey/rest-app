FROM registry.access.redhat.com/devtools/go-toolset-rhel7
COPY . .
RUN CGO_ENABLED=0 GOOS=linux /opt/rh/go-toolset-1.11/root/usr/bin/go build -a -ldflags '-extldflags "-static"' .
