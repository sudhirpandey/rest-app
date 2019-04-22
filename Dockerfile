FROM registry.access.redhat.com/devtools/go-toolset-rhel7
RUN /opt/rh/go-toolset-1.11/root/usr/bin/go get github.com/rs/cors && \
    /opt/rh/go-toolset-1.11/root/usr/bin/go get github.com/gorilla/mux && \
    /opt/rh/go-toolset-1.11/root/usr/bin/go get github.com/jinzhu/gorm  && \
    /opt/rh/go-toolset-1.11/root/usr/bin/go get github.com/prometheus/client_golang/prometheus && \
    /opt/rh/go-toolset-1.11/root/usr/bin/go get github.com/prometheus/client_golang/prometheus/promhttp && \
    /opt/rh/go-toolset-1.11/root/usr/bin/go get github.com/jinzhu/gorm && \
    /opt/rh/go-toolset-1.11/root/usr/bin/go get github.com/go-sql-driver/mysql
RUN pwd 
COPY . .
RUN cd /opt/app-root/src 
RUN ls /opt/app-root/src
RUN CGO_ENABLED=0 GOOS=linux /opt/rh/go-toolset-1.11/root/usr/bin/go build -a -ldflags '-extldflags "-static"' .
