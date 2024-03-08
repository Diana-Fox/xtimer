package utils

type Logger interface {
	DEBUG(msg string, arg ...any)
	INFO(msg string, arg ...any)
	WARNING(msg string, arg ...any)
	ERROR(msg string, arg ...any)
	CRITICAL(msg string, arg ...any)
}
