package calculators

import (
	"errors"
	"fmt"
);

type SegmentType int

const (
	Diameter SegmentType = 0
	Radius               = 1
)

type GeometrySolidsCalc struct{ Calculator }

func (c Calculator) GeometrySolidsCalc() GeometrySolidsCalc {
	return GeometrySolidsCalc{c}
}

// all this does is return a diamter if radius is given, and return radius when diameter is given.
func (c GeometrySolidsCalc) CheckSegmentType(t SegmentType, input float64) (error, float64) {
	if t == SegmentType(Diameter) {
		return nil, input / 2
	} else if t == SegmentType(Radius) {
		return nil, input * 2
	} else {
		return errors.New("No valid segment type selected!"), 0;
	};
};

// Circle Circumference Formula = C = 2PIr

func (c GeometrySolidsCalc) GetCircumference(r float64) float64 {
  return 2 * PI * r;
};

// Cylinder Volume Formula = V = PIr^2h

func (c GeometrySolidsCalc) CylidnerVolume(r float64, h float64) float64 {
  result := PI * (float64(c.Power(r, 2))) * h;
  fmt.Printf("Cylinder V = %f\n", result);

  return result;
};

// Cylinder Height Formula = h = V / PIr^2

func (c GeometrySolidsCalc) CylinderHeight(r float64, v float64) float64 {
  result := v / (PI * c.Power(r, 2));
  fmt.Printf("Cylinder h = %f\n", result);
  return result;
};


// Prism Height Formula h = V/B

func (c GeometrySolidsCalc) PrismHeight(b float64, v float64) float64 {
  result := v / b;

  fmt.Println(result);

  return result;
};

// Prism Height Formula V = Bh

func (c GeometrySolidsCalc) PrismVolume(b float64, h float64) float64 {
  result := b * h;

  fmt.Println(result);

  return result;
};

// Sphere Volume Formula = 4/3(PI)(r^3)

func (c GeometrySolidsCalc) SphereVolume(r float64) float64 {
 result := 4.0 / 3.0 * (PI * c.Power(r, 3));

 fmt.Printf("Sphere V = %f\n", result);

 return result;
};

// Pyramid Volume Formula = V = (1/3)Bh

func (c GeometrySolidsCalc) PyramidVolume(b float64, h float64) float64 {
  result := 1.0 / 3.0 * (b * h);

  fmt.Printf("Pyramid V = %f\n", result);

  return result;
};

// Pyramid Height Formula = h = 3(V/b)

func (c GeometrySolidsCalc) PyramidHeight(b float64, v float64) float64 {
  result := 3 * (v / b);

  fmt.Printf("Pyramid h = %f\n", result);

  return result;
};
