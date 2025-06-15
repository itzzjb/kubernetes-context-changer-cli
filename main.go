package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load kubeconfig: %v\n", err)
		os.Exit(1)
	}

	contexts := []string{}
	for name := range config.Contexts {
		contexts = append(contexts, name)
	}

	if len(contexts) == 0 {
		fmt.Println("No Kubernetes contexts found in kubeconfig.")
		os.Exit(1)
	}

	selected := ""
	prompt := &survey.Select{
		Message: "Choose a Kubernetes context:",
		Options: contexts,
		Default: config.CurrentContext,
	}
	err = survey.AskOne(prompt, &selected)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prompt failed: %v\n", err)
		os.Exit(1)
	}

	if selected == config.CurrentContext {
		fmt.Printf("'%s' is already the current context.\n", selected)
		return
	}

	config.CurrentContext = selected
	err = clientcmd.WriteToFile(*config, kubeconfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to update kubeconfig: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Switched to context: %s\n", selected)
}
