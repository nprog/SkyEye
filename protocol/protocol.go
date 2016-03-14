package protocol

type CmdInterface interface {
	GetIndicates() uint16
	GetContent() interface{}
	GetReport() interface{}
	SetIndicates(indicates uint16)
	SetContent(content interface{})
	SetReport(report interface{})
}

type Cmd struct {
	indicates uint32      `json:"i"`
	content   interface{} `json:"c"`
	report    interface{} `json:"report"`
}

func NewCmd(indicates uint32) *Cmd {
	return &Cmd{
		indicates: indicates,
	}
}

func (c *Cmd) GetIndicates() uint32 {
	return c.indicates
}

func (c *Cmd) GetContent() interface{} {
	return c.content
}

func (c *Cmd) GetReport() interface{} {
	return c.report
}

func (c *Cmd) SetIndicates(indicates uint32) {
	c.indicates = indicates
}

func (c *Cmd) SetContent(content interface{}) {
	c.content = content
}

func (c *Cmd) SetReport(report interface{}) {
	c.report = report
}

type Ping struct {
	Time uint64 `json:"time"`
}
