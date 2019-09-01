########################################################################################################################
#					swagger partner api client generator
########################################################################################################################
generate-client:
	DEBUG=1 swagger generate client -f partners/partner_api_1_0_0.yaml  -A partner-api
build:
	goreleaser --snapshot --skip-publish --rm-dist
init-cobra:
	 cobra init --pkg-name github.com/NickTaporuk/channels_booking_clients