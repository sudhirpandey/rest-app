# FROM registry.access.redhat.com/devtools/go-toolset-rhel7 
FROM scratch

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Update the DB host with the IP of your mysql container
ENV ENVIRONMENT=PROD
ENV DB_HOST=172.17.0.2
ENV DB_USERNAME=testuser
ENV DB_PASSWORD=testpasswd
ENV DB_NAME=test

# Expose port 8081
EXPOSE 8081

# Run the executable
CMD ["./rest-app"]

