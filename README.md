
# Med Appointment Service
## Description
This project is a web-based application for scheduling appointments with healthcare professionals at polyclinics. The application allows patients to register for an account, provide their personal information, and schedule appointments with their preferred doctors.


## Installation
#### Prerequires
Since the project is written in go, the computer must have go installed. 
Go: [download and install.](https://go.dev/dl/)
[Postgres too.](https://www.postgresql.org/download/)
[Docker.](https://docs.docker.com/get-docker/)
[Docker compose.]()
#### Next
First, we clone the project from github:
```bash
  git clone https://github.com/husanmusa/med-appointment-service
```

The following command is written to download the modules:
```bash
  go mod tidy
  go mod vendor
```
To use microservices, [protobuf](https://github.com/gogo/protobuf) must be installed.
Then:
```bash
protoc-gen-go: go get -u github.com/golang/protobuf/protoc-gen-go
protoc-gen-go-grpc: go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
You need to download the submodule where the protos are.
```bash
  git submodule add git@github.com:husanmusa/mh_protos.git
```
The code that runs the protos is written in a Makefile. This is why you need to install a Makefile. For [MacOS.](https://formulae.brew.sh/formula/make)

Let's clone the protos.
```bash
    make copy-proto-module
    make update-proto-module
```
Added migration for easy database management. [Installation.](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

Then:
```bash
    make migrate-up
```

## Run
You can run the project in docker if you want, or locally if you want.
Dockerfile and docker-compose are already written. You just need to install some environments :).
```bash
    go run cmd/main.go
```

## Structure
The reason for the structure of the program is that, first of all, 90% of projects in Uzbekistan have a similar structure. That's why I also chose this structure.
Second, the project is understandable and scalable.
