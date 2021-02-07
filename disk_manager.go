package bpm

// DiskManager is responsible for interacting with disk
type DiskManager interface {
	ReadPage(PageID) (*Page, error)
	WritePage(*Page) error
	AllocatePage() *PageID
	DeallocatePage(PageID)
	PagesInDisk() []int
	MaxNumPages() int
}
