package fungolog

type Level uint8

const (
	Debug Level = iota
	Info
	Warning
	Error
	Panic
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Panic:
		return "Panic"
	default:
		return "UNKNOWN"
	}
}
