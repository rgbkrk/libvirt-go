build:
	docker build -t libvirt-go:master .

unit-test: build
	docker run -ti --privileged --rm libvirt-go:master go test

integration-test: build
	docker run -ti --privileged --rm libvirt-go:master go test -tags integration

test: unit-test integration-test
