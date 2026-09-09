package conf

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/share"
)

// RemoveUnreachable removes unreachable nodes from sing-box config
func RemoveUnreachable() {
	start := time.Now()

	lenAll := checkAllTags()
	aliveServers := api.GetAliveServers()
	lenOnline := len(aliveServers)

	if lenOnline == 0 {
		log.Println("No proxies online. Exiting")
		return
	}

	fmt.Println()
	log.Println("INFO Scanned servers:", lenAll)
	log.Println("INFO Found online:", lenOnline)
	log.Println("INFO Scan finished in", time.Since(start).Round(time.Millisecond))

	if share.Settings.BestN != 0 && lenOnline > share.Settings.BestN {
		aliveServers = aliveServers[0:share.Settings.BestN]
	}

	editConfig(aliveServers)
}

func checkAllTags() int {

	tags := api.GetAllTags()
	l := len(tags)
	lStr := strconv.Itoa(l)

	type Job struct {
		i   int
		tag string
	}

	jobs := make(chan Job)
	var wg sync.WaitGroup

	for range 5 {
		wg.Go(func() {
			for job := range jobs {
				api.CheckOneProxy(job.tag, "["+strconv.Itoa(job.i+1)+"-"+lStr+"]")
			}
		})
	}

	for i, tag := range tags {
		jobs <- Job{i: i, tag: tag}
	}

	close(jobs)
	wg.Wait()

	return l
}
