# Используем базовый образ с установленным protoc и Go
FROM golang:latest

# Аргумент командной строки для передачи имени файла протокола
ARG ABS_FILEPATH
ARG REL_FILEPATH
ARG FILENAME

# Устанавливаем protoc
RUN apt-get update && \
    apt-get install -y protobuf-compiler && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Создаем релятивный путь для копии protofile.
RUN mkdir /go/proto

# Переносим protofile в соответствующую дирректорию.
COPY $REL_FILEPATH /go/proto

# Создаем релятивный путь для сгенерированных файлов.
RUN mkdir /go/generated
RUN mkdir /go/generated/swagger

# Генерируем файлы из протокола и перемещаем их в папку generated.
RUN protoc --go_out=/go/generated  \
    --grpc-gateway_out=/go/generated  \
    --go_opt=paths=import  \
    --go-grpc_out=/go/generated  \
    --proto_path=/go/proto  \
    --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out /go/generated/swagger  \
    $FILENAME