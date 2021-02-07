package bpm

// PageID is the type of the page identifier
type PageID int

const pageSize = 5

// Page represents a page on disk
type Page struct {
	id       PageID
	pinCount int
	isDirty  bool
	data     [pageSize]byte
}

// PinCount retunds the pin count
func (p *Page) PinCount() int {
	return p.pinCount
}

// ID retunds the page id
func (p *Page) ID() PageID {
	return p.id
}

// DecPinCount decrements pin count
func (p *Page) DecPinCount() {
	if p.pinCount > 0 {
		p.pinCount--
	}
}
