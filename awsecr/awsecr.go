package awsecr

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/docker/docker-credential-helpers/credentials"
)

var awsECRServerPattern = regexp.MustCompile(
	`(?P<regId>[0-9]*)\.dkr\.ecr\.(ap|eu|us|us-gov|sa|cn)-(north|south)?(central|east|west)?-[0-9]\.amazonaws\.com`,
)

// Subhelper delegates secret handling to other helpers.
type AWSECR struct{}

// Add passes credentials to the configured subhelper for storage
func (s AWSECR) Add(creds *credentials.Credentials) error {
	logf("Called Add %+v\n", creds)
	if !awsECRServerPattern.MatchString(creds.ServerURL) {
		return fmt.Errorf("Not an ECR registry")
	}
	return fmt.Errorf("Not implemented")
}

// Delete passes the given registry server URL to the configured helper for deletion.
func (s AWSECR) Delete(serverURL string) error {
	logf("Called Delete %s\n", serverURL)
	if !awsECRServerPattern.MatchString(serverURL) {
		return fmt.Errorf("Not an ECR registry")
	}
	return fmt.Errorf("Not implemented")
}

// Get passes the given registry server URL to the configured helper and returns the username and secret.
func (s AWSECR) Get(serverURL string) (string, string, error) {
	logf("Called Get %s\n", serverURL)
	urlInfo := awsECRServerPattern.FindStringSubmatch(serverURL)
	if urlInfo == nil {
		return "", "", fmt.Errorf("Not an ECR registry")
	}
	regId := urlInfo[1]

	sess := session.New()
	logf("Region: %s\n", *sess.Config.Region)
	svc := ecr.New(sess)
	input := ecr.GetAuthorizationTokenInput{
		RegistryIds: []*string{aws.String(regId)},
	}
	logf("Input: %s\n", input)
	resp, err := svc.GetAuthorizationToken(&input)
	if err != nil {
		return "", "", err
	}
	logf("AuthorizationData: %+v", *resp.AuthorizationData[0])
	token := *resp.AuthorizationData[0].AuthorizationToken

	time.Sleep(3 * time.Second)
	return "AWS", token, nil
}

func logf(tmpl string, params ...interface{}) {
	f, err := os.OpenFile("awsecr.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf(tmpl, params...)); err != nil {
		panic(err)
	}
}
