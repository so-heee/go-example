package cmd

import (
	"fmt"
	"log"

	"firebase.google.com/go/auth"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	id       string
	email    string
	phone    string
	password string
	name     string
	url      string
	disabled bool
}

// updateCmd represents the update command
func NewCmdUpdate() *cobra.Command {
	var (
		o = &updateOptions{}
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Get firebase authentication user",
		Run: func(cmd *cobra.Command, args []string) {
			runCmdUpdate(o)
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

func runCmdUpdate(opt *updateOptions) {
	if len(opt.id) != 0 {
		update, user := setUserToUpdate(opt)
		if update {
			updateUser(opt.id, user)
		} else {
			fmt.Println("Required update params")
		}
	} else {
		fmt.Println("Required firebase authentication UID")
	}
}

func updateUser(uid string, user *auth.UserToUpdate) {
	ctx, client := getFirebaseClient()
	u, err := client.UpdateUser(ctx, uid, user)
	if err != nil {
		log.Fatalf("Error updating user: %v\n", err)
	}
	log.Printf("Success updated user: %v\n", u)
}

func setUserToUpdate(opt *updateOptions) (bool, *auth.UserToUpdate) {
	user := &auth.UserToUpdate{}
	update := false
	if len(opt.email) != 0 {
		user.Email(opt.email)
		user.EmailVerified(true)
		update = true
	}
	if len(opt.phone) != 0 {
		user.PhoneNumber(opt.phone)
		update = true
	}
	if len(opt.password) != 0 {
		user.Password(opt.password)
		update = true
	}
	if len(opt.name) != 0 {
		user.DisplayName(opt.name)
		update = true
	}
	if len(opt.url) != 0 {
		user.PhotoURL(opt.url)
		update = true
	}
	if opt.disabled {
		user.Disabled(false)
		update = true
	}
	return update, user
}
