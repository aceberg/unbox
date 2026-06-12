package api

import (
	"strconv"
	"time"
)

func testCurrentProxy() {

	for {
		if currentProxy != "" {
			ok := testOneProxy(currentProxy, colorMain+"[MAIN]"+colorReset)
			if !ok {
				alive = false
			}
		}

		time.Sleep(time.Duration(Config.DelayMain) * time.Second)
	}
}

func updateAliveTags() {

	for {
		for i, tag := range aliveTags {

			if tag.Tag == currentProxy {
				continue
			}

			if i > 3 {
				break
			}

			testOneProxy(tag.Tag, colorBkp+"[BKP] "+colorReset)
		}

		aliveTags = getAliveTags()

		time.Sleep(time.Duration(Config.DelayBkp) * time.Second)
	}
}

func testAllTagsRoutine() {
	for {
		testAllTags()

		aliveTags = getAliveTags()

		time.Sleep(time.Duration(Config.DelayAll) * time.Second)
	}
}

func testAllTags() {

	tags := getAllTags()
	l := len(tags)

	for i, tag := range tags {

		if tag == currentProxy {
			continue
		}

		testOneProxy(tag, "["+strconv.Itoa(i+1)+"-"+strconv.Itoa(l)+"]")
	}
}
