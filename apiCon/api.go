package apiConn

import (
	"context"
	"golang.org/x/time/rate"
	"sort"
	"time"
)

type APIConn struct {
	apiLimit,
	dbLimit RateLimiter
}

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}

	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

func Open() *APIConn {
	return &APIConn{
		apiLimit: MultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 1),
			rate.NewLimiter(Per(5, time.Minute), 5),
		),
		dbLimit: MultiLimiter(
			rate.NewLimiter(rate.Every(time.Second*5), 1),
		),
	}
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

//func Open() *APIConn {
//	return &APIConn{
//		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 1),
//	}
//}
//
//func (c *APIConn) Read(ctx context.Context) (string, error) {
//	if err := c.rateLimiter.Wait(ctx); err != nil {
//		return "", err
//	}
//
//	return "Read", nil
//}
//
//func (c *APIConn) Resolve(ctx context.Context) error {
//	if err := c.rateLimiter.Wait(ctx); err != nil {
//		return err
//	}
//
//	return nil
//}
