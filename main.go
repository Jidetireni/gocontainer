package main

import (
	"log"

	"github.com/Jidetireni/gocontainer/internal/container"
)

func main() {
	log.Println("🔧 Setting up root filesystem...")
	if err := container.SetUpRootFS(); err != nil {
		log.Fatalf("❌ Setup failed: %v", err)
	}

	log.Println("✅ Root filesystem is ready.")
	if err := container.ChrootIntoRootFS(); err != nil {
		log.Fatalf("❌ Chroot failed: %v", err)
	}

	log.Println("👋 Exiting chroot.")

}
