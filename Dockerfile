#build stage
# FROM golang:alpine AS builder
# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .
# RUN go get -d -v ./...
# RUN go build -o /go/bin/app -v ./...

#final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=urfuprojectsmartsheduleformer Version=0.0.1
# EXPOSE 5432

FROM golang:1.21.3

WORKDIR /urfu-project-smart-schedule-former

COPY go.sum .
COPY go.mod .
RUN go mod download

COPY . .

# ENV PATH=$PATH:/custom/dir/bin

RUN CGO_ENABLED=0 GOOS=linux go build -o ./urfu-project-smart-schedule-former

# RUN chmod +x app

# RUN chmod +x /app
# ENTRYPOINT [ "./app" ]
CMD [ "./urfu-project-smart-schedule-former" ]

EXPOSE 8000


# COPY go.mod go.sum ./
# RUN go mod download
# COPY *.go ./


# RUN go build -o bin .

# ENTRYPOINT ["/app/bin"]
