package subhelper

import (
	"fmt"
	"regexp"

	"github.com/docker/docker-credential-helpers/credentials"
	credclient "github.com/docker/docker/cliconfig/credentials"
)

// HelperPattern matches a registry server URL with its credential helper
type HelperPattern struct {
	Regexp regexp.Regexp
	Helper string
}

func callHelperAdd(helper string, creds *credentials.Credentials) error {
	return nil
}

func callHelperDelete(helper, serverURL string) error {
	return nil
}

func callHelperGet(helper, serverURL string) (string, string, error) {
	return "", "", nil
}

// Subhelper delegates secret handling to other helpers.
type Subhelper struct {
	helperPatterns []HelperPattern
}

// NewSubhelper creates a new Subhelper with the specified helper config
func NewSubhelper(helpers []HelperPattern) Subhelper {
	return Subhelper{
		helperPatterns: helpers,
	}
}

// findHelper identifies the helper for the given registry server URL
func (s Subhelper) findHelper(serverURL string) (string, error) {
	for _, patt := range s.helperPatterns {
		if patt.Regexp.MatchString(serverURL) {
			return patt.Helper, nil
		}
	}
	return "", fmt.Errorf("No helper configured for %s", serverURL)
}

// Add passes credentials to the configured subhelper for storage
func (s Subhelper) Add(creds *credentials.Credentials) error {
	helper, err := s.findHelper(creds.ServerURL)
	if err != nil {
		return err
	}
	return callHelperAdd(helper, creds)
}

// Delete passes the given registry server URL to the configured helper for deletion.
func (s Subhelper) Delete(serverURL string) error {
	return nil
}

// Get passes the given registry server URL to the configured helper and returns the username and secret.
func (s Subhelper) Get(serverURL string) (string, string, error) {
	return "", "", nil
}
