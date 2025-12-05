package ui

import (
	"github.com/derailed/tview"
	"github.com/ramonvermeulen/whosthere/internal/ui/views"
)

// Router is both the visual pages container and the logical router
// that knows about Page implementations and current page state.
type Router struct {
	*tview.Pages

	pages       map[string]views.Page
	currentPage string
}

func NewRouter() *Router {
	return &Router{
		Pages: tview.NewPages(),
		pages: make(map[string]views.Page),
	}
}

// Register adds a Page and attaches its primitive to the underlying tview.Pages.
func (r *Router) Register(p views.Page) {
	name := p.GetName()
	r.pages[name] = p
	r.AddPage(name, p.GetPrimitive(), true, false)
}

// NavigateTo switches to a previously registered page by name.
func (r *Router) NavigateTo(name string) {
	if _, ok := r.pages[name]; !ok {
		return
	}
	r.currentPage = name
	r.SwitchToPage(name)
}

// Current returns the currently active page name.
func (r *Router) Current() string { return r.currentPage }

// Page returns a registered page by name.
func (r *Router) Page(name string) views.Page {
	return r.pages[name]
}
