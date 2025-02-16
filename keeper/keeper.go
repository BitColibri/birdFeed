package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/bitcolibri/birdFeed"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// state management
	Schema       collections.Schema
	Params       collections.Item[birdFeed.Params]
	Users        collections.Map[string, birdFeed.User]
	Followers    collections.KeySet[collections.Pair[string, string]]
	Follows      collections.KeySet[collections.Pair[string, string]]
	Tweets       collections.Map[string, birdFeed.Tweet]
	AuthorTweets collections.Map[collections.Pair[string, string], bool]
	Likes        collections.Map[collections.Pair[string, string], bool]
	Comments     collections.Map[collections.Triple[string, string, string], bool]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, birdFeed.ParamsKey, "params", codec.CollValue[birdFeed.Params](cdc)),
		Users:        collections.NewMap(sb, birdFeed.UsersKey, "users", collections.StringKey, codec.CollValue[birdFeed.User](cdc)),
		Followers:    collections.NewKeySet(sb, birdFeed.FollowersKey, "followers", collections.PairKeyCodec(collections.StringKey, collections.StringKey)),
		Follows:      collections.NewKeySet(sb, birdFeed.FollowsKey, "follows", collections.PairKeyCodec(collections.StringKey, collections.StringKey)),
		Tweets:       collections.NewMap(sb, birdFeed.TweetsKey, "tweets", collections.StringKey, codec.CollValue[birdFeed.Tweet](cdc)),
		AuthorTweets: collections.NewMap(sb, birdFeed.AuthorTweetsKey, "author_tweets",
			collections.PairKeyCodec(collections.StringKey, collections.StringKey),
			collections.BoolValue,
		),
		Likes: collections.NewMap(sb, birdFeed.LikesKey, "likes",
			collections.PairKeyCodec(collections.StringKey, collections.StringKey),
			collections.BoolValue,
		),
		Comments: collections.NewMap(sb, birdFeed.CommentsKey, "comments",
			collections.TripleKeyCodec(collections.StringKey, collections.StringKey, collections.StringKey),
			collections.BoolValue,
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}
