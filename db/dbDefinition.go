package db

type Operations interface {
	Select(dbIndexStr string) (int, error)
	Set(dbIndex int, key string, value interface{})
	Get(dbIndex int, key string) interface{}
	Del(dbIndex int, key string) interface{}
	Show(dbIndex int) <-chan string
}