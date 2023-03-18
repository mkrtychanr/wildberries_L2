package pattern

import "fmt"

type Server interface {
	SetConfig()
	ConnectToDB()
	ConnectToEventBroker()
	SetHttpRouting()
	StartRouting()
}

type ApiServer struct{}

func (s *ApiServer) SetConfig() {
	fmt.Println("The config is set")
}

func (s *ApiServer) ConnectToDB() {
	fmt.Println("Connected to database")
}

func (s *ApiServer) ConnectToEventBroker() {
	fmt.Println("Connected to event broker")
}

func (s *ApiServer) SetHttpRouting() {
	fmt.Println("Http router is set")
}

func (s *ApiServer) StartRouting() {
	fmt.Println("Routing is started")
}

func (s *ApiServer) StartServer() {
	s.SetConfig()
	s.ConnectToDB()
	s.ConnectToEventBroker()
	s.SetHttpRouting()
	s.StartRouting()
	fmt.Println("Server is up")
}

// Паттерн строитель относится к порождающим паттернам уровня объекта.

// Плсюы: позволяет создавать продукты поэтапно, позволяет переиспользовать код для других продуктов
// изолирует код сборки от бизнес логики

// Минусы: усложняет код из-за введения дополнительных структур
