package cmd

import (
	"fmt"
	"log"

	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"

	"github.com/spf13/cobra"
)

type getOptions struct {
	id   string
	list bool
}

// getCmd represents the get command
func NewCmdGet() *cobra.Command {
	var (
		o = &getOptions{}
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get firebase authentication user",
		Run: func(cmd *cobra.Command, args []string) {
			runCmdGet(o)
		},
	}

	cmd.Flags().BoolVarP(&o.list, "list", "l", false, "Get user list")
	cmd.Flags().StringVarP(&o.id, "id", "i", "", "Get user UID")
	return cmd
}

func init() {
}

func runCmdGet(opt *getOptions) {
	if opt.list {
		getUsers()
	} else {
		if len(opt.id) != 0 {
			getUser(opt.id)
		} else {
			fmt.Println("Required firebase authentication UID")
		}
	}
}

func getUser(uid string) *auth.UserRecord {
	ctx, client := getFirebaseClient()
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u.UserInfo.Email)
	return u
}

func getUsers() {
	ctx, client := getFirebaseClient()
	iter := client.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error listing users: %s\n", err)
		}
		log.Printf("Success user user: %v\n", user.UserInfo)
	}

	pager := iterator.NewPager(client.Users(ctx, ""), 100, "")
	for {
		var users []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			log.Fatalf("paging error %v\n", err)
		}
		for _, u := range users {
			log.Printf("read user user: %v\n", u)
		}
		if nextPageToken == "" {
			break
		}
	}
}
