package Bitcask

var HM *HashMap

type HashMapEntry struct {
	key    []byte
	offset uint64
}

type HashMap struct {
	Entries []HashMapEntry
}

func NewHashMap() *HashMap {
	entries := make([]HashMapEntry, MemSize)
	return &HashMap{
		Entries: entries,
	}
}

func init() {
	HM = NewHashMap()
}
