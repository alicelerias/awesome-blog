package cache

import (
	"encoding/json"
	"time"
)

// SetKey
// value must be json serializable

func (s *RedisClient) genKey(name string, nameSpace string) string {
	return name + "_" + nameSpace
}

func (s *RedisClient) SetKey(name string, nameSpace string, value interface{}, expiration time.Duration) error {
	key := s.genKey(name, nameSpace)
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.client.Set(key, bytes, expiration).Err()
}

func (s *RedisClient) GetKey(name string, nameSpace string, model interface{}) error {
	key := s.genKey(name, nameSpace)
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

func (s *RedisClient) DeleteKey(name string, nameSpace string) error {
	key := s.genKey(name, nameSpace)
	err := s.client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
