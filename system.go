package rxclient

import "time"

type CommandBody struct {
	Command string
	Arg     string
	Args    []string
	Timeout int
}

type Response struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}

type StatusResp struct {
	Status       string    `json:"status,omitempty"`
	RunningSince time.Time `json:"runningSince,omitempty"`
	Uptime       string    `json:"uptime,omitempty"`
	PID          int       `json:"pid,omitempty"`
	Memory       string    `json:"memory,omitempty"`
	CPU          string    `json:"cpu,omitempty"`
	IsEnabled    bool      `json:"isEnabled"`
	IsActive     bool      `json:"isActive"`
	IsFailed     bool      `json:"isFailed"`
	RestartCount int       `json:"restartCount"`
}

func (c *rxClient) RunCommand(body *CommandBody) (*Response, error) {
	var cmdResp = &Response{}
	_, err := c.rx.Send("cmd/run", &body, body.Timeout+1, &cmdResp, "any")
	return cmdResp, err
}

func (c *rxClient) SystemdStatus(uint string, timeout int) (*StatusResp, error) {
	body := &CommandBody{
		Command: uint,
		Args:    nil,
		Timeout: timeout,
	}
	var cmdResp = &StatusResp{}
	_, err := c.rx.Send("cmd/systemctl/status", &body, body.Timeout+1, &cmdResp, "any")
	return cmdResp, err
}

type SystemCTLCommand string

const (
	SystemCTLDisable SystemCTLCommand = "disable"
	SystemCTLEnable  SystemCTLCommand = "enable"
	SystemCTLRestart SystemCTLCommand = "restart"
	SystemCTLStop    SystemCTLCommand = "start"
	SystemCTLStart   SystemCTLCommand = "start"
)

func (c *rxClient) SystemdCommand(uint string, commandType SystemCTLCommand, timeout int) (*Response, error) {
	body := &CommandBody{
		Command: uint,
		Arg:     string(commandType),
		Timeout: timeout,
	}
	var cmdResp = &Response{}
	_, err := c.rx.Send("cmd/systemctl/command", &body, body.Timeout+1, &cmdResp, "any")
	return cmdResp, err
}
