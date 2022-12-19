# syntax=docker/dockerfile:1

#Starting from golang base image
FROM golang:1.16-alpine as base
 
#Adding Maintainer info
LABEL maintainer = "Garima Rawat <graw3.14@gmail.com>"

#Installing git.
#Git is required for fetching the dependencies
Run apk add --no-cache git

# Setting the current work directory inside the container
WORKDIR /app/explore-devops

#Copying go mod and sum files
Copy go.mod go.sum .

#Downloading dependenices
#Dependencies will be cached if the go.mod and go.sum files are not changed
Run go mod download

#Copying the source from the current directory to the working directory inside the container
COPY . .

# Unit test
RUN CGO_ENABLED=0 go test -v

#Build the Go App
RUN go build -o ./explore-devops .

# notifying Docker that the container listens on specified network ports at runtime
EXPOSE 8080

#command to be used to execute when the image is used to start a container
CMD [ "./explore-devops" ]
