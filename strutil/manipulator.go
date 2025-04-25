package strutil

type Manipulator struct {
	data []byte
	err  error
}

func NewManipulator(data string) *Manipulator {
	return &Manipulator{
		data: []byte(data),
	}
}

func (m *Manipulator) GetString() (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return string(m.data), nil
}
