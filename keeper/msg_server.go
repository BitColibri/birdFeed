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

func (s msgServer) LikeTweet(ctx context.Context, msg *birdFeed.MsgLikeTweet) (*birdFeed.MsgLikeTweetResponse, error) {
	// check if tweet exists
	tweet, err := s.Tweets.Get(ctx, msg.TweetID)
	if err != nil {
		return nil, err
	}

	// check if user has already liked the tweet
	likeKey := collections.Join(msg.TweetID, msg.From)
	liked, err := s.Likes.Get(ctx, likeKey)
	if err != nil {
		if !errors.Is(err, collections.ErrNotFound) {
			return nil, err
		}
	}
	if liked {
		return nil, fmt.Errorf("%s already liked this tweet", msg.From)
	}

	// set the like
	if err := s.Likes.Set(ctx, likeKey, true); err != nil {
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
	liked, err := s.Likes.Get(ctx, likeKey)
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
	if err := s.Comments.Set(ctx, commentID, true); err != nil {
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
	if err := s.AuthorTweets.Set(ctx, authorTweetsKey, true); err != nil {
		return err
	}

	return nil
}
