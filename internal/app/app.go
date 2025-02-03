package app

import (
	"bufio"
	"context"
	"fmt"
	"github.com/osamikoyo/tic-tac-toe-p2p/internal/host"
	"github.com/osamikoyo/tic-tac-toe-p2p/pkg/loger"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Commands:")
	fmt.Println("/connect <peer-address> - Connect to a peer")
	fmt.Println("/quit - Exit the chat")
	fmt.Print("> ")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "/quit" {
			break
		}

		if input == "/get field" {

		}
	}
}