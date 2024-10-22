# 🐋 ReAgent in a Docker Container

## Description

This utility is designed to generate `protobuf` and `swagger` files that describe
grpc requests of arbitrary services based on a `*.proto` file within a Docker container.
The utility allows you to get rid of the need to install the necessary packages,
configure paths, and perform the necessary manipulations on your machine to
generate `protobuf` and `swagger`  files. Simply describe what you want to get,
generate all files with one command, and enjoy life! 🍹

## Usage Instructions

Place a folder with a unique name `<SERVICE_NAME>` in the `/protofiles` directory,
describing the grpc requests of some service. Note that the file name
`<SERVICE_NAME>.proto` must be identical to the directory name!
This allows the generator to focus on a specific `proto` file.

Also worth noting is that the output directory is determined by defining the
`<IMPORT_PATH>` parameter.

```bash
option go_package = "<IMPORT_PATH>";
```

Note that `<IMPORT_PATH>` must start with a `/` character (e.g. `/api`).
Please do not forget to include this directive in your `<SERVICE_NAME>.proto`
file.

For example, the `/protofiles` directory contains a subdirectory `service_a`
with a `service_a.proto` file that describes the `/health` grpc request.
In this case, the `<SERVICE_NAME>` variable is equal to `service_a`.

To run the generation of all necessary protobuf and swagger files,
you need to execute a make-command of the form

```bash
make generate <SERVICE_NAME>
```

For example, if we want to generate files for the `service_a`
service, run the command

```bash
make generate service_a
```

## Voilà! 🎉

Now, under the path `/protofiles/<SERVICE_NAME>`, two subdirectories
will be generated:

`<IMPORT_PATH>` - contains protobuf files (`<SERVICE_NAME>.pb.go`,
`<SERVICE_NAME>.pb.gw.go`, `<SERVICE_NAME>_grpc.pb.go`).

`swagger` - contains a swagger file (`<SERVICE_NAME>_grpc.swagger.json`).

Note that the utility automatically deletes the created Docker container
and image after itself, so you don't need to manually clean up the system.

## Requirements

For correct operation, it is required to install and run Docker.
