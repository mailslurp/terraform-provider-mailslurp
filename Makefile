TEST?=$$(go list ./... |grep -v 'vendor')

build:
	go build

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -count=1

test:
	go test $(TEST) || exit 1