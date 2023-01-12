package RedditHandler

import (
	"context"
	"fmt"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

/*
TODO: 1 - Get un post
*/
type RedditHandler struct {
	client *reddit.Client
}

/* TODO: DB pour garder post déjà used par la chaine
 */
func (redditHandler *RedditHandler) GetClient(id string, secret string, username string, password string) error {
	credentials := reddit.Credentials{ID: id, Secret: secret, Username: username, Password: password}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return err
	}

	redditHandler.client = client
	return nil
}

func (redditHandler *RedditHandler) GetTopPosts(subredditName string) error {
	posts, _, err := redditHandler.client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 5,
		},
		Time: "all",
	})
	fmt.Printf("Received %d posts.\n", len(posts))
	if err != nil {
		return err
	}

	return nil
}
