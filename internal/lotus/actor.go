package lotus

import "github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"

func CheckForActorAddress(service services.LotusService) (string, error) {
	return service.GetMinerAddress()
}