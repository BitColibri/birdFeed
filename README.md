# Bird Feed

## Overview

Bird Feed is a Cosmos SDK module that allows users to create and manage their own bird feeders.
It is a simple implementation of Twitter.
> **Note**: This module is not production ready and is only for learning purposes.

## Features

- Create a user
- Follow other users
- Publish posts
- Like posts
- Comment on posts
- Search for users, posts, comments

## Collections Overview

The BirdFeed module uses several collections to store and manage social media data:

### Core Collections

- **Params**: Stores global module parameters
  - Type: `collections.Item`
  - Single instance storage, no key required

- **Users**: Stores user profile information
  - Type: `collections.Map[string, birdFeed.User]`
  - Key: User address (string)

### Relationship Collections

- **Followers**: Tracks who follows whom
  - Type: `collections.KeySet[collections.Pair[string, string]]`
  - Key Pair: `[follower_address, followed_address]`
  - Represents: "follower_address follows followed_address"

- **Follows**: Inverse index of followers relationship
  - Type: `collections.KeySet[collections.Pair[string, string]]`
  - Key Pair: `[followed_address, follower_address]`
  - Represents: "followed_address is followed by follower_address"
  - Enables efficient querying of a user's followers

### Content Collections

- **Tweets**: Stores tweet content
  - Type: `collections.Map[string, birdFeed.Tweet]`
  - Key: Tweet ID (string)

- **AuthorTweets**: Links authors to their tweets
  - Type: `collections.KeySet[collections.Pair[string, string]]`
  - Key Pair: `[author_address, tweet_id]`
  - Enables efficient querying of tweets by author

- **Likes**: Tracks tweet likes
  - Type: `collections.KeySet[collections.Pair[string, string]]`
  - Key Pair: `[user_address, tweet_id]`
  - Represents: "user_address liked tweet_id"

- **Comments**: Stores comment relationships
  - Type: `collections.KeySet[collections.Triple[string, string, string]]`
  - Key Triple: `[tweet_id, user_address, comment_id]`
  - Represents: "user_address commented on tweet_id with comment_id"




