package keeper

import (
	"context"

	"cosmossdk.io/collections"

	"github.com/bitcolibri/birdFeed"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *birdFeed.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, user := range data.IndexedUsers {
		if err := k.Users.Set(ctx, user.Index, *user.User); err != nil {
			return err
		}
	}

	for _, follow := range data.IndexedFollows {
		key := collections.Join(follow.K1, follow.K2)
		if err := k.Follows.Set(ctx, key); err != nil {
			return err
		}
	}

	for _, follow := range data.IndexedFollowers {
		key := collections.Join(follow.K1, follow.K2)
		if err := k.Followers.Set(ctx, key); err != nil {
			return err
		}
	}
	for _, tweet := range data.IndexedTweets {
		if err := k.Tweets.Set(ctx, tweet.Index, *tweet.Tweet); err != nil {
			return err
		}
	}

	for _, like := range data.IndexedLikes {
		key := collections.Join(like.K1, like.K2)
		if err := k.Likes.Set(ctx, key, like.Like); err != nil {
			return err
		}
	}

	for _, comment := range data.IndexedComments {
		key := collections.Join3(comment.K1, comment.K2, comment.K3)
		if err := k.Comments.Set(ctx, key, comment.Comment); err != nil {
			return err
		}
	}
	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*birdFeed.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var indexedUsers []birdFeed.IndexedUser
	if err := k.Users.Walk(ctx, nil, func(index string, user birdFeed.User) (bool, error) {
		indexedUsers = append(indexedUsers, birdFeed.IndexedUser{
			Index: index,
			User:  &user,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	var indexedFollows []birdFeed.IndexedFollow
	if err := k.Follows.Walk(ctx, nil, func(index collections.Pair[string, string]) (bool, error) {
		indexedFollows = append(indexedFollows, birdFeed.IndexedFollow{
			K1: index.K1(),
			K2: index.K2(),
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	var indexedFollowers []birdFeed.IndexedFollow
	if err := k.Followers.Walk(ctx, nil, func(index collections.Pair[string, string]) (bool, error) {
		indexedFollowers = append(indexedFollowers, birdFeed.IndexedFollow{
			K1: index.K1(),
			K2: index.K2(),
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	var indexedTweets []birdFeed.IndexedTweet
	if err := k.Tweets.Walk(ctx, nil, func(index string, tweet birdFeed.Tweet) (bool, error) {
		indexedTweets = append(indexedTweets, birdFeed.IndexedTweet{
			Index: index,
			Tweet: &tweet,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	var indexedLikes []birdFeed.IndexedLike
	if err := k.Likes.Walk(ctx, nil, func(index collections.Pair[string, string], _ bool) (bool, error) {
		indexedLikes = append(indexedLikes, birdFeed.IndexedLike{
			K1:   index.K1(),
			K2:   index.K2(),
			Like: true,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	var indexedComments []birdFeed.IndexedComment
	if err := k.Comments.Walk(ctx, nil, func(index collections.Triple[string, string, string], _ bool) (bool, error) {
		indexedComments = append(indexedComments, birdFeed.IndexedComment{
			K1:      index.K1(),
			K2:      index.K2(),
			K3:      index.K3(),
			Comment: true,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &birdFeed.GenesisState{
		Params:           params,
		IndexedUsers:     indexedUsers,
		IndexedFollows:   indexedFollows,
		IndexedFollowers: indexedFollowers,
		IndexedTweets:    indexedTweets,
		IndexedLikes:     indexedLikes,
		IndexedComments:  indexedComments,
	}, nil
}
