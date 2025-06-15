package main

import (
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// ktx - Easily switch Kubernetes contexts
// Production-ready CLI using Cobra, Viper, Survey, and Color
// Usage: ktx [flags] [commands]

// version is set at build-time via -ldflags (default: v0.1.0)
var version = "v1.0.0"

func main() {
	rootCmd := &cobra.Command{
		Use:   "ktx [context]",
		Short: "Easily switch Kubernetes contexts",
		Args:  cobra.MaximumNArgs(1),
		Run:   runKtx,
		Long: `ktx is a fast CLI for switching Kubernetes contexts.

Usage:
  ktx                 # interactive mode
  ktx <context>       # switch to context directly
  ktx list            # list all contexts
  ktx version         # show version
`,
	}

	// Add --kubeconfig flag
	rootCmd.Flags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file")
	viper.BindPFlag("kubeconfig", rootCmd.Flags().Lookup("kubeconfig"))

	// Version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ktx",
		Run: func(cmd *cobra.Command, args []string) {
			color.Blue("🔷 ktx version %s", version)
		},
	}
	rootCmd.AddCommand(versionCmd)

	// List command
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all available Kubernetes contexts",
		Run:   runList,
	}
	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		color.Red("❌ Error: %v", err)
		os.Exit(1)
	}
}

// runKtx is the main context switching logic
func runKtx(cmd *cobra.Command, args []string) {
	kubeconfigPath := resolveKubeconfigPath()
	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		color.Red("Failed to load kubeconfig: %v", err)
		os.Exit(2)
	}

	contexts := getContextNames(config)
	if len(contexts) == 0 {
		color.Yellow("No Kubernetes contexts found in kubeconfig.")
		os.Exit(3)
	}

	// 1. If positional argument is given, use it as context
	if len(args) > 0 && args[0] != "" {
		argContext := args[0]
		if _, ok := config.Contexts[argContext]; !ok {
			color.Red("❌ Context '%s' not found.", argContext)
			os.Exit(4)
		}
		if argContext == config.CurrentContext {
			color.Cyan("ℹ️  '%s' is already the current context.", argContext)
			return
		}
		config.CurrentContext = argContext
		err = clientcmd.WriteToFile(*config, kubeconfigPath)
		if err != nil {
			color.Red("Failed to update kubeconfig: %v", err)
			os.Exit(5)
		}
		color.Green("✅ Switched to context: %s", argContext)
		return
	}


	// 3. Interactive: prompt user
	selected := ""
	promptOptions := make([]string, 0, len(contexts))
	promptOptions = append(promptOptions, contexts...)
	prompt := &survey.Select{
		Message: "🔄 Choose a Kubernetes context:",
		Options: promptOptions,
		Default: config.CurrentContext,
	}
	err = survey.AskOne(prompt, &selected)
	if err != nil {
		color.Red("❌ Prompt failed: %v", err)
		os.Exit(6)

	}
	if selected == config.CurrentContext {
		color.Cyan("ℹ️  '%s' is already the current context.", selected)
		return
	}
	config.CurrentContext = selected
	err = clientcmd.WriteToFile(*config, kubeconfigPath)
	if err != nil {
		color.Red("Failed to update kubeconfig: %v", err)
		os.Exit(7)
	}
	color.Green("✅ Switched to context: %s", selected)
}

// runList prints all contexts, highlighting the current one
func runList(cmd *cobra.Command, args []string) {
	kubeconfigPath := resolveKubeconfigPath()
	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		color.Red("Failed to load kubeconfig: %v", err)
		os.Exit(2)
	}
	contexts := getContextNames(config)
	if len(contexts) == 0 {
		color.Yellow("No Kubernetes contexts found in kubeconfig.")
		os.Exit(3)
	}
	for _, ctx := range contexts {
		if ctx == config.CurrentContext {
			color.Green("* %s (current) ✅", ctx)
		} else {
			color.White("  %s", ctx) // Not current, no emoji
		}
	}
}

// resolveKubeconfigPath determines the kubeconfig path (flag > env > default)
func resolveKubeconfigPath() string {
	if path := viper.GetString("kubeconfig"); path != "" {
		return path
	}
	if env := os.Getenv("KUBECONFIG"); env != "" {
		return env
	}
	home, err := os.UserHomeDir()
	if err != nil {
		color.Red("❌ Unable to determine home directory: %v", err)
		os.Exit(10)
	}
	return filepath.Join(home, ".kube", "config")
}

// getContextNames returns all context names from kubeconfig
func getContextNames(config *clientcmdapi.Config) []string {
	contexts := make([]string, 0, len(config.Contexts))
	contexts = append(contexts, getMapKeys(config.Contexts)...)
	return contexts
}

// getMapKeys returns the keys of a map[string]*clientcmdapi.Context as a slice of strings
func getMapKeys(m map[string]*clientcmdapi.Context) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}