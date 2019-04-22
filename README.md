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


# Running on openshift

### Create the build config and start the build
```
oc create -f openshift-configs/multistagebuild

oc start-build product-api-build
```



### Deploy the build application 
```
#deploy the the mysql db the app will use
oc new-app mysql MYSQL_USER=user MYSQL_PASSWORD=pass MYSQL_DATABASE=testdb -l db=mysql

#deploy the app the the image stream creted
```
oc new-app --image-stream=product-api -e DB_USERNAME=user DB_PASSWORD=pass -e DB_NAME=testdb -e ENVIRONMENT=PROD -e DB_HOST=mysql
```

