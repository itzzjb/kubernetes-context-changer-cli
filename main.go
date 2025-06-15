package main

import (
	"os"
	"path/filepath"
	"strings"

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
var version = "v0.1.0"

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
  ktx -c <context>    # switch to context directly (flag, legacy)
  ktx list            # list all contexts
  ktx version         # show version
`,
	}

	// Add --context flag for non-interactive switching (legacy/optional)
	rootCmd.Flags().StringP("context", "c", "", "Context to switch to (non-interactive, legacy; prefer positional argument)")
	rootCmd.Flags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file")
	viper.BindPFlag("kubeconfig", rootCmd.Flags().Lookup("kubeconfig"))
	viper.BindPFlag("context", rootCmd.Flags().Lookup("context"))

	// Version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ktx",
		Run: func(cmd *cobra.Command, args []string) {
			color.Blue("ktx version %s", version)
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
		color.Red("Error: %v", err)
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
			color.Red("Context '%s' not found.", argContext)
			os.Exit(4)
		}
		if argContext == config.CurrentContext {
			color.Cyan("'%s' is already the current context.", argContext)
			return
		}
		config.CurrentContext = argContext
		err = clientcmd.WriteToFile(*config, kubeconfigPath)
		if err != nil {
			color.Red("Failed to update kubeconfig: %v", err)
			os.Exit(5)
		}
		color.Green("Switched to context: %s", argContext)
		return
	}

	// 2. If --context flag is given, use it
	flagContext := viper.GetString("context")
	if flagContext != "" {
		if _, ok := config.Contexts[flagContext]; !ok {
			color.Red("Context '%s' not found.", flagContext)
			os.Exit(4)
		}
		if flagContext == config.CurrentContext {
			color.Cyan("'%s' is already the current context.", flagContext)
			return
		}
		config.CurrentContext = flagContext
		err = clientcmd.WriteToFile(*config, kubeconfigPath)
		if err != nil {
			color.Red("Failed to update kubeconfig: %v", err)
			os.Exit(5)
		}
		color.Green("Switched to context: %s", flagContext)
		return
	}

	// 3. Interactive: prompt user
	selected := ""
	promptOptions := make([]string, 0, len(contexts))
	for _, ctx := range contexts {
		if ctx == config.CurrentContext {
			promptOptions = append(promptOptions, color.New(color.FgGreen, color.Bold).Sprint(ctx+" (current)"))
		} else {
			promptOptions = append(promptOptions, ctx)
		}
	}
	prompt := &survey.Select{
		Message: "Choose a Kubernetes context:",
		Options: promptOptions,
		Default: color.New(color.FgGreen, color.Bold).Sprint(config.CurrentContext+" (current)"),
	}
	err = survey.AskOne(prompt, &selected)
	if err != nil {
		color.Red("Prompt failed: %v", err)
		os.Exit(6)
	}
	// Remove (current) and color codes for comparison
	selected = stripContextName(selected)
	if selected == config.CurrentContext {
		color.Cyan("'%s' is already the current context.", selected)
		return
	}
	config.CurrentContext = selected
	err = clientcmd.WriteToFile(*config, kubeconfigPath)
	if err != nil {
		color.Red("Failed to update kubeconfig: %v", err)
		os.Exit(7)
	}
	color.Green("Switched to context: %s", selected)
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
			color.Green("* %s (current)", ctx)
		} else {
			color.White("  %s", ctx)
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
		color.Red("Unable to determine home directory: %v", err)
		os.Exit(10)
	}
	return filepath.Join(home, ".kube", "config")
}

// getContextNames returns all context names from kubeconfig
func getContextNames(config *clientcmdapi.Config) []string {
	contexts := make([]string, 0, len(config.Contexts))
	for name := range config.Contexts {
		contexts = append(contexts, name)
	}
	return contexts
}

// stripContextName removes color and (current) marker for comparison
func stripContextName(s string) string {
	// Remove (current) and any color codes
	plain := s
	plain = strings.ReplaceAll(plain, " (current)", "")
	plain = color.New().SprintFunc()(plain) // Remove color
	return plain
}

