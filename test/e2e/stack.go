// SiGG-GoLang-On-the-Fly //
package e2e

import (
	"encoding/json"
	"os"
)

type Stack struct {
	Name                  string    `json:"name,omitempty"`
	ExposedBlockchainPort int       `json:"exposedBlockchainPort,omitempty"`
	BlockchainProvider    string    `json:"blockchainProvider"`
	TokenProviders        []string  `json:"tokenProviders"`
	Members               []*Member `json:"members,omitempty"`
	Database              string    `json:"database"`
}

type StackState struct {
	Accounts []interface{} `json:"accounts"`
}

type Member struct {
	ExposedFireflyPort   int         `json:"exposedFireflyPort,omitempty"`
	ExposedAdminPort     int         `json:"exposedFireflyAdminPort,omitempty"`
	FireflyHostname      string      `json:"fireflyHostname,omitempty"`
	Username             string      `json:"username,omitempty"`
	Password             string      `json:"password,omitempty"`
	UseHTTPS             bool        `json:"useHttps,omitempty"`
	ExposedConnectorPort int         `json:"exposedConnectorPort,omitempty"`
	OrgName              string      `json:"orgName,omitempty"`
	Account              interface{} `json:"account,omitempty"`
}

func GetMemberPort(filename string, n int) (int, error) {
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	var stack Stack
	err = json.Unmarshal(jsonBytes, &stack)
	if err != nil {
		return 0, err
	}

	return stack.Members[n].ExposedFireflyPort, nil
}

func GetMemberHostname(filename string, n int) (string, error) {
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var stack Stack
	err = json.Unmarshal(jsonBytes, &stack)
	if err != nil {
		return "", err
	}

	return stack.Members[n].FireflyHostname, nil
}

func ReadStackFile(filename string) (*Stack, error) {
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var stack *Stack
	err = json.Unmarshal(jsonBytes, &stack)
	if err != nil {
		return nil, err
	}

	// Apply defaults, in case this stack.json is a local CLI environment
	for _, member := range stack.Members {
		if member.FireflyHostname == "" {
			member.FireflyHostname = "127.0.0.1"
		}

	}

	return stack, nil
}

func ReadStackStateFile(filename string) (*StackState, error) {
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var stackState *StackState
	err = json.Unmarshal(jsonBytes, &stackState)
	if err != nil {
		return nil, err
	}

	return stackState, nil
}
