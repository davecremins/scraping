workspace:
	@docker run -it --rm --name Golang_Workspace -v ${PWD}:/usr/src/app -w /usr/src/app golang /bin/bash