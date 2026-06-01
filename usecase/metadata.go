package usecase

import (
	"math/rand"
	"strings"
	"time"
)

func generateInitials(name string) string {
	words := strings.Fields(name)
	var initials string
	for i, word := range words {
		if i >= 2 {
			break
		}
		if len(word) > 0 {
			initials += string(word[0])
		}
	}
	return strings.ToUpper(initials)
}

func generateHandle(name string) string {
	cleanName := strings.ToLower(name)
	cleanName = strings.ReplaceAll(cleanName, " ", "")
	if len(cleanName) > 12 {
		cleanName = cleanName[:12]
	}
	return "@" + cleanName
}

func generateAvatarGradient() string {
	gradients := []string{
		"linear-gradient(135deg, #7c3aed, #2563eb)",
		"linear-gradient(135deg, #0ea5e9, #2563eb)",
		"linear-gradient(135deg, #f43f5e, #e11d48)",
		"linear-gradient(135deg, #10b981, #059669)",
		"linear-gradient(135deg, #a78bfa, #7c3aed)",
		"linear-gradient(135deg, #f59e0b, #d97706)",
	}
	rand.Seed(time.Now().UnixNano())
	return gradients[rand.Intn(len(gradients))]
}
