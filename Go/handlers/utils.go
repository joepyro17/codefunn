package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
)

func GetIP(ctx *fiber.Ctx) error {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	var currentIP string

	for _, addr := range addresses {

		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				currentIP = ipnet.IP.String()
			}
		}
	}

	return ctx.SendString("IP address = " + currentIP)
}