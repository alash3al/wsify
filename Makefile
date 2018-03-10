test:
	go test -v ./...

install-gometalinter:
	go get -v -u github.com/alecthomas/gometalinter
	gometalinter --install

LINT=$(eval export GOGC=400)\
gometalinter --enable-all -D dupl -D lll -D gas -D goconst -D interfacer -D safesql -D test -D testify -D vetshadow\
 --tests --vendor --warn-unmatched-nolint --deadline=10m --concurrency=2 --enable-gc ./...
lint: install-gometalinter
	$(LINT)

install-dep:
	go get -v -u github.com/golang/dep/cmd/dep

dep-init: install-dep
	dep init

dep-ensure: install-dep
	dep ensure

dep-update: install-dep
	dep ensure -update

.PHONY:: \
	test \
	lint \
	dep-ensure \
	dep-update \
