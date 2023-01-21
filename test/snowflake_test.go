package test

import (
	"sync"
	"testing"

	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"
)

func TestGenID(t *testing.T) {
	// use 10 goroutines to generate 1000IDs
	// save them to a slice
	// and check if it is unique
	idSlice := make([]int64, 1000)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			for j := 0; j < 100; j++ {
				idSlice[i*100+j] = utils.GenID()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// check if it is unique
	for i := 0; i < len(idSlice); i++ {
		for j := i + 1; j < len(idSlice); j++ {
			if idSlice[i] == idSlice[j] {
				t.Errorf("ID is not unique")
			}
		}
	}

}
