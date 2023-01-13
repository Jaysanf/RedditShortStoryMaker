package RedditHandler

import (
	"context"
	"fmt"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/exp/slices"
)

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

func (redditHandler *RedditHandler) GetTopPosts(subredditName string, amount int, timePeriod string) ([]*reddit.Post, error) {
	posts, _, err := redditHandler.client.Subreddit.TopPosts(context.Background(), subredditName, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: amount,
		},
		Time: timePeriod,
	})
	fmt.Printf("Received %d posts.\n", len(posts))
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (redditHandler *RedditHandler) GetUnusedPost(posts []*reddit.Post, ids []string) *reddit.Post {
	for _, post := range posts {
		if !slices.Contains(ids, post.ID) {
			return post
		}
	}
	return nil
}
