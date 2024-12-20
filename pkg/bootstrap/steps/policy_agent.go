package steps

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	policyAgentInstallInfoMsg    = "installing Policy Agent ..."
	policyAgentInstallConfirmMsg = "Policy Agent is installed successfully"
)

const (
	agentControllerURL        = "https://raw.githubusercontent.com/weaveworks/policy-agent/dev/docs/examples/policy-agent-helmrelease.yaml"
	agentHelmReleaseFileName  = "policy-agent-helmrelease.yaml"
	agentHelmReleaseCommitMsg = "Add Policy Agent HelmRelease YAML file"
)

// NewInstallPolicyAgentStep creates the policy agent installation step
func NewInstallPolicyAgentStep(config Config) BootstrapStep {

	if slices.Contains(config.ComponentsExtra.Existing, policyAgentController) {
		config.Logger.Warningf(" not installing %s: found in the cluster", policyAgentController)
		return BootstrapStep{
			Name:  "existing policy agent installation",
			Input: []StepInput{},
			Step:  doNothingStep,
		}
	}

	return BootstrapStep{
		Name:  "install Policy Agent",
		Input: []StepInput{},
		Step:  installPolicyAgent,
	}
}

// installPolicyAgent start installing policy agent helm chart
func installPolicyAgent(input []StepInput, c *Config) ([]StepOutput, error) {
	c.Logger.Actionf(policyAgentInstallInfoMsg)
	c.Logger.Warningf(" please note that the Policy Agent requires cert-manager to be installed!")

	// download agent file
	bodyBytes, err := doGetRequest(agentControllerURL)
	if err != nil {
		return []StepOutput{}, fmt.Errorf("error getting Policy Agent HelmRelease: %v", err)
	}

	helmreleaseFile := fileContent{
		Name:      agentHelmReleaseFileName,
		Content:   string(bodyBytes),
		CommitMsg: agentHelmReleaseCommitMsg,
	}

	c.Logger.Successf(policyAgentInstallConfirmMsg)
	return []StepOutput{
		{
			Name:  agentHelmReleaseFileName,
			Type:  typeFile,
			Value: helmreleaseFile,
		},
	}, nil
}
