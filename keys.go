package birdFeed

import "cosmossdk.io/collections"

const ModuleName = "birdFeed"
const MaxIndexLength = 256

var (
	ParamsKey       = collections.NewPrefix("Params")
	TweetsKey       = collections.NewPrefix("Tweets")
	AuthorTweetsKey = collections.NewPrefix("AuthorTweets")
	LikesKey        = collections.NewPrefix("Likes")
	CommentsKey     = collections.NewPrefix("Comments")
)
