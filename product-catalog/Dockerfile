FROM golang:1.17.3-buster AS build
WORKDIR /src/
COPY . /src/

RUN go get -u github.com/gorilla/mux
RUN go get -u gopkg.in/yaml.v2
RUN go get -u github.com/google/uuid
RUN go get -u github.com/jmoiron/sqlx
RUN go get -u github.com/joho/godotenv
RUN go get -u github.com/lib/pq
RUN go get -u google.golang.org/grpc
RUN go get -u google.golang.org/protobuf


RUN CGO_ENABLED=0 go build -o /bin/main

FROM scratch
COPY --from=build /bin/main /bin/main
EXPOSE 5002 50002
ENTRYPOINT ["/bin/main"]