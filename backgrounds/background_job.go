package backgrounds

import (
	"credit/services"
	"credit/services/interfaces"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func StartBackground(service *services.Service) {
	var wg sync.WaitGroup

	errCh := make(chan error, 1)

	quit := make(chan struct{})

	go BatchUpdateOverdueInstallmentLine(service.DebtorService, &wg, errCh, quit)

	go func() {
		for err := range errCh {
			if err != nil {
				log.Printf("Error occured during batch update installment line: %v\n", err)
			}
		}
	}()

	go func() {
		wg.Wait()
		log.Println("All background jobs completed.")
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	log.Println("Shutting down background job service...")
	close(quit)
}

func BatchUpdateOverdueInstallmentLine(debtorService interfaces.DebtorInterface, wg *sync.WaitGroup, errCh chan<- error, quit <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Minute) // Set default timer to 1 minute

	for {
		select {
		case <-ticker.C:
			wg.Add(1)

			go debtorService.BatchUpdateOverdueInstallmentLine(wg, errCh)
		case <-quit:
			ticker.Stop()
			log.Println("Batch update overdue installment line job stopped.")
			return
		}
	}
}
