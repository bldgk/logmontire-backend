package logging

func Init() {
	RegisterConsoleLogger()
	RegisterExternalLogger()
}
