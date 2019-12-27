package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type deleteOptions struct {
	id string
}

// deleteCmd represents the delete command
func NewCmdDelete() *cobra.Command {
	var (
		o = &deleteOptions{}
	)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete firebase authentication user",
		Run: func(cmd *cobra.Command, args []string) {
			runCmdDelete(o)
		},
	}

	cmd.Flags().StringVarP(&o.id, "id", "i", "", "Delete user UID")
	return cmd
}

func init() {
}

func runCmdDelete(opt *deleteOptions) {
	if len(opt.id) != 0 {
		deleteUser(opt.id)
	} else {
		fmt.Println("Required firebase authentication UID")
	}
}

func deleteUser(uid string) {
	ctx, client := getFirebaseClient()
	err := client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("Error deleting user: %v\n", err)
	}
	log.Printf("Success deleted user: %s\n", uid)
}
