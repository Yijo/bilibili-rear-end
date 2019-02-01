package handlers

type success int

const (
	none success = iota
)

func (this success) Message() string{
	switch this {
	}
	return "成功"
}

func (this success) DisplayMsg() string{
	switch this {
	}

	return "成功"
}