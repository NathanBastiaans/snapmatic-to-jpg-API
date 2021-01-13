FROM golang:1.15 as builder

## Create the workdir
RUN mkdir /app

## Set the workdir
WORKDIR /app

## Copy the go mod and sum
COPY go.mod /app
COPY go.sum /app

## Add this go mod download command to pull in any dependencies
RUN go mod download

## Copy all files to the work dir
COPY . /app

## Our project will now successfully build with the necessary go libraries included.
RUN go build -o api .

## Our start command which kicks off our newly created binary executable
CMD ["/app/api"]