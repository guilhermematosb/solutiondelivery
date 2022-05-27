# --- Base ----
FROM golang:1.18 AS base
WORKDIR $GOPATH/src/github.com/guilhermematosb/solutiondelivery

# ---- Dependencies ----
FROM base AS dependencies
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

# ---- Build ----
FROM dependencies AS build
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]

# FROM base as dev

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# WORKDIR /opt/app/api
# CMD ["air"]