package keeper

import (
	"context"

	"github.com/bitcolibri/birdFeed"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *birdFeed.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, tweet := range data.IndexedTweets {
		if err := k.Tweets.Set(ctx, tweet.Index, *tweet.Tweet); err != nil {
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

	return &birdFeed.GenesisState{
		Params:        params,
		IndexedTweets: indexedTweets,
	}, nil
}
