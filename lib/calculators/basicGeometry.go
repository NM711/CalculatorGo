package calculators;

// These serve more as helper functions for bigger formulas.
// For example when dealing with 3D shapes, you may or may not need to calculate a 2D base

type BasicGeometryCalc struct {Calculator};

func (c Calculator) BasicGeometryCalc() BasicGeometryCalc {
  return BasicGeometryCalc{c};
};

// Rectangle Area = A = lw

func (c BasicGeometryCalc) RectangleArea(l float64, w float64) float64 {
  return l * w;
};

// Triangle Area = A = (1/2)bh

func (c BasicGeometryCalc) TrianguleArea(b float64, h float64) float64 {
  return 1 / 2 * (b * h)
};

// Square Area = A = s^2

func (c BasicGeometryCalc) SquareArea(s float64) float64 {
  return c.Power(s, 2);
};
