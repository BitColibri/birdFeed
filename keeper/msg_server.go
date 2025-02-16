package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

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
	if err := s.Tweets.Set(ctx, tweetID, tweet); err != nil {
		return nil, err
	}

	return &birdFeed.MsgPublishTweetResponse{}, nil
}
