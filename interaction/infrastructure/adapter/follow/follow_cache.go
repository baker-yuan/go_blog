// Package adapter 端口-适配器模式的 adapter
package adapter

// FollowerZSetPrefix 关注者缓存，缓存关注了哪些 uin
const FollowerZSetPrefix = "Running::Public::follower_zset_"

// FollowCache Follow 的 Cache
type FollowCache struct {
	// RunningRedis redis.RunningRedis
}

//
// // Save 保存缓存
// func (c *FollowCache) Save(ctx context.Context, followerUIN, followeeUIN int64) (err error) {
// 	// 得到 redisHelper
// 	redisHelper := c.RunningRedis.GenHelper()
//
// 	// 得到 key
// 	key := FollowerZSetPrefix + strconv.FormatInt(followerUIN, 10)
//
// 	// ZAdd 到 Redis
// 	_, err = redisHelper.ZAdd(ctx, key, followeeUIN, time.Now().UnixNano()/1e6)
// 	if err != nil {
// 		log.Error(ctx, "application - service - ZAdd err: %+v", err)
// 		return err
// 	}
// 	return nil
// }
//
// // Delete 删除缓存
// func (c *FollowCache) Delete(ctx context.Context, followerUIN, followeeUIN int64) (err error) {
// 	// 得到 redisHelper
// 	redisHelper := c.RunningRedis.GenHelper()
//
// 	// 得到 key
// 	key := FollowerZSetPrefix + strconv.FormatInt(followerUIN, 10)
//
// 	// ZAdd 到 Redis
// 	_, err = redisHelper.ZRem(ctx, key, followeeUIN)
// 	if err != nil {
// 		log.Error(ctx, "application - service - ZRem err: %+v", err)
// 		return err
// 	}
// 	return nil
// }
