package cmd

import (
    "errors"
	"fmt"

	"github.com/spf13/cobra"
)

// timescaledbGetPasswordCmd represents the timescaledb get-password command
var timescaledbGetPasswordCmd = &cobra.Command{
	Use:   "get-password",
	Short: "Gets the TimescaleDB/PostgreSQL password for a specific user",
	RunE:  timescaledbGetPassword,
}

func init() {
	timescaledbCmd.AddCommand(timescaledbGetPasswordCmd)
    timescaledbGetPasswordCmd.Flags().StringP("user", "u", "postgres", "user whose password to get")
}

func timescaledbGetPassword(cmd *cobra.Command, args []string) error {
    var err error

    if len(args) != 0 {
        return errors.New("\"ts-obs timescaledb get-password\" requires 0 arguments")
    }

    var user string
    user, err = cmd.Flags().GetString("user")
    if err != nil {
        return err
    }

    secret, err := kubeGetSecret("ts-obs-timescaledb-passwords")
    if err != nil {
        return err
    }

    pass := secret.Data[user]
    fmt.Printf("Password: %v\n", string(pass))

    return nil
}