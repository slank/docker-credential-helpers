package main

import (
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/slank/docker-credential-helpers/awsecr"
)

func main() {
	credentials.Serve(awsecr.AWSECR{})
}
