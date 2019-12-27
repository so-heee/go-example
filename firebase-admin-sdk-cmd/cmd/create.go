package cmd

import (
	"log"

	"firebase.google.com/go/auth"
	"github.com/spf13/cobra"
)

type createOptions struct {
	id       string
	email    string
	phone    string
	password string
	name     string
	url      string
	disabled bool
}

// createCmd represents the create command
func NewCmdCreate() *cobra.Command {
	var (
		o = &createOptions{}
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create firebase authentication user",
		Run: func(cmd *cobra.Command, args []string) {
			runCmdCreate(o)
		},
	}
	cmd.Flags().StringVarP(&o.id, "id", "i", "", "Create user UID")
	cmd.Flags().StringVarP(&o.email, "email", "e", "", "Create user Email")
	cmd.Flags().StringVarP(&o.phone, "phone", "p", "", "Create user Phone Number")
	cmd.Flags().StringVarP(&o.password, "password", "w", "", "Create user Password")
	cmd.Flags().StringVarP(&o.name, "name", "n", "", "Create user Name")
	cmd.Flags().StringVarP(&o.url, "url", "u", "", "Create user Photo Url")
	cmd.Flags().BoolVarP(&o.disabled, "disabled", "d", false, "Create user Disabled")
	return cmd
}

func init() {
}

func runCmdCreate(opt *createOptions) {
	user := setUserToCreate(opt)
	createUser(user)
}

func createUser(user *auth.UserToCreate) *auth.UserRecord {
	ctx, client := getFirebaseClient()
	u, err := client.CreateUser(ctx, user)
	if err != nil {
		log.Fatalf("Error creating user: %v\n", err)
	}
	log.Printf("Success created user: %v\n", u)
	return u
}

func setUserToCreate(opt *createOptions) *auth.UserToCreate {
	user := &auth.UserToCreate{}
	if len(opt.id) != 0 {
		user.UID(opt.id)
	}
	if len(opt.email) != 0 {
		user.Email(opt.email)
		user.EmailVerified(true)
	}
	if len(opt.phone) != 0 {
		user.PhoneNumber(opt.phone)
	}
	if len(opt.password) != 0 {
		user.Password(opt.password)
	}
	if len(opt.name) != 0 {
		user.DisplayName(opt.name)
	}
	if len(opt.url) != 0 {
		user.PhotoURL(opt.url)
	}
	if opt.disabled {
		user.Disabled(false)
	}
	return user
}
