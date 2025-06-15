package main

import (
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
)

var version = "v0.1.0"

func main() {
	rootCmd := &cobra.Command{
		Use:   "kontext",
		Short: "Easily switch Kubernetes contexts",
		Run:   runKontext,
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of kontext",
		Run: func(cmd *cobra.Command, args []string) {
			color.Blue("kontext version %s", version)
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file")
	viper.BindPFlag("kubeconfig", rootCmd.Flags().Lookup("kubeconfig"))

	if err := rootCmd.Execute(); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}

func runKontext(cmd *cobra.Command, args []string) {
	// Kubeconfig path resolution: flag > env > default
	kubeconfigPath := viper.GetString("kubeconfig")
	if kubeconfigPath == "" {
		if env := os.Getenv("KUBECONFIG"); env != "" {
			kubeconfigPath = env
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				color.Red("Unable to determine home directory: %v", err)
				os.Exit(1)
			}
			kubeconfigPath = filepath.Join(home, ".kube", "config")
		}
	}

	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		color.Red("Failed to load kubeconfig: %v", err)
		os.Exit(1)
	}

	contexts := []string{}
	for name := range config.Contexts {
		contexts = append(contexts, name)
	}

	if len(contexts) == 0 {
		color.Yellow("No Kubernetes contexts found in kubeconfig.")
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
		color.Red("Prompt failed: %v", err)
		os.Exit(1)
	}

	if selected == config.CurrentContext {
		color.Cyan("'%s' is already the current context.", selected)
		return
	}

	config.CurrentContext = selected
	err = clientcmd.WriteToFile(*config, kubeconfigPath)
	if err != nil {
		color.Red("Failed to update kubeconfig: %v", err)
		os.Exit(1)
	}

	color.Green("Switched to context: %s", selected)
}
