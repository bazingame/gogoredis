package dict

//type DictHt struct {
//	table    []*DictEntry
//	size     int64
//	sizeMark int64 // used to calculate index, always equal to size-1
//	used     int64
//}
//
//type DictEntry struct {
//	key   *interface{}
//	value *interface{}
//	next  *DictEntry
//}
//
//type Dict struct {
//	dictHt     [2]DictHt
//	tRehashIdx int
//}

type Dict interface {
	DictAdd(key string, val *interface{}) int
	DictReplace(key string, val *interface{}) int
	DictFetchValue(key string) *interface{}
	DictGetRandomKey() string
	DictDelete(key string) int
	DictLen() int
	//DictRelease()
}
