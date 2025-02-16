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
	Tweets       collections.Map[string, birdFeed.Tweet]
	AuthorTweets collections.Map[collections.Pair[string, string], bool] // (author, tweetID) -> Tweet
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
		Tweets:       collections.NewMap(sb, birdFeed.TweetsKey, "tweets", collections.StringKey, codec.CollValue[birdFeed.Tweet](cdc)),
		AuthorTweets: collections.NewMap(sb, birdFeed.AuthorTweetsKey, "author_tweets",
			collections.PairKeyCodec(collections.StringKey, collections.StringKey),
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
