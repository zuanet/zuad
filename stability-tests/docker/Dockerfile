ARG ZUAD_IMAGE
ARG ZUAMINER_IMAGE

FROM ${ZUAD_IMAGE} as zuad
FROM ${ZUAMINER_IMAGE} as zuaminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/zuanet/zuad

WORKDIR /go/src/github.com/zuanet/zuad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=zuad /app/ /app/
COPY --from=zuaminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/zuanet/zuad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]