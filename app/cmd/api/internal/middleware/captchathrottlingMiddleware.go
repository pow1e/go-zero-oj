package middleware

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/pkg/utils"
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"strconv"
	"time"
)

type CaptchaThrottlingMiddleware struct {
	maxTime int
	redis   *redis.Redis
}

func NewCaptchaThrottlingMiddleware(redis *redis.Redis, maxTime int) *CaptchaThrottlingMiddleware {
	return &CaptchaThrottlingMiddleware{maxTime, redis}
}

func (m *CaptchaThrottlingMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := utils.GetIP(r)
		err := getCurrentIpCount(r.Context(), ip, m.redis, m.maxTime)
		if err != nil {
			if err.Error() == consts.ErrOverTakeCaptchaMaxTimes {
				response.JsonBaseResponseCtx(r.Context(), w, err)
				return
			}
			response.JsonBaseResponseCtx(r.Context(), w, errors.New(consts.ErrUnKnow))
		}
		next(w, r)
	}
}

// 从redis中获取当前ip的次数
func getCurrentIpCount(ctx context.Context, clientIP string, client *redis.Redis, maxTime int) error {
	key := consts.GetCurrentIPCaptchaMaxTimeCachePrefix + clientIP
	// 判断上一次请求和请求计数器
	count, err := client.GetCtx(ctx, key)
	if err != nil && err != redis.Nil {
		logx.Debug("从缓存中获取当前ip的请求次数失败", err)
		return err
	}

	// 如果没有上次请求记录，则保存到redis中
	if count == "" {
		// 初始化设置当前ip的请求次数为1,过期时间为一天
		expireSeconds := int64(time.Hour * 24 / time.Second)
		_, err = client.SetnxExCtx(ctx, key, "1", int(expireSeconds))
		if err != nil {
			logx.Debug("设置当前用户ip缓存失败", err)
			return err
		}
		// 讲count设置为1
		count = "1"
	}

	// 判断是否超过限制
	times, err := strconv.Atoi(count)
	if err != nil {
		logx.Debug("转换类型失败", err)
		return err
	}
	if times >= maxTime {
		return errors.New(consts.ErrOverTakeCaptchaMaxTimes)
	}

	// 如果没有超过限制则对当前ip的访问的值+1
	_, err = client.IncrbyCtx(ctx, key, 1)
	if err != nil {
		return err
	}

	return nil
}
