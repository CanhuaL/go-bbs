package redis

import "time"

// redis key

// redis key注意使用命名空间的方式,方便查询和拆分
// redis中存有帖子的时间，分数以及对应帖子用户的投票状态
const (
	MaxAttempts        = 10             //  最大尝试次数
	Expiration         = 24 * time.Hour // 过期时间为一天
	Prefix             = "go_bbs:"      // 项目key前缀
	KeyPostTimeZSet    = "post:time"    // zset;贴子及发帖时间
	KeyPostScoreZSet   = "post:score"   // zset;贴子及投票的分数
	KeyPostVotedZSetPF = "post:voted:"  // zset;记录用户及投票类型;参数是post id，记录帖子是0，1还是-1的状态
	KeyAttempts        = "attempts:"

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
