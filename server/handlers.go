package server

import (
	"github.com/hot-leaf-juice/fgwm/grid"
	"github.com/hot-leaf-juice/fgwm/wmutils"
)

func (s *Server) Snap(struct{}, *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Snap(wid)
}

func (s *Server) Center(struct{}, *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Center(wid)
}

func (s *Server) Fullscreen(struct{}, *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Fullscreen(wid)
}

func (s *Server) Kill(struct{}, *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return wmutils.Kill(wid)
}

func (s *Server) Move(diff grid.Size, _ *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Move(wid, diff)
}

func (s *Server) Grow(diff grid.Size, _ *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Grow(wid, diff)
}

func (s *Server) Throw(direction grid.Direction, _ *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Throw(wid, direction)
}

func (s *Server) Spread(direction grid.Direction, _ *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Spread(wid, direction)
}

func (s *Server) Focus(nextOrPrev grid.NextOrPrev, _ *struct{}) error {
	return s.grid.Focus(nextOrPrev)
}

func (s *Server) Teleport(rectangle grid.Rectangle, _ *struct{}) error {
	wid, err := wmutils.Focussed()
	if err != nil {
		return err
	}
	return s.grid.Teleport(wid, rectangle)
}
