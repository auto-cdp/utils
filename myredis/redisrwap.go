package redis

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"utils/log"
)

type DataFunc func(string)

func NewRedisPool(host string, MaxIdle int, MaxActive int, idletimeout time.Duration) redis.Pool {
	return redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: idletimeout * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("redis connection error: %s", err)
			}
			return nil
		},
	}
}

type RedisConn struct {
	pool redis.Pool
}

func NewRedisConn(pool redis.Pool) RedisConn {
	return RedisConn{pool: pool}
}

func (rp *RedisConn) Publish(channel, message string) (int, error) {
	conn := rp.pool.Get()
	defer conn.Close()

	n, err := redis.Int(conn.Do("PUBLISH", channel, message))

	return n, err
}

func (rp *RedisConn) Subscribe(channel string) (redis.PubSubConn, error) {
	conn := rp.pool.Get()
	psc := redis.PubSubConn{conn}
	if err := psc.Subscribe(channel); err != nil {
		log.Slogger.Errorf("[Redis] 订阅通道[%s]失败.[%s]", channel, err)
		return psc, err
	}
	log.Slogger.Debugf("[Redis] 订阅通道[%s]成功.", channel)
	return psc, nil
}

// 处理普通消息
func (rp *RedisConn) HandleCommMessage(psc redis.PubSubConn, handleMessage DataFunc) {
	done := make(chan error, 1)
	go func() {
		defer psc.Close()
		for {
			switch msg := psc.Receive().(type) {
			case redis.Message:
				log.Slogger.Debugf(fmt.Sprintf("[Redis] 收到通道[%s]信息:%s", msg.Channel, msg.Data))
				log.Slogger.Info("[Redis] 启动goroutine处理.")
				go handleMessage(string(msg.Data))
			case redis.Subscription:
				if msg.Count == 0 {
					done <- nil
					return
				}
			case error:
				done <- fmt.Errorf("[Redis] 接收信息发生错误:%v", msg)
				log.Slogger.Error(msg)
				return
			}
		}
	}()

	ctx := context.TODO()
	err := healthchek(ctx, psc, done)
	if err != nil {
		log.Slogger.Error(err)
	}
}

// 处理命令结果消息
func (rp *RedisConn) HandleCMDResultMessage(psc redis.PubSubConn, result chan string) {
	done := make(chan error, 1)
	go func() {
		defer psc.Close()
		for {
			switch msg := psc.Receive().(type) {
			case redis.Message:
				log.Slogger.Debugf(fmt.Sprintf("[Redis] 收到通道[%s]信息:%s", msg.Channel, msg.Data))
				result <- string(msg.Data)
			case redis.Subscription:
				if msg.Count == 0 {
					done <- nil
					return
				}
			case error:
				done <- fmt.Errorf("[Redis] 接收信息发生错误:%v", msg)
				log.Slogger.Error(msg)
				return
			}
		}
	}()

	ctx := context.TODO()
	err := healthchek(ctx, psc, done)
	if err != nil {
		log.Slogger.Error(err)
	}
}

func healthchek(ctx context.Context, psc redis.PubSubConn, done chan error) error {
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			if err := psc.Unsubscribe(); err != nil {
				return fmt.Errorf("[Redis] 取消订阅失败:%v", err)
			}
			return nil
		case err := <-done:
			return err
			/*case <-tick.C:
			if err := psc.Ping("ping"); err != nil {
				return nil
			}*/

		}
	}
}
