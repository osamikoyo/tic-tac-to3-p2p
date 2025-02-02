package app

import (
	"context"
	"github.com/osamikoyo/tic-tac-toe-p2p/internal/host"
	"github.com/osamikoyo/tic-tac-toe-p2p/pkg/loger"
)

type App struct {
	logger loger.Logger
	ChatHost *host.Host
}

func Init(port int) (App, error) {
	hostname, err := host.NewChatHost(context.Background(), port)
	if err != nil{
		return App{}, err
	}
	defer hostname.Close()

	return App{
		logger: loger.New(),
		ChatHost: hostname,
	}, nil
}

func (a App) Run() error {

}