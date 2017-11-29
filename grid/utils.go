package grid

import "github.com/hot-leaf-juice/fgwm/wmutils"

func (g *Grid) getAttributes(wid wmutils.WindowID) (Position, Position, error) {
	pPos, pSize, err := wmutils.GetAttributes(wid)
	if err != nil {
		return Position{}, Position{}, err
	}
	return g.closestPoint(pPos), g.closestPoint(pPos.Offset(pSize)), nil
}

func (g *Grid) closestPoint(p wmutils.Position) Position {
	return Position{
		X: int((p.X - g.margin.W + g.cell.W/2) / g.cell.W),
		Y: int((p.Y - g.margin.H + g.cell.H/2) / g.cell.H),
	}
}

func (g *Grid) inGrid(p Position) bool {
	return 0 <= p.X && p.X <= g.size.W && 0 <= p.Y && p.Y <= g.size.H
}

func (g *Grid) pixelSize(size Size) wmutils.Size {
	return wmutils.Size{
		W: wmutils.Pixels(size.W) * g.cell.W,
		H: wmutils.Pixels(size.H) * g.cell.H,
	}
}

func (g *Grid) pixelPosition(pos Position) wmutils.Position {
	return wmutils.Position{
		X: g.margin.W + wmutils.Pixels(pos.X)*g.cell.W,
		Y: g.margin.H + wmutils.Pixels(pos.Y)*g.cell.H,
	}
}