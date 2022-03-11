package acceptance

import (
	"fmt"
	"os/exec"
	"path"
	"regexp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func DescribeCliUpgrade(gitopsTestRunner GitopsTestRunner) {
	var _ = Describe("Gitops upgrade Tests", func() {

		UI_NODEPORT := "30081"
		NATS_NODEPORT := "31491"
		var capi_endpoint_url string
		var test_ui_url string
		var stdOut string
		var stdErr string

		BeforeEach(func() {

			By("Given I have a gitops binary installed on my local machine", func() {
				Expect(fileExists(gitops_bin_path)).To(BeTrue(), fmt.Sprintf("%s can not be found.", gitops_bin_path))
			})
		})

		Context("[CLI] When Wego core is installed in the cluster", func() {
			var current_context string
			var public_ip string
			kind_upgrade_cluster_name := "test-upgrade"

			templateFiles := []string{}

			JustBeforeEach(func() {
				current_context, _ = runCommandAndReturnStringOutput("kubectl config current-context")

				// Create vanilla cluster for WGE upgrade
				createCluster("kind", kind_upgrade_cluster_name, "upgrade-kind-config.yaml")

			})

			JustAfterEach(func() {

				gitopsTestRunner.DeleteApplyCapiTemplates(templateFiles)
				templateFiles = []string{}

				err := runCommandPassThrough("kubectl", "config", "use-context", current_context)
				Expect(err).ShouldNot(HaveOccurred())

				deleteClusters("kind", []string{kind_upgrade_cluster_name})

			})

			It("@upgrade @git Verify wego core can be upgraded to wego enterprise", func() {
				repoAbsolutePath := configRepoAbsolutePath(gitProviderEnv)

				By("When I create a private repository for cluster configs", func() {
					initAndCreateEmptyRepo(gitProviderEnv, true)
				})

				By("When I install gitops/wego to my active cluster", func() {
					installAndVerifyGitops(GITOPS_DEFAULT_NAMESPACE, getGitRepositoryURL(repoAbsolutePath))
				})

				By("And I install the entitlement for cluster upgrade", func() {
					Expect(gitopsTestRunner.KubectlApply([]string{}, path.Join(getCheckoutRepoPath(), "test", "utils", "scripts", "entitlement-secret.yaml")), "Failed to create/configure entitlement")
				})

				By("And I install the git repository secret for cluster service", func() {
					cmd := fmt.Sprintf(`kubectl create secret generic git-provider-credentials --namespace=%s --from-literal="GIT_PROVIDER_TOKEN=%s"`, GITOPS_DEFAULT_NAMESPACE, gitProviderEnv.Token)
					stdOut, stdErr = runCommandAndReturnStringOutput(cmd)
					Expect(stdErr).Should(BeEmpty(), "Failed to create git repository secret for cluster service")
				})

				By("And I should update/modify the default upgrade manifest ", func() {
					public_ip = clusterWorkloadNonePublicIP("KIND")
				})

				prBranch := "wego-upgrade-enterprise"
				version := "0.0.19"
				By(fmt.Sprintf("And I run gitops upgrade command from directory %s", repoAbsolutePath), func() {
					natsURL := public_ip + ":" + NATS_NODEPORT
					upgradeCommand := fmt.Sprintf(" %s upgrade --version %s --branch %s --set 'agentTemplate.natsURL=%s' --set 'nats.client.service.nodePort=%s'", gitops_bin_path, version, prBranch, natsURL, NATS_NODEPORT)
					logger.Infof("Upgrade command: '%s'", upgradeCommand)
					stdOut, stdErr = runCommandAndReturnStringOutput(fmt.Sprintf("cd %s && %s", repoAbsolutePath, upgradeCommand))
					Expect(stdErr).Should(BeEmpty())
				})

				By("Then I should see pull request created to management cluster", func() {
					re := regexp.MustCompile(`Pull Request created.*:[\s\w\d]+(?P<URL>https.*\/\d+)`)
					match := re.FindSubmatch([]byte(stdOut))
					Eventually(match[1]).ShouldNot(BeNil(), "Failed to Create pull request")
				})

				By("Then I should merge the pull request to start weave gitops enterprise upgrade", func() {
					upgradePRUrl := verifyPRCreated(gitProviderEnv, repoAbsolutePath)
					mergePullRequest(gitProviderEnv, repoAbsolutePath, upgradePRUrl)
				})

				By("And I should see cluster upgraded from 'wego core' to 'wego enterprise'", func() {
					verifyEnterpriseControllers("weave-gitops-enterprise", "mccp-", GITOPS_DEFAULT_NAMESPACE)
				})

				By("And I can also use upgraded enterprise UI/CLI after port forwarding (for loadbalancer ingress controller)", func() {
					serviceType, _ := runCommandAndReturnStringOutput(fmt.Sprintf(`kubectl get service clusters-service -n %s -o jsonpath="{.spec.type}"`, GITOPS_DEFAULT_NAMESPACE))
					if serviceType == "NodePort" {
						capi_endpoint_url = "http://" + public_ip + ":" + UI_NODEPORT
						test_ui_url = "http://" + public_ip + ":" + UI_NODEPORT
					} else {
						commandToRun := fmt.Sprintf("kubectl port-forward --namespace %s svc/clusters-service 8000:80", GITOPS_DEFAULT_NAMESPACE)

						cmd := exec.Command("sh", "-c", commandToRun)
						session, _ := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

						go func() {
							_ = session.Command.Wait()
						}()

						test_ui_url = "http://localhost:8000"
						capi_endpoint_url = "http://localhost:8000"
					}
					initializeWebdriver(test_ui_url)
				})

				By("And the Cluster service is healthy", func() {
					gitopsTestRunner.CheckClusterService(capi_endpoint_url)
				})

				By("Then I should run enterprise CLI commands", func() {
					testGetCommand := func(subCommand string) {
						logger.Infof("Running 'gitops get %s --endpoint %s'", subCommand, capi_endpoint_url)

						cmd := fmt.Sprintf(`%s get %s --endpoint %s`, gitops_bin_path, subCommand, capi_endpoint_url)
						stdOut, stdErr = runCommandAndReturnStringOutput(cmd)
						Expect(stdErr).Should(BeEmpty(), fmt.Sprintf("'%s get %s' command failed", gitops_bin_path, subCommand))
						Expect(stdOut).Should(MatchRegexp(fmt.Sprintf(`No %s[\s\w]+found`, subCommand)), fmt.Sprintf("'%s get %s' command failed", gitops_bin_path, subCommand))
					}

					testGetCommand("templates")
					testGetCommand("credentials")
					testGetCommand("clusters")
				})

				By("And I can connect cluster to itself", func() {
					leaf := LeafSpec{
						Status:          "Ready",
						IsWKP:           false,
						AlertManagerURL: "",
						KubeconfigPath:  "",
					}
					connectACluster(webDriver, gitopsTestRunner, leaf)
				})
			})
		})
	})
}
