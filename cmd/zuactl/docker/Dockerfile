# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/zuanet/zuad

WORKDIR /go/src/github.com/zuanet/zuad

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/zuanet/zuad/cmd/zuactl

RUN GOOS=linux go build -a -installsuffix cgo -o zuactl .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/zuanet/zuad/cmd/zuactl/zuactl /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
