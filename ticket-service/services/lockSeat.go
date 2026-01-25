package services

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type SeatLockService struct {
	RedisClient *redis.Client
}

func NewSeatLockService(redisClient *redis.Client) *SeatLockService {
	return &SeatLockService{
		RedisClient: redisClient,
	}
}

func (s *SeatLockService) LockSeat(ctx context.Context, seatID string, userID string) error {
	lockKey := "lock:seat:" + seatID
	userCartKey := "user:" + userID + ":seats"
	ttl := 10 * time.Minute

	success, err := s.RedisClient.SetNX(ctx, lockKey, userID, ttl).Result()
	if err != nil {
		return err
	}
	if !success {
		existingOwner, _ := s.RedisClient.Get(ctx, lockKey).Result()
		if existingOwner != userID {
			return errors.New("seat_already_locked")
		}
	}

	err = s.RedisClient.SAdd(ctx, userCartKey, seatID).Err()
	if err != nil {
		return err
	}

	allSeats, err := s.RedisClient.SMembers(ctx, userCartKey).Result()
	if err != nil {
		return err
	}

	pipe := s.RedisClient.Pipeline()
	pipe.Expire(ctx, userCartKey, ttl)
	for _, sID := range allSeats {
		seatKey := "lock:seat:" + sID
		pipe.Expire(ctx, seatKey, ttl)
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
