package bootstrap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/frain-dev/convoy/database/postgres"
	"github.com/frain-dev/convoy/datastore"
	"github.com/frain-dev/convoy/internal/pkg/cli"
	"github.com/frain-dev/convoy/pkg/log"
	"github.com/frain-dev/convoy/util"
	"github.com/oklog/ulid/v2"
	"time"

	"github.com/spf13/cobra"
)

func AddBootstrapCommand(a *cli.App) *cobra.Command {
	var firstName string
	var lastName string
	var format string
	var email string

	cmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "bootstrap creates a new user account",
		Annotations: map[string]string{
			"ShouldBootstrap": "false",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if format != "json" && format != "human" {
				return errors.New("unsupported output format")
			}

			if util.IsStringEmpty(email) {
				return errors.New("email is required")
			}

			password, err := util.GenerateSecret()
			if err != nil {
				return err
			}

			p := datastore.Password{Plaintext: password}
			err = p.GenerateHash()
			if err != nil {
				return err
			}

			user := &datastore.User{
				UID:       ulid.Make().String(),
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  string(p.Hash),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			userRepo := postgres.NewUserRepo(a.DB)
			err = userRepo.CreateUser(context.Background(), user)
			if err != nil {
				return err
			}

			type JsonUser struct {
				FirstName string `json:"first_name,omitempty"`
				LastName  string `json:"last_name,omitempty"`
				Email     string `json:"email,omitempty"`
				Password  string `json:"password,omitempty"`
			}

			jsUser := &JsonUser{
				Email:     user.Email,
				Password:  p.Plaintext,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			switch format {
			case "json":
				data, err := json.MarshalIndent(jsUser, "", "    ")
				if err != nil {
					log.Fatalf("Error printing config: %v\n", err)
				}

				fmt.Println(string(data))
				break
			case "human":
				fmt.Printf("Email: %s\n", jsUser.Email)
				fmt.Printf("Password: %s\n", jsUser.Password)
				fmt.Printf("First Name: %s\n", jsUser.FirstName)
				fmt.Printf("Last Name: %s\n", jsUser.LastName)
				break
			default:
				return errors.New("unsupported output format")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&email, "email", "", "Email")
	cmd.Flags().StringVar(&firstName, "first-name", "admin", "Email")
	cmd.Flags().StringVar(&lastName, "last-name", "admin", "Email")
	cmd.Flags().StringVar(&format, "format", "json", "Output Format")

	return cmd
}
