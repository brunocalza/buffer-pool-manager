package bpm

import (
	"errors"
	"sort"
)

//DiskManagerMock is a memory mock for disk manager
type DiskManagerMock struct {
	numPage     int // tracks the number of pages. -1 indicates that there is no page, and the next to be allocates is 0
	pages       map[PageID]*Page
	maxNumPages int
}

//ReadPage reads a page from pages
func (d *DiskManagerMock) ReadPage(pageID PageID) (*Page, error) {
	if page, ok := d.pages[pageID]; ok {
		return page, nil
	}

	return nil, errors.New("Page not found")
}

//WritePage writes a page in memory to pages
func (d *DiskManagerMock) WritePage(page *Page) error {
	d.pages[page.id] = page
	return nil
}

//AllocatePage allocates one more page
func (d *DiskManagerMock) AllocatePage() *PageID {
	if d.numPage == d.maxNumPages-1 {
		return nil
	}
	d.numPage = d.numPage + 1
	pageID := PageID(d.numPage)
	return &pageID
}

//DeallocatePage removes page from disk
func (d *DiskManagerMock) DeallocatePage(pageID PageID) {
	delete(d.pages, pageID)
}

//PagesInDisk returns the pages id that are in disk in order
func (d *DiskManagerMock) PagesInDisk() []int {
	keys := make([]int, 0)
	for k := range d.pages {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	return keys
}

//MaxNumPages returns the maximun number of pages allowed in disk
func (d *DiskManagerMock) MaxNumPages() int {
	return d.maxNumPages
}

//NewDiskManagerMock returns a in-memory mock of disk manager
func NewDiskManagerMock(maxNumPages int) *DiskManagerMock {
	return &DiskManagerMock{-1, make(map[PageID]*Page), maxNumPages}
}
