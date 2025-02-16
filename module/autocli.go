package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	birdFeedv1 "github.com/bitcolibri/birdFeed/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: birdFeedv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetTweet",
					Use:       "get-tweet id",
					Short:     "Get tweet by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
				{
					RpcMethod: "GetAuthorTweets",
					Use:       "author-tweets author",
					Short:     "Get tweets by author",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "author"},
					},
				},
				{
					RpcMethod: "GetTweetLikes",
					Use:       "likes",
					Short:     "Get all likes of a tweet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: birdFeedv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "PublishTweet",
					Use:       "tweet",
					Short:     "Publish a new tweet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "content"},
						{ProtoField: "hashtags"},
					},
				},
				{
					RpcMethod: "LikeTweet",
					Use:       "like",
					Short:     "Like a tweet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "tweetID"},
					},
				},
				{
					RpcMethod: "UnlikeTweet",
					Use:       "unlike",
					Short:     "Unlike a tweet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "tweetID"},
					},
				},
				{
					RpcMethod: "CommentTweet",
					Use:       "comment",
					Short:     "Unlike a tweet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "tweetID"},
						{ProtoField: "content"},
						{ProtoField: "hashtags"},
					},
				},
			},
			SubCommands:          nil,
			EnhanceCustomCommand: false,
		},
	}
}
