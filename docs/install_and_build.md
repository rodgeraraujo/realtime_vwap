# Install and Build

This manual gives an overview about how to install, run and build the project.


## Installation
To isntall all project dependencies tun the command bellow:
```sh
make install
```

## Running
For running the project in development mod, run the command bellow in the root of the project:

```sh
go run cmd/vwap/main.go
```
This will be start the app.

## Build

To build the project into a binary file, run the command bellow:
```sh
make build
```
This will generate a file called `vwap`, that representes the binary file for the project. To run this binary file, run the command:
```sh
./vwap
````
