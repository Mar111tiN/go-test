## First Steps in GO
### repository for testing golang

#### test code in golang container
+ compile and run in golang docker:
`docker run -it -v $(pwd):/go/src/go-test golang:buster`
+ in container:
  * `cd /go/src/go-test`
  * make binary with `go install github.com/Mar111tiN/go-test`
  * run binary `/go/bin/test`
+ alternatively build a binary container from Dockerfile and run it (see [here](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324)):
  * docker built -t go-test .
  * docker run go-test



* ..let´s start the ride