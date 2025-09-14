package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	lg := log.New(
		os.Stdout,
		"[6spr] ", // префикс
		log.LstdFlags,
	)

	// Создаем сервер
	s := server.Srvr(lg)
	// Запуск сервера
	s.Serve()
}
