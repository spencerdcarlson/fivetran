# syntax=docker/dockerfile:1
FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

# NOTE: is there a better way then manually copying every module directory?
COPY *.go ./
COPY api api
COPY args args
COPY entrypoint.sh .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./fivetran && \
    chmod +x ./fivetran && \
    ln -s "$(pwd)/fivetran" /usr/local/bin/

# TODO: use Docker layers to keep final image source code free and minimal

ENTRYPOINT ["./entrypoint.sh"]
CMD ["fivetran"]