base_version = 1.28.0

print:
	$(MAKE) hello

hello:
	echo ${base_version}

test:
	$(DOCKER_RUN) --entrypoint=go ./main.go

release:
	version = go run main.go
	echo "$$version"
