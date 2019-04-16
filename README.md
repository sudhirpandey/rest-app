# rest-app

## Compile static binary
```
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .
```

### Build docker container
```
sudo docker build -t rest-app -f Dockerfile .
```

### Run docker container
```
sudo docker run -p 8081:8081 -it rest-app
```
