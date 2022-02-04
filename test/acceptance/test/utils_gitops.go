package acceptance

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func waitForResource(resourceType string, resourceName string, namespace string, kubeconfig string, timeout time.Duration) error {
	pollInterval := 5
	if timeout < 5*time.Second {
		timeout = 5 * time.Second
	}

	if kubeconfig != "" {
		kubeconfig = "--kubeconfig=" + kubeconfig
	}

	timeoutInSeconds := int(timeout.Seconds())
	for i := pollInterval; i < timeoutInSeconds; i += pollInterval {
		logger.Infof("Waiting for %s in namespace: %s... : %d second(s) passed of %d seconds timeout", resourceType+"/"+resourceName, namespace, i, timeoutInSeconds)
		err := runCommandPassThroughWithoutOutput("sh", "-c", fmt.Sprintf("kubectl %s get %s %s -n %s", kubeconfig, resourceType, resourceName, namespace))
		if err == nil {
			stdOut, _ := runCommandAndReturnStringOutput(fmt.Sprintf("kubectl %s get %s %s -n %s", kubeconfig, resourceType, resourceName, namespace))

			noResourcesFoundMessage := ""
			if namespace == "default" {
				noResourcesFoundMessage = "No resources found"
			} else {
				noResourcesFoundMessage = fmt.Sprintf("No resources found in %s namespace", namespace)
			}
			if len(stdOut) == 0 || strings.Contains(stdOut, noResourcesFoundMessage) {
				logger.Infof("Got message => {" + noResourcesFoundMessage + "} Continue looking for resource(s)")
			} else {
				return nil
			}
		}
		time.Sleep(time.Duration(pollInterval) * time.Second)
	}
	return fmt.Errorf("error: Failed to find the resource %s of type %s, timeout reached", resourceName, resourceType)
}

func verifyCoreControllers(namespace string) {
	Expect(waitForResource("deploy", "helm-controller", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", "kustomize-controller", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", "notification-controller", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", "source-controller", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", "image-automation-controller", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", "image-reflector-controller", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("pods", "", namespace, "", ASSERTION_2MINUTE_TIME_OUT))

	By("And I wait for the gitops core controllers to be ready", func() {
		_, _ = runCommandAndReturnStringOutput(fmt.Sprintf("kubectl wait --for=condition=Ready --timeout=180s -n %s --all pod --selector='app!=wego-app'", namespace), ASSERTION_3MINUTE_TIME_OUT)
	})
}

func verifyEnterpriseControllers(releaseName string, mccpPrefix, namespace string) {
	// SOMETIMES (?) (with helm install ./local-path), the mccpPrefix is skipped
	Expect(waitForResource("deploy", releaseName+"-"+mccpPrefix+"gitops-repo-broker", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", releaseName+"-"+mccpPrefix+"event-writer", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", releaseName+"-"+mccpPrefix+"cluster-service", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", releaseName+"-nginx-ingress-controller", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	// FIXME
	// const maxDeploymentLength = 63
	// Expect(waitForResource("deploy", (releaseName + "-nginx-ingress-controller-default-backend")[:maxDeploymentLength], namespace, ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("deploy", releaseName+"-wkp-ui-server", namespace, "", ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("pods", "", namespace, "", ASSERTION_2MINUTE_TIME_OUT))

	By("And I wait for the gitops enterprise controllers to be ready", func() {
		_, _ = runCommandAndReturnStringOutput(fmt.Sprintf("kubectl wait --for=condition=Ready --timeout=180s -n %s --all pod --selector='app!=wego-app'", namespace), ASSERTION_3MINUTE_TIME_OUT)
	})
}

func controllerStatus(controllerName, namespace string) error {
	return runCommandPassThroughWithoutOutput("sh", "-c", fmt.Sprintf("kubectl rollout status deployment %s -n %s", controllerName, namespace))
}

func runWegoAddCommand(repoAbsolutePath string, addCommand string, namespace string) {
	logger.Infof("Add command to run: %s in namespace %s from dir %s", addCommand, namespace, repoAbsolutePath)
	_, errOutput := runCommandAndReturnStringOutput(fmt.Sprintf("cd %s && %s %s", repoAbsolutePath, GITOPS_BIN_PATH, addCommand))
	Expect(errOutput).Should(BeEmpty())
}

func verifyWegoAddCommand(appName string, namespace string) {
	cmd := fmt.Sprintf(" kubectl wait --for=condition=Ready --timeout=60s -n %s GitRepositories --all", namespace)
	_, _ = runCommandAndReturnStringOutput(cmd, ASSERTION_5MINUTE_TIME_OUT)
	Expect(waitForResource("GitRepositories", appName, namespace, "", ASSERTION_6MINUTE_TIME_OUT)).To(Succeed())
}

func installAndVerifyGitops(gitopsNamespace string, manifestRepoURL string) {

	// Deploy key secret should not exist already
	deleteGitopsDeploySecret(GITOPS_DEFAULT_NAMESPACE)

	cmd := fmt.Sprintf("%s install --config-repo %s --namespace=%s --auto-merge", GITOPS_BIN_PATH, manifestRepoURL, gitopsNamespace)
	By(fmt.Sprintf("And I run '%s'", cmd), func() {
		_, stdErr := runCommandAndReturnStringOutput(cmd, ASSERTION_5MINUTE_TIME_OUT)
		Expect(stdErr).Should(BeEmpty())

		verifyCoreControllers(gitopsNamespace)
	})
}

func removeGitopsCapiClusters(appName string, clusternames []string, nameSpace string) {
	susspendGitopsApplication(appName, nameSpace)

	deleteClusters("capi", clusternames)

	deleteGitopsApplication(appName, nameSpace)
}

func susspendGitopsApplication(appName string, nameSpace string) {
	cmd := fmt.Sprintf("%s suspend app %s", GITOPS_BIN_PATH, appName)
	By(fmt.Sprintf("And I run '%s'", cmd), func() {
		_, _ = runCommandAndReturnStringOutput(cmd)
	})
}

func listGitopsApplication(appName string, nameSpace string) string {
	var stdOut string
	cmd := fmt.Sprintf("%s get app %s", GITOPS_BIN_PATH, appName)
	By(fmt.Sprintf("And I run '%s'", cmd), func() {
		stdOut, _ = runCommandAndReturnStringOutput(cmd)
	})
	return stdOut
}

func deleteGitopsApplication(appName string, nameSpace string) {
	cmd := fmt.Sprintf("%s delete app %s --auto-merge=true", GITOPS_BIN_PATH, appName)
	By(fmt.Sprintf("And I run '%s'", cmd), func() {
		_, _ = runCommandAndReturnStringOutput(cmd)

		appDeleted := func() bool {
			status := listGitopsApplication(appName, nameSpace)
			return status == ""
		}
		Eventually(appDeleted, ASSERTION_2MINUTE_TIME_OUT, POLL_INTERVAL_5SECONDS).Should(BeTrue(), fmt.Sprintf("%s application failed to delete", appName))
	})
}

func deleteGitopsDeploySecret(nameSpace string) {
	cmd := fmt.Sprintf(`kubectl get secrets -n %[1]v  | grep Opaque | grep wego- | cut -d' ' -f1 | xargs kubectl delete secrets -n %[1]v`, nameSpace)
	By("And I delete deploy key secret", func() {
		_, _ = runCommandAndReturnStringOutput(cmd)
	})
}

func clusterWorkloadNonePublicIP(clusterKind string) string {
	var expernal_ip string
	if clusterKind == "EKS" || clusterKind == "GKE" {
		node_name, _ := runCommandAndReturnStringOutput(`kubectl get node --selector='!node-role.kubernetes.io/master' -o name | head -n 1`)
		worker_name := strings.Split(node_name, "/")[1]
		expernal_ip, _ = runCommandAndReturnStringOutput(fmt.Sprintf(`kubectl get nodes -o jsonpath="{.items[?(@.metadata.name=='%s')].status.addresses[?(@.type=='ExternalIP')].address}"`, worker_name))
	} else {
		switch runtime.GOOS {
		case "darwin":
			expernal_ip, _ = runCommandAndReturnStringOutput(`ifconfig en0 | grep -i MASK | awk '{print $2}' | cut -f2 -d:`)
		case "linux":
			expernal_ip, _ = runCommandAndReturnStringOutput(`ifconfig eth0 | grep -i MASK | awk '{print $2}' | cut -f2 -d:`)
		}
	}
	return expernal_ip
}

func createCluster(clusterType string, clusterName string, configFile string) {
	if clusterType == "kind" {
		err := runCommandPassThrough("kind", "create", "cluster", "--name", clusterName, "--image=kindest/node:v1.20.7", "--config", "../../utils/data/"+configFile)
		Expect(err).ShouldNot(HaveOccurred())
	} else {
		Fail(fmt.Sprintf("%s cluster type is not supported for test WGE upgrade", clusterType))
	}
}

func deleteClusters(clusterType string, clusters []string) {
	for _, cluster := range clusters {
		if clusterType == "kind" {
			logger.Infof("Deleting cluster: %s", cluster)
			err := runCommandPassThrough("kind", "delete", "cluster", "--name", cluster)
			Expect(err).ShouldNot(HaveOccurred())
		} else {
			err := runCommandPassThrough("kubectl", "get", "cluster", cluster)
			if err == nil {
				logger.Infof("Deleting cluster: %s", cluster)
				err := runCommandPassThrough("kubectl", "delete", "cluster", cluster)
				Expect(err).ShouldNot(HaveOccurred())
				err = runCommandPassThrough("kubectl", "get", "cluster", cluster)
				Expect(err).Should(HaveOccurred(), fmt.Sprintf("Failed to delete cluster %s", cluster))
			}
		}
	}
}

func verifyCapiClusterKubeconfig(kubeconfigPath string, capiCluster string) {
	contents, err := ioutil.ReadFile(kubeconfigPath)
	Expect(err).ShouldNot(HaveOccurred())
	Eventually(contents).Should(MatchRegexp(fmt.Sprintf(`context:\s+cluster: %s`, capiCluster)))

	if runtime.GOOS == "darwin" {
		// Point the kubeconfig to the exposed port of the load balancer, rather than the inaccessible container IP.
		_, stdErr := runCommandAndReturnStringOutput(fmt.Sprintf(`sed -i -e "s/server:.*/server: https:\/\/$(docker port %s-lb 6443/tcp | sed "s/0.0.0.0/127.0.0.1/")/g" %s`, capiCluster, kubeconfigPath))
		Expect(stdErr).Should(BeEmpty(), "Failed to delete ClusterBootstrapConfig secret")
	}
}

func verifyCapiClusterHealth(kubeconfigPath string, capiCluster string) {

	Expect(waitForResource("nodes", "", "default", kubeconfigPath, ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("pods", "", "kube-system", kubeconfigPath, ASSERTION_2MINUTE_TIME_OUT))
	Expect(waitForResource("pods", "", "wego-system", kubeconfigPath, ASSERTION_2MINUTE_TIME_OUT))
}
