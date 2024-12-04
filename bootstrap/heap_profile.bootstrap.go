package bootstrap

import (
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func HeapProfile(wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := os.Create("heap_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Ambil profil heap
	pprof.WriteHeapProfile(f)
}
