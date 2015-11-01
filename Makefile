.PHONY: build_docker cmd shell test run

default: shell

build_docker:
	@./environment/dev/scripts/build_docker.sh

up:
	@./environment/dev/scripts/up.sh

run:
	@./environment/dev/scripts/exec.sh go run main.go

stop:
	@./environment/dev/scripts/stop.sh

rm:
	@./environment/dev/scripts/rm.sh

test:
	@./environment/dev/scripts/exec.sh go test ./...
    #MONF_SETTINGS=`pwd`/environment/dev/settings-test.yaml go test ./... -v

cmd:
	@./environment/dev/scripts/exec.sh ${ARGS}

shell:
	@./environment/dev/scripts/exec.sh /bin/bash
