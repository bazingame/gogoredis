package impl

const (
	SDSMaxPrealloc = 1024 * 1024 // 最大预分配长度 1M
)

type SDS struct {
	free int
	len  int
	buf  []byte
}

// SDSNew Create a new sds string starting from a null terminated string
func SDSNew(initStr string) *SDS {
	var strLen int
	strLen = len(initStr)
	return sdsNewLen(&initStr, strLen)
}

// SDSEmpty Create an empty (zero length) sds string
func SDSEmpty() *SDS {
	return sdsNewLen(nil, 0)
}

func (s *SDS) Len() int {
	return s.len
}
func (s *SDS) Avail() int {
	return s.free
}
func (s *SDS) Buf() []byte {
	return s.buf[:s.len]
}

// Clear  Modify an sds string on-place to make it empty (
func (s *SDS) Clear() *SDS {
	s.free += s.len
	s.len = 0
	return s
}

// Cat Append the specified null terminated string to the sds string 's'
func (s *SDS) Cat(str string) *SDS {
	return s.sdsCatLen(&str, len(str))
}

//func (s *SDS) Free() *SDS {
//	s.free += s.len
//	s.len = 0
//	return s
//}

func (s *SDS) Cpy(str string) *SDS {
	return s.sdsCpyLen(&str, len(str))
}

func (s *SDS) Range(start, end int) *SDS {
	s.buf = s.buf[start:end]
	s.len = end - start + 1
	return s
}

// Trim Remove the part of the string from left and from right
func (s *SDS) Trim() *SDS {
	panic("implement me")
	return s
}

// Cmp Compare two sds strings s1 and s2
func (s *SDS) Cmp() *SDS {
	panic("implement me")
	return s
}

// SDSNewLen  Create a new sds string with the content specified by the 'initStr' pointer and 'initLen'.
func sdsNewLen(initStr *string, initLen int) *SDS {
	var sh = new(SDS) // todo optimize memory apply
	sh.len = initLen
	sh.free = 0
	if initLen != 0 && initStr != nil {
		sh.buf = []byte(*initStr)
	}
	return sh
}

func (s *SDS) sdsCatLen(str *string, len int) *SDS {
	// 计算是否申请新内存, 小于的话直接追加

	s.sdsMakeRoomFor(len)

	copy(s.buf[s.len:], *str)
	s.len = s.len + len

	s.free = s.free - len
	return s
}

func (s *SDS) sdsCpyLen(str *string, len int) *SDS {
	s.sdsMakeRoomFor(len - s.len)

	copy(s.buf, *str)
	s.free = s.len + s.free - len
	s.len = len

	return s
}

func (s *SDS) sdsMakeRoomFor(addLen int) {
	if addLen <= s.free {
		return
	}

	var newLen = s.len + addLen
	if newLen < SDSMaxPrealloc {
		newLen *= 2
	} else {
		newLen += SDSMaxPrealloc
	}
	oldBuf := s.buf
	s.buf = make([]byte, newLen)
	copy(s.buf, oldBuf)
	s.free = newLen - s.len
	return
}
