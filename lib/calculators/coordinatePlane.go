package calculators

import "fmt"
import "math";

type CoordinateCalc struct{Calculator};

func (c Calculator) CoordinateCalc() CoordinateCalc {
  return CoordinateCalc{c};
};

// Reflection over x axis (x, y) = (x, -y)
// Reflection over y axis (x, y) = (-x, y)

type Axis int;

const (
  X Axis = iota;
  Y
);

func (c CoordinateCalc) ReflectionXY(x float64, y float64, reflect Axis) {
  if (reflect == Axis(X)) {
    fmt.Printf("X-Axist Reflection = (%f, -%f)\n", x, y);
  } else {
    fmt.Printf("Y-Axis Reflection = (-%f, %f)\n", x, y);
  };
};

type DoublePoint struct {
  X1 float64;
  Y1 float64;
  X2 float64;
  Y2 float64;
};

// Distance Formula = d = sqrt((x2 - x1)^2 + (y2 - y1)^2)

func (c CoordinateCalc) Distance(f DoublePoint) float64 {
  result := math.Sqrt( c.Power((f.X2 - f.X1), 2)  + c.Power((f.Y2 - f.Y1), 2));

  fmt.Printf("Distance = %f\n", result);

  return result;
};

// Single Point Slope Formula = m = y/x

func (c CoordinateCalc) SinglePointSlope(x float64, y float64) float64 {
  result := y / x;

  fmt.Printf("Slope = %f\n", result);

  return result;
};

// Double Point Slope Formula = m = (y2 - y1)/(x2 - x1)

func (c CoordinateCalc) DoublePointSlope(f DoublePoint) float64 {
  result := (f.Y2 - f.Y1) / (f.X2 - f.X1);

  fmt.Printf("Slope = %f\n", result);

  return result;
};

// Midpoint Formula = (xm, ym) = ((x1 + x2) / 2, (y1 + y2/ 2))

func (c CoordinateCalc) Midpoint(f DoublePoint) (float64, float64) {
  xm := (f.X1 + f.X2) / 2;
  ym := (f.Y1 + f.Y2) / 2;

  fmt.Printf("Midpoint = (%f, %f)\n", xm, ym);

  return xm, ym;
};
