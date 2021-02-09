package bpm

import "sort"

// ClockReplacerRepresentation represents the clock replacer for serialization
type ClockReplacerRepresentation struct {
	ClockHand int
	Clock     []ClockValue
}

// ClockValue represents the clock node
type ClockValue struct {
	ClockFrame     int
	ReferenceValue bool
}

//Response is the JSON response of the web server
type Response struct {
	PagesInDisk     []int
	MaxPoolSize     int
	PagesTable      map[PageID]FrameID
	ClockReplacer   ClockReplacerRepresentation
	MaxDiskNumPages int
	PinCount        map[int]int
}

func getClockReplacerRepresentation(clockReplacer *ClockReplacer) ClockReplacerRepresentation {
	clockValues := []ClockValue{}
	var clockHand int
	ptr := clockReplacer.cList.head
	for i := 0; i < clockReplacer.Size(); i++ {
		clockValues = append(clockValues, ClockValue{int(ptr.key.(FrameID)), ptr.value.(bool)})
		if *clockReplacer.clockHand == ptr {
			clockHand = i
		}

		ptr = ptr.next
	}

	return ClockReplacerRepresentation{clockHand, clockValues}
}

func pagesInDisk(diskManager DiskManager) []int {
	keys := make([]int, 0)
	for k := range diskManager.(*DiskManagerMock).pages {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	return keys
}

//NewResponse creates a response for web server given a buffer pool manager
func NewResponse(bufferPool *BufferPoolManager) Response {
	pagePinCount := make(map[int]int)
	for i := 0; i < len(bufferPool.pages); i++ {
		if bufferPool.pages[i] != nil {
			pagePinCount[int(bufferPool.pages[i].ID())] = bufferPool.pages[i].PinCount()
		}
	}

	return Response{
		pagesInDisk(bufferPool.diskManager),
		MaxPoolSize,
		bufferPool.pageTable,
		getClockReplacerRepresentation(bufferPool.replacer),
		DiskMaxNumPages,
		pagePinCount,
	}
}
