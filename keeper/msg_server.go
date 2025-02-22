package keeper

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"github.com/bitcolibri/birdFeed"
)

var _ birdFeed.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) birdFeed.MsgServer {
	return &msgServer{keeper}
}

func (s msgServer) InitUser(ctx context.Context, msg *birdFeed.MsgInitUser) (*birdFeed.MsgInitUserResponse, error) {
	hasUser, err := s.Users.Has(ctx, msg.Address)
	if err != nil {
		return nil, err
	}
	if hasUser {
		return nil, fmt.Errorf("user %s already exists", msg.Address)
	}

	user := birdFeed.User{
		Address: msg.Address,
		Alias:   msg.Alias,
		Picture: msg.Picture,
	}

	if err := s.Users.Set(ctx, msg.Address, user); err != nil {
		return nil, err
	}

	return &birdFeed.MsgInitUserResponse{}, nil
}

func (s msgServer) FollowUser(ctx context.Context, msg *birdFeed.MsgFollowUser) (*birdFeed.MsgFollowUserResponse, error) {
	from, err := s.Users.Get(ctx, msg.From)
	if err != nil {
		return nil, err
	}

	to, err := s.Users.Get(ctx, msg.To)
	if err != nil {
		return nil, err
	}

	if from.Address == to.Address {
		return nil, fmt.Errorf("you cannot follow yourself")
	}

	followingKey := collections.Join(from.Address, to.Address)
	if err = s.Follows.Set(ctx, followingKey); err != nil {
		return nil, err
	}

	followerKey := collections.Join(to.Address, from.Address)
	if err = s.Followers.Set(ctx, followerKey); err != nil {
		return nil, err
	}

	to.Followers++
	if err := s.Users.Set(ctx, to.Address, to); err != nil {
		return nil, err
	}

	from.Follows++
	if err := s.Users.Set(ctx, from.Address, from); err != nil {
		return nil, err
	}

	return &birdFeed.MsgFollowUserResponse{}, nil
}

func (s msgServer) UnfollowUser(ctx context.Context, msg *birdFeed.MsgUnfollowUser) (*birdFeed.MsgUnfollowUserResponse, error) {
	followingKey := collections.Join(msg.From, msg.To)
	if err := s.Followers.Remove(ctx, followingKey); err != nil {
		return nil, err
	}

	followerKey := collections.Join(msg.To, msg.From)
	if err := s.Follows.Remove(ctx, followerKey); err != nil {
		return nil, err
	}

	from, err := s.Users.Get(ctx, msg.From)
	if err != nil {
		return nil, err
	}

	from.Follows--
	if err := s.Users.Set(ctx, from.Address, from); err != nil {
		return nil, err
	}

	to, err := s.Users.Get(ctx, msg.To)
	if err != nil {
		return nil, err
	}

	to.Followers--
	if err := s.Users.Set(ctx, to.Address, to); err != nil {
		return nil, err
	}

	return &birdFeed.MsgUnfollowUserResponse{}, nil
}

func (s msgServer) PublishTweet(ctx context.Context, msg *birdFeed.MsgPublishTweet) (*birdFeed.MsgPublishTweetResponse, error) {
	timeStamp := time.Now().UnixNano()
	tweetID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%d", msg.Author, msg.Content, timeStamp))))
	tweet := birdFeed.Tweet{
		Author:    msg.Author,
		Content:   msg.Content,
		Timestamp: timeStamp,
		Hashtags:  msg.Hashtags,
	}

	if err := tweet.Validate(); err != nil {
		return nil, err
	}
	if err := s.saveTweet(ctx, tweet, tweetID); err != nil {
		return nil, err
	}

	if err := s.saveTweet(ctx, tweet, tweetID); err != nil {
		return nil, err
	}

	return &birdFeed.MsgPublishTweetResponse{}, nil
}

func (s msgServer) RemoveTweet(ctx context.Context, msg *birdFeed.MsgRemoveTweet) (*birdFeed.MsgRemoveTweetResponse, error) {
	tweet, err := s.Tweets.Get(ctx, msg.TweetID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, fmt.Errorf("tweet %s does not exist", msg.TweetID)
		}
		return nil, err
	}

	if tweet.Author != msg.Author {
		return nil, fmt.Errorf("you are not the author of this tweet")
	}

	if err := s.Tweets.Remove(ctx, msg.TweetID); err != nil {
		return nil, err
	}

	if err := s.AuthorTweets.Remove(ctx, collections.Join(tweet.Author, msg.TweetID)); err != nil {
		return nil, err
	}

	likesRange := collections.NewPrefixedPairRange[string, string](msg.TweetID)
	err = s.Likes.Walk(ctx, likesRange, func(key collections.Pair[string, string]) (stop bool, err error) {
		if err := s.Likes.Remove(ctx, key); err != nil {
			return false, err
		}
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return &birdFeed.MsgRemoveTweetResponse{}, nil
}

func (s msgServer) LikeTweet(ctx context.Context, msg *birdFeed.MsgLikeTweet) (*birdFeed.MsgLikeTweetResponse, error) {
	// check if tweet exists
	tweet, err := s.Tweets.Get(ctx, msg.TweetID)
	if err != nil {
		return nil, err
	}

	// check if user has already liked the tweet
	likeKey := collections.Join(msg.TweetID, msg.From)
	liked, err := s.Likes.Has(ctx, likeKey)
	if err != nil {
		if !errors.Is(err, collections.ErrNotFound) {
			return nil, err
		}
	}
	if liked {
		return nil, fmt.Errorf("%s already liked this tweet", msg.From)
	}

	// set the like
	if err := s.Likes.Set(ctx, likeKey); err != nil {
		return nil, err
	}

	// increment the like count
	tweet.Likes++

	// save the tweet
	if err := s.Tweets.Set(ctx, msg.TweetID, tweet); err != nil {
		return nil, err
	}

	return &birdFeed.MsgLikeTweetResponse{}, nil
}

func (s msgServer) UnlikeTweet(ctx context.Context, msg *birdFeed.MsgUnlikeTweet) (*birdFeed.MsgUnlikeTweetResponse, error) {
	// check if tweet exists
	tweet, err := s.Tweets.Get(ctx, msg.TweetID)
	if err != nil {
		return nil, err
	}

	// check if user has already liked the tweet
	likeKey := collections.Join(msg.TweetID, msg.From)
	liked, err := s.Likes.Has(ctx, likeKey)
	if err != nil {
		return nil, err
	}
	if !liked {
		return nil, fmt.Errorf("%s has not liked this tweet", msg.From)
	}

	// delete the like
	if err := s.Likes.Remove(ctx, likeKey); err != nil {
		return nil, err
	}

	// decrement the like count
	tweet.Likes--

	// save the tweet
	if err := s.Tweets.Set(ctx, msg.TweetID, tweet); err != nil {
		return nil, err
	}

	return &birdFeed.MsgUnlikeTweetResponse{}, nil
}

func (s msgServer) CommentTweet(ctx context.Context, msg *birdFeed.MsgCommentTweet) (*birdFeed.MsgCommentTweetResponse, error) {
	// check if tweet exists
	tweet, err := s.Tweets.Get(ctx, msg.TweetID)
	if err != nil {
		return nil, err
	}

	timeStamp := time.Now().UnixNano()
	comment := birdFeed.Tweet{
		Author:    msg.Author,
		Content:   msg.Content,
		Timestamp: timeStamp,
		Hashtags:  msg.Hashtags,
		ParentId:  msg.TweetID,
	}

	tweetID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%d", msg.Author, msg.Content, timeStamp))))
	commentID := collections.Join3(msg.TweetID, msg.Author, tweetID)
	if err := s.Comments.Set(ctx, commentID); err != nil {
		return nil, err
	}

	if err := s.Tweets.Set(ctx, tweetID, comment); err != nil {
		return nil, err
	}

	tweet.Responses++

	if err := s.Tweets.Set(ctx, msg.TweetID, tweet); err != nil {
		return nil, err
	}

	return &birdFeed.MsgCommentTweetResponse{}, nil
}

func (s msgServer) saveTweet(ctx context.Context, tweet birdFeed.Tweet, tweetID string) error {
	if err := s.Tweets.Set(ctx, tweetID, tweet); err != nil {
		return err
	}

	authorTweetsKey := collections.Join(tweet.Author, tweetID)
	if err := s.AuthorTweets.Set(ctx, authorTweetsKey); err != nil {
		return err
	}

	return nil
}
