package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/bitcolibri/birdFeed"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ birdFeed.QueryServer = queryServer{}

type queryServer struct {
	k Keeper
}

func NewQueryServerImpl(k Keeper) birdFeed.QueryServer {
	return &queryServer{k: k}
}

func (s queryServer) GetTweet(ctx context.Context, msg *birdFeed.QueryGetTweetRequest) (*birdFeed.QueryGetTweetResponse, error) {
	tweet, err := s.k.Tweets.Get(ctx, msg.Id)
	if err == nil {
		return &birdFeed.QueryGetTweetResponse{
			Tweet: &tweet,
		}, nil
	}

	if errors.Is(err, collections.ErrNotFound) {
		return &birdFeed.QueryGetTweetResponse{}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}

func (s queryServer) GetAuthorTweets(ctx context.Context, msg *birdFeed.QueryGetAuthorTweetsRequest) (*birdFeed.QueryGetAuthorTweetsResponse, error) {
	var tweets []*birdFeed.Tweet

	rng := collections.NewPrefixedPairRange[string, string](msg.Author)
	err := s.k.AuthorTweets.Walk(ctx, rng, func(key collections.Pair[string, string], _ bool) (bool, error) {
		tweetID := key.K2()

		tweet, err := s.k.Tweets.Get(ctx, tweetID)
		if err != nil {
			return false, err
		}
		tweets = append(tweets, &tweet)
		return false, nil // continue iteration
	})

	return &birdFeed.QueryGetAuthorTweetsResponse{
		Tweets: tweets,
	}, err
}
