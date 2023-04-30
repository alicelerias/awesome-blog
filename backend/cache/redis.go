package cache

import (
	"encoding/json"
	"time"
)

// SetKey
// value must be json serializable

func (s *RedisClient) genKey(name string, id string) string {
	return name + "_" + id
}

func (s *RedisClient) SetKey(name string, id string, value interface{}, expiration time.Duration) error {
	key := s.genKey(name, id)
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.client.Set(key, bytes, expiration).Err()
}

func (s *RedisClient) GetKey(name string, id string, model interface{}) error {
	key := s.genKey(name, id)
	result, err := s.client.Get(key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), model)
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisClient) DeleteKey(name string, id string) error {
	key := s.genKey(name, id)
	err := s.client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
