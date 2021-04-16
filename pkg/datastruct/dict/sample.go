package dict

import (
	"math/rand"
	"time"

	"github.com/bazingame/gogoredis/pkg/macro"
)

/*
Use golang's data struct 'map' to implement dict.
*/

type SampleDict struct {
	m   map[string]interface{}
	len int
}

var _ Dict = NewSimpleDict()

func NewSimpleDict() *SampleDict {
	return &SampleDict{m: make(map[string]interface{})}
}

func (s *SampleDict) DictAdd(key string, val *interface{}) int {
	// if already exists return err
	if _, ok := s.m[key]; ok {
		return macro.DictErr
	} else {
		s.m[key] = val
		s.len++
		return macro.DictOk
	}
}

func (s *SampleDict) DictReplace(key string, val *interface{}) int {
	if s.DictAdd(key, val) == macro.DictOk {
		return macro.DictAddNew
	}
	s.m[key] = val
	return macro.DictAddReplace
}

func (s *SampleDict) DictFetchValue(key string) *interface{} {
	if v, ok := s.m[key]; ok {
		return &v
	} else {
		return nil
	}
}

func (s *SampleDict) DictLen() int {
	return s.len
}

func (s *SampleDict) DictGetRandomKey() (res string) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(s.len)

	for k, _ := range s.m {
		if random == 0 {
			res = k
		}
		random--
	}
	return res
}

func (s *SampleDict) DictDelete(key string) int {
	if _, ok := s.m[key]; ok {
		delete(s.m, key)
		return macro.DictOk
	} else {
		return macro.DictErr
	}
}
