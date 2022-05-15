FROM golang:1.16 AS builder

#For multistage build need to disable CGO_ENABLED
ENV CGO_ENABLED=0
ENV GOOS=linux


# putting all of the project code in app
RUN mkdir /app
ADD . /app
WORKDIR /app

#Copying all the files
COPY go.mod go.sum ./

# Build the app
RUN go build -a .

# Light weight alpine version
# The size of the image will be relatively smaller.
FROM alpine:latest AS app


RUN apk --no-cache add ca-certificates

# copy the output from to /dns path
COPY --from=builder /app/ /dns

#setting the path
ENV PATH /dns:$PATH

EXPOSE 8080
# Executing the command
CMD ["dns", "serve"]