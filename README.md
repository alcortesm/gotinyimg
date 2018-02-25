# Gotinyimg

An example go project to show how to:
- Build Go programs in a reproducible way using Docker.
- Test Go programs in a reproducible way using Docker.
- Generate small Docker images to run Go programs.

## The Example Project

The current project is simple
but has some features commonly found in more complex projects:
- A main package to generate a command line program,
- another package with some functionality used by the main package,
- and an external dependency (using `dep`).

## Instalation

```
bash$ go get -d github.com/alcortesm/gotinyimg/...
bash$ cd ${GOPATH%%:*}/src/github.com/alcortesm/gotinyimg
bash$ dep ensure -vendor-only
bash$ go install ./cmd/gotinyimg
```

## Example of use

Run the command outside Docker:
```
bash$ gotinyimg
Hello, world!
```

Generate a Docker image with the program and run it:
```
bash$ cd ${GOPATH%%:*}/src/github.com/alcortesm/gotinyimg
bash$ docker build -t gotinyimg .
[...]
Successfully tagged gotinyimg:latest
bash$ docker run --rm gotinyimg:latest
Hello, world!
```

Check the size of the image:
```
bash$ docker images gotinyimg
REPOSITORY          TAG                 IMAGE ID            CREATED              SIZE
gotinyimg           latest              a68c8e4b5e7a        About a minute ago   2.86MB
```

List the files in a container created with the image:
```
bash$ docker export `docker create gotinyimg` | tar -tvf -
-rwxr-xr-x 0/0               0 2018-02-26 00:25 .dockerenv
drwxr-xr-x 0/0               0 2018-02-26 00:25 dev/
-rwxr-xr-x 0/0               0 2018-02-26 00:25 dev/console
drwxr-xr-x 0/0               0 2018-02-26 00:25 dev/pts/
drwxr-xr-x 0/0               0 2018-02-26 00:25 dev/shm/
drwxr-xr-x 0/0               0 2018-02-26 00:25 etc/
-rwxr-xr-x 0/0               0 2018-02-26 00:25 etc/hostname
-rwxr-xr-x 0/0               0 2018-02-26 00:25 etc/hosts
lrwxrwxrwx 0/0               0 2018-02-26 00:25 etc/mtab -> /proc/mounts
-rwxr-xr-x 0/0               0 2018-02-26 00:25 etc/resolv.conf
drwxr-xr-x 0/0               0 2018-02-26 00:21 etc/ssl/
drwxr-xr-x 0/0               0 2018-02-26 00:21 etc/ssl/certs/
-rw-r--r-- 0/0          261407 2018-02-15 04:55 etc/ssl/certs/ca-certificates.crt
drwxr-xr-x 0/0               0 2018-02-26 00:25 proc/
-rwxr-xr-x 0/0         2236878 2018-02-26 00:15 runme
drwxr-xr-x 0/0               0 2018-02-26 00:25 sys/
drwxr-xr-x 0/0               0 2018-02-26 00:21 usr/
drwxr-xr-x 0/0               0 2018-02-26 00:21 usr/local/
drwxr-xr-x 0/0               0 2018-02-26 00:21 usr/local/go/
drwxr-xr-x 0/0               0 2018-02-26 00:21 usr/local/go/lib/
drwxr-xr-x 0/0               0 2018-02-26 00:21 usr/local/go/lib/time/
-rw-r--r-- 0/0          364985 2018-02-16 18:12 usr/local/go/lib/time/zoneinfo.zip
```

Run the tests of the project inside Docker:
```
bash$ docker build -t gotinyimg:test --target=src .
[...]
Successfully tagged gotinyimg:test
bash$ docker run --rm gotinyimg:test go test -race -cover ./...
ok    github.com/alcortesm/gotinyimg  1.007s  coverage: 100.0% of statements
?     github.com/alcortesm/gotinyimg/cmd/gotinyimg  [no test files]
```

## Rationale

I use a multistage Dockerfile to have small images and reproducible builds:

- An initial "builder" image is generated
  with all the tools needed to compile and test the project,
  including the Go compiler
  and the "dep" dependency manager.

- Then an "src" image is generated
  to store the source code of our program
  and its dependencies.

- Now, a "build" image is used
  to build a static binary of our program.

- And finally a "run" image is generated from "scratch",
  including only the static binary compiled above
  and some required files (see below).

## Gotchas

The "src" Docker image is based in the official "golang" image
instead of the smaller "golang:alpine" image
because we want to run our tests inside Docker with the "-race" option,
and that option requires gcc support
that is missing from the alpine version.

The "run" image also includes some files needed by the golang stdlib,
like the timezone database
and the SSL certificates,
required by common operations like connecting through HTTPS.
