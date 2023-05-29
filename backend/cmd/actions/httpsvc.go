package actions

import (
	"github.com/spf13/cobra"
	service "tabelf/backend/service/api"
)

func RunHTTPSvc(cmd *cobra.Command, args []string) {
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		Fatalf("get env failed: %v\n", err)
	}
	service.StartHTTP(env)
}

func init() {
	cmd := &cobra.Command{
		Use: "start_http",
		Run: RunHTTPSvc,
	}
	rootCmd.AddCommand(cmd)
}
