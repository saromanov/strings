package strings

type KV struct {
	Key   string
	Value int
}

type Kvs []KV

func (a Kvs) Len() int           { return len(a) }
func (a Kvs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Kvs) Less(i, j int) bool { return a[i].Value < a[j].Value }
