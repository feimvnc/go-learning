FROM wamuir/golang-tf
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /tfapp
CMD ["/tfapp"]