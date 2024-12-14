package bootstrap

import (
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"sync"
	"syscall"
	"time"
)

func HeapProfile(wg *sync.WaitGroup) {
	defer wg.Done()

	// Membuat file untuk profil heap
	f, err := os.Create("heap_profile.prof")
	if err != nil {
		log.Fatalf("Error creating heap profile file: %v", err)
	}
	defer func() {
		f.Close()
		log.Println("Heap profile file closed gracefully.")
	}()

	// Channel untuk menangkap sinyal OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	log.Println("Heap profiling started. Waiting for shutdown signal...")

	select {
	case <-quit:
		log.Println("Shutdown signal received. Generating heap profile...")
		// Ambil profil heap sebelum keluar
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Printf("Error writing heap profile: %v", err)
		} else {
			log.Println("Heap profile successfully written.")
		}
	case <-time.After(10 * time.Second): // Simulasi waktu jika diperlukan
		log.Println("Heap profiling completed without interruption.")
	}
}
