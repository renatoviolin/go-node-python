package ports

type PortHttp interface {
	SetupAndRun(port string)
}
