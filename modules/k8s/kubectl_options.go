package k8s

import (
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/testing"
)

// KubectlOptions represents common options necessary to specify for all Kubectl calls
type KubectlOptions struct {
	ContextName string
	ConfigPath  string
	Namespace   string
	Env         map[string]string
	Logger      *logger.Logger // Set a non-default logger that should be used. See the logger package for more info.
}

// NewKubectlOptions will return a pointer to new instance of KubectlOptions with the configured options
func NewKubectlOptions(contextName string, configPath string, namespace string) *KubectlOptions {
	return &KubectlOptions{
		ContextName: contextName,
		ConfigPath:  configPath,
		Namespace:   namespace,
		Env:         map[string]string{},
	}
}

// SetKubectlOptionsLogger will return a pointer to new instance of KubectlOptions with the configured Logger
func SetKubectlOptionsLogger(options *KubectlOptions, log *logger.Logger) {
	options.Logger = log
}

// GetConfigPath will return a sensible default if the config path is not set on the options.
func (kubectlOptions *KubectlOptions) GetConfigPath(t testing.TestingT) (string, error) {
	// We predeclare `err` here so that we can update `kubeConfigPath` in the if block below. Otherwise, go complains
	// saying `err` is undefined.
	var err error

	kubeConfigPath := kubectlOptions.ConfigPath
	if kubeConfigPath == "" {
		kubeConfigPath, err = GetKubeConfigPathE(t)
		if err != nil {
			return "", err
		}
	}
	return kubeConfigPath, nil
}
