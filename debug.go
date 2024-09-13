package debugger

type Debug bool

func InitDebug() Debug {
	return Debug(false)
}

func NewDebug() Debug {
	return InitDebug()
}

func (d *Debug) IsDebug() bool {
	return bool(*d)
}

func (d *Debug) SetDebug() {
	*d = Debug(true)
}

func (d *Debug) UnsetDebug() {
	*d = Debug(false)
}

func (d *Debug) ToggleDebug() {
	*d = Debug(!bool(*d))
}
