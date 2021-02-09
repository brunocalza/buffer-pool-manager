package bpm

import (
	"errors"
)

//DiskManagerMock is a memory mock for disk manager
type DiskManagerMock struct {
	numPage int // tracks the number of pages. -1 indicates that there is no page, and the next to be allocates is 0
	pages   map[PageID]*Page
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
	if d.numPage == DiskMaxNumPages-1 {
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

//NewDiskManagerMock returns a in-memory mock of disk manager
func NewDiskManagerMock() *DiskManagerMock {
	return &DiskManagerMock{-1, make(map[PageID]*Page)}
}
