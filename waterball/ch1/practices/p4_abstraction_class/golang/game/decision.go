package game

// Decision 定義列舉型別
type Decision string

// 定義列舉常量
const (
	Paper    Decision = "布"
	Scissors Decision = "剪刀"
	Stone    Decision = "石頭"
)

// String 方法實現 toString 功能
func (d Decision) String() string {
	return string(d)
}
