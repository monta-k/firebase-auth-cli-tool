package cmd

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var getUserByUIDCmd = &cobra.Command{
	Use:   "get-user-by-uid",
	Short: "get Firebase Auth User By UID",
	RunE: func(cmd *cobra.Command, args []string) error {
		tenantID, err := cmd.PersistentFlags().GetString("tenant")
		if err != nil {
			return err
		}

		uid, err := cmd.PersistentFlags().GetString("uid")
		if err != nil {
			return err
		}

		ctx := context.Background()

		opt := option.WithCredentialsFile("./secret/serviceAccount.json")
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			return fmt.Errorf("error initializing app: %v\n", err)
		}

		client, err := app.Auth(ctx)
		if err != nil {
			return fmt.Errorf("error getting Auth client: %v\n", err)
		}

		tenantClient, err := client.TenantManager.AuthForTenant(tenantID)
		if err != nil {
			return fmt.Errorf("error getting Tenant Auth client: %v\n", err)
		}

		user, err := tenantClient.GetUser(ctx, uid)
		if err != nil {
			return fmt.Errorf("error getting user: %v\n", err)
		}

		fmt.Println(user)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getUserByUIDCmd)

	getUserByUIDCmd.PersistentFlags().StringP("tenant", "t", "tenant_id", "identity platform tenant id")
	getUserByUIDCmd.PersistentFlags().StringP("uid", "u", "user_id", "Firebase Auth User UID")
}
