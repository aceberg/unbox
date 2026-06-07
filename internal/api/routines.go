package api

import "time"

func testCurrentProxy() {

	for {
		if currentProxy != "" {
			ok := testOneProxy(currentProxy, colorMain+"[MAIN]"+colorReset)
			if !ok {
				alive = false
			}
		}

		time.Sleep(time.Duration(5) * time.Second)
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

		time.Sleep(time.Duration(30) * time.Second)
	}
}

func testAllTagsRoutine() {
	for {
		testAllTags()

		aliveTags = getAliveTags()

		time.Sleep(time.Duration(5*60) * time.Second)
	}
}

func testAllTags() {
	for _, tag := range getAllTags() {

		if tag == currentProxy {
			continue
		}

		testOneProxy(tag, "[ALL] ")
	}
}
