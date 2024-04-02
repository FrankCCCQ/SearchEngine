package redis

import (
	"context"

	"github.com/RoaringBitmap/roaring"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// PushInvertedPath 把存放db的path信息放到redis中
func PushInvertedPath(ctx context.Context, key string, paths []string) error {
	for _, v := range paths {
		err := RedisClient.LPush(ctx, key, v).Err()
		if err != nil {
			return errors.Wrap(err, "failed to push inverted path in redis")
		}
	}
	return nil
}

// SetInvertedPath 把存放db的path信息放到redis中
func SetInvertedPath(ctx context.Context, key string, path string) error {
	err := RedisClient.Set(ctx, key, path, redis.KeepTTL).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set inverted path in redis")
	}
	return nil
}

// GetInvertedPath 获取存储的path信息
func GetInvertedPath(ctx context.Context, key string) (path string, err error) {
	path, err = RedisClient.Get(ctx, key).Result()
	if err != nil {
		return path, errors.Wrap(err, "failed to get inverted path")
	}
	return
}

// ListInvertedPath 把存放在redis中的信息放到path中，包括day，month和season
func ListInvertedPath(ctx context.Context, keys []string) (paths []string, err error) {
	for _, key := range keys {
		switch key {
		case InvertedIndexDbPathDayKey, TrieTreeDbPathDayKey:
			results := RedisClient.LRange(ctx, key, 0, -1)
			if err != nil {
				return paths, errors.Wrap(err, "failed to get value")
			}
			paths = append(paths, results.Val()...)
		case InvertedIndexDbPathMonthKey, TrieTreeDbPathMonthKey, InvertedIndexDbPathSeasonKey, TrieTreeDbPathSeasonKey:
			prefixKey := getAllDbPaths(key)
			results, errx := ListInvertedIndexByPrefixKey(ctx, prefixKey)
			if errx != nil {
				return paths, errors.Wrap(errx, "failed to list inverted index")
			}
			paths = append(paths, results...)
		default:
			return
		}
	}
	return
}

// DeleteInvertedIndexPath 删除inverted index path
func DeleteInvertedIndexPath(ctx context.Context, key string) (err error) {
	err = RedisClient.Del(ctx, key).Err()
	if err != nil {
		return errors.Wrap(err, "failed to delete inverted index path")
	}
	return
}

// BatchDelateInvertedIndexPath 批量删除
func BatchDelateInvertedIndexPath(ctx context.Context, keys []string) (err error) {
	for _, key := range keys {
		_ = DeleteInvertedIndexPath(ctx, key)
	}
	if err != nil {
		return errors.Wrap(err, "failed to batch delete inverted index path")
	}
	return
}

// SetInvertedIndexTokenDocIds 搜索缓存过的结果
func SetInvertedIndexTokenDocIds(ctx context.Context, token string, docIds *roaring.Bitmap) (err error) {
	docIdsByte, _ := docIds.MarshalBinary()
	err = RedisClient.Set(ctx, getQueryTokenDocsIdsKey(token), docIdsByte, QueryTokenDocIdsDefaultTimeout).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set inverted index token docIds")
	}
	return
}

// GetInvertedIndexTokenDocIds 获取缓存的结果
func GetInvertedIndexTokenDocIds(ctx context.Context, token string) (docIds *roaring.Bitmap, err error) {
	res, err := RedisClient.Get(ctx, getQueryTokenDocsIdsKey(token)).Result()
	if err != nil {
		return docIds, errors.Wrap(err, "failed to get query token docIds key from redis")
	}
	docIds = roaring.NewBitmap()
	err = docIds.UnmarshalBinary([]byte(res))
	if err != nil {
		return docIds, errors.Wrap(err, "failed to unmarshal binary")
	}
	return
}

// PushInvertedIndexToken 存储用户搜索的历史记录
func PushInvertedIndexToken(ctx context.Context, userId int64, token string) (err error) {
	err = RedisClient.LPush(ctx, getUserQueryTokenKey(userId)).Err()
	if err != nil {
		return errors.Wrap(err, "failed to push inverted index token")
	}
	return
}

// ListInvertedToken 获取用户搜索的历史记录
func ListInvertedToken(ctx context.Context, userId int64) (tokens []string, err error) {
	tokens, err = RedisClient.LRange(ctx, getUserQueryTokenKey(userId), 0, -1).Result()
	if err != nil {
		return tokens, errors.Wrap(err, "failed to list inverted index token")
	}
	return
}

// PushInvertedMonthPath 把存到db的path信息放到redis的month维度
func PushInvertedMonthPath(ctx context.Context, key string, paths []string) (err error) {
	for _, v := range paths {
		err = RedisClient.LPush(ctx, key, v).Err()
		if err != nil {
			return errors.Wrap(err, "failed to push invertd month path")
		}
	}
	return
}

// ListInvertedIndexByPrefixKey 通过前缀获取所有的valve index_platform:inverted_index:month:*
func ListInvertedIndexByPrefixKey(ctx context.Context, prefixKey string) (paths []string, err error) {
	iter := RedisClient.Scan(ctx, 0, prefixKey, 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		value, _ := RedisClient.Get(ctx, key).Result()
		paths = append(paths, value)
	}
	if err != nil {
		err = errors.Wrap(err, "failed to get value")
	}
	return
}

// ListAllPrefixKey 通过前缀获取所有的value
func ListAllPrefixKey(ctx context.Context, prefixKey string) (paths []string, err error) {
	iter := RedisClient.Scan(ctx, 0, prefixKey, 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		paths = append(paths, key)
	}
	if err != nil {
		err = errors.Wrap(err, "failed to get value")
	}
	return
}
