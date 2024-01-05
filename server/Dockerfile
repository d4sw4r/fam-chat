# syntax=docker/dockerfile:1

# USAGE: docker build -t godocker:multistage -f Dockerfile.multistage .

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.21.5-alpine3.18 AS build

# create a working directory inside the image
WORKDIR /app

#enable v1 dependencies
RUN export GO111MODULE=on

# copy directory files i.e all files ending with .go
COPY . ./

# download Go modules and dependencies
RUN go mod download

# compile application
RUN go build -o /godocker

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /godocker /godocker

EXPOSE 110

ENTRYPOINT ["/godocker"]