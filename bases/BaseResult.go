package bases

type IResult interface {
	SetData(interface{})
	Init()
}

type BaseResult struct {
	LogId  string
	ErrNo  string
	ErrMsg string
	Data   interface{}
}

func (this *BaseResult) SetData(data interface{}) {
	this.Data = data
}

func (this *BaseResult) Init() {
}
