app: compile
	bin/app

compile:
	go build -o bin/app cmd/*.go
