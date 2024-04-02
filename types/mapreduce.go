package types

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ByKey []*KeyValue

func (bk ByKey) Len() int {
	return len(bk)
}

func (bk ByKey) Less(i, j int) bool {
	return bk[i].Key < bk[j].Key
}
func (bk ByKey) Swap(i, j int) {
	bk[i], bk[j] = bk[j], bk[i]
}

type State int

const (
	Map State = iota + 1
	Reduce
	Exit
	Wait
)

type MasterTaskStatus int

const (
	Idle      MasterTaskStatus = iota + 1 // 未开始
	Inprocess                             //进行中
	Completed                             //完成
)

type MapReduceTask struct {
	Input         string   `json:"input"`
	TaskState     State    `json:"task_state"`
	NReducer      int      `json:"n_reducer"`     // reducer数量
	TaskNumber    int      `json:"task_number"`   //任务数量
	Intermediates []string `json:"intermediates"` // map之后文件存储的地址
	Output        string   `json:"output"`        // 输出的地址
}

// Tokenization 分词返回结构
type Tokenization struct {
	Token string //词条
	DocId int64
}
