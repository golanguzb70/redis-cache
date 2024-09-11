## Redis Cache Package

The `rediscache` package provides a convenient way to implement caching using Redis. It combines Redis and hashing algorithms in one place, making it faster to implement caching in your applications.

### Installation

To install the `rediscache` package, use the following command:

```shell
go get github.com/golanguzb70/redis-cache
```

### Usage

Import the `rediscache` package in your Go code:

```go
import "github.com/golanguzb70/redis-cache"
```

Create a new Redis cache instance by calling the `New` function and passing a `Config` object:

```go
cfg := &rediscache.Config{
    RedisHost:     "localhost",
    RedisPort:     6379,
    RedisUsername: "your-redis-username",
    RedisPassword: "your-redis-password",
}

cache, err := rediscache.New(cfg)
if err != nil {
    // handle error
}
```

#### Cache Operations
##### Set

Store a key-value pair in the cache:

```go
err := cache.Set(ctx, key, value, expiration)
if err != nil {
    // handle error
}
```

##### Get

Retrieve a value from the cache using a key:

```go
value, err := cache.Get(ctx, key)
if err != nil {
    // handle error
}
```

##### Del

Delete a key from the cache:

```go
err := cache.Del(ctx, key)
if err != nil {
    // handle error
}
```

##### DelWildCard

Delete all keys that match a wildcard pattern:

```go
err := cache.DelWildCard(ctx, wildcard)
if err != nil {
    // handle error
}
```

##### Ping

Check if the cache is available:

```go
err := cache.Ping(ctx)
if err != nil {
    // handle error
}
```

##### HashOject

Generate a hash for a given object:

```go
hash := cache.HashOject(obj)
```

##### Hash

Generate a hash for a given key:

```go
hash := cache.Hash(key)
```

For more information, refer to the [Redis Go client documentation](https://pkg.go.dev/github.com/redis/go-redis/v9).

