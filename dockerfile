FROM golang:1.19

WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/

# FROM scratch
# COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
