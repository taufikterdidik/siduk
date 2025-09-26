package utils

import "fmt"

func SendPINWhatsApp(noWhatsapp, pin string) error {
	// Integrasi WhatsApp API asli di sini.
	fmt.Printf("[WA] Send PIN %s ke %s\n", pin, noWhatsapp)
	return nil
}