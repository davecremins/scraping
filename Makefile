workspace:
	@docker run -it --rm --name Golang_Workspace -v ${PWD}:/go/src/davecremins/scraperProject/app -w /go/src/davecremins/scraperProject/app golang /bin/bash

run:
	go get
	go build
	./app