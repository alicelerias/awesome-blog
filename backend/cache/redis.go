package cache

import (
	"encoding/json"
	"time"
)

// SetKey
// value must be json serializable
func (s *RedisClient) SetKey(key string, id string, value interface{}, expiration time.Duration) error {
	// TODO uma conexao por server
	concat := key + "_" + id
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.client.Set(concat, bytes, expiration).Err()
}

func (s *RedisClient) GetKey(key string, id string, model interface{}) error {
	concat := key + "_" + id
	result, err := s.client.Get(concat).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), model)
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisClient) DeleteKey(key string, id string) error {
	concat := key + "_" + id
	err := s.client.Del(concat).Err()
	if err != nil {
		return err
	}
	return nil
}
