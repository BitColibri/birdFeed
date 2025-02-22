package keeper

import (
	"context"
	"errors"
	"fmt"

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
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &birdFeed.QueryGetTweetResponse{}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	var comments []*birdFeed.Tweet
	rng := collections.NewPrefixedTripleRange[string, string, string](msg.Id)
	err = s.k.Comments.Walk(ctx, rng, func(key collections.Triple[string, string, string]) (bool, error) {
		fmt.Printf("Found comment key: K1=%s, K2=%s, K3=%s\n", key.K1(), key.K2(), key.K3())

		commentId := key.K3()
		comment, err := s.k.Tweets.Get(ctx, commentId)
		if err != nil {
			return false, err
		}
		comments = append(comments, &comment)
		return false, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &birdFeed.QueryGetTweetResponse{
		Tweet:    &tweet,
		Comments: comments,
	}, nil
}

func (s queryServer) GetAuthorTweets(ctx context.Context, msg *birdFeed.QueryGetAuthorTweetsRequest) (*birdFeed.QueryGetAuthorTweetsResponse, error) {
	var tweets []*birdFeed.Tweet

	rng := collections.NewPrefixedPairRange[string, string](msg.Author)
	err := s.k.AuthorTweets.Walk(ctx, rng, func(key collections.Pair[string, string]) (bool, error) {
		tweetID := key.K2()

		tweet, err := s.k.Tweets.Get(ctx, tweetID)
		if err != nil {
			return false, err
		}
		tweet.Id = tweetID
		tweets = append(tweets, &tweet)
		return false, nil // continue iteration
	})

	return &birdFeed.QueryGetAuthorTweetsResponse{
		Tweets: tweets,
	}, err
}

func (s queryServer) GetTweetLikes(ctx context.Context, msg *birdFeed.QueryGetTweetLikesRequest) (*birdFeed.QueryGetTweetLikesResponse, error) {
	var likes []string

	rng := collections.NewPrefixedPairRange[string, string](msg.Id)
	err := s.k.Likes.Walk(ctx, rng, func(key collections.Pair[string, string]) (bool, error) {
		likes = append(likes, key.K2())
		return false, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &birdFeed.QueryGetTweetLikesResponse{
		Likes: likes,
	}, nil
}

func (s queryServer) GetUser(ctx context.Context, msg *birdFeed.QueryGetUserRequest) (*birdFeed.QueryGetUserResponse, error) {
	user, err := s.k.Users.Get(ctx, msg.Address)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &birdFeed.QueryGetUserResponse{User: &user}, nil
}

func (s queryServer) GetUserFollowers(ctx context.Context, msg *birdFeed.QueryGetUserFollowersRequest) (*birdFeed.QueryGetUserFollowersResponse, error) {
	var followers []string

	rng := collections.NewPrefixedPairRange[string, string](msg.Address)
	err := s.k.Followers.Walk(ctx, rng, func(key collections.Pair[string, string]) (bool, error) {
		followers = append(followers, key.K2())
		return false, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &birdFeed.QueryGetUserFollowersResponse{
		Followers: followers,
	}, nil
}

func (s queryServer) GetUserFollows(ctx context.Context, msg *birdFeed.QueryGetUserFollowsRequest) (*birdFeed.QueryGetUserFollowsResponse, error) {
	var follows []string

	rng := collections.NewPrefixedPairRange[string, string](msg.Address)
	err := s.k.Follows.Walk(ctx, rng, func(key collections.Pair[string, string]) (bool, error) {
		follows = append(follows, key.K2())
		return false, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &birdFeed.QueryGetUserFollowsResponse{
		Follows: follows,
	}, nil
}
