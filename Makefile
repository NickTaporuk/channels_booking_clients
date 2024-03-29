########################################################################################################################
#					swagger partner api client generator
########################################################################################################################
generate-client:
	DEBUG=1 swagger generate client -f partners/partner_api_1_0_0.yaml  -A partner-api
build:
	goreleaser --snapshot --skip-publish --rm-dist
lint:
	GO111MODULE=on go get
	golangci-lint run -v