package calculators

import "fmt"

type CircularTrigCalc struct{ Calculator }

func (c Calculator) CircularTrigCalc() CircularTrigCalc {
	return CircularTrigCalc{c}
};

// Circle Area Formula = A = pi * r^2

func (c Calculator) CircleArea(r float64) float64 {
  result := PI * c.Power(r, 2);

  println(result);

  return result;
};

// Tangent Ratio Formula = Tan (Theta) = opposite/adjacent

// In reality I probably do not even need a calculator for this because its just division
// But I mean at least you have a function that remembers the formula for you ig.

func (c CircularTrigCalc) TangentRatio(opp float64, adj float64) float64 {
	result := opp / adj
	fmt.Printf("Tangent Ratio %f\n", result);
  return result;
};

// Sector Area Formula #1 = A = (theta/360 deg) * pi * r^2;
// Sector Area Formula #2 = A = (theta/2rad) * pi * r^2;

// We will use the first case for this function and provide a helper function that can convert radiants to degrees.

func (c CircularTrigCalc) SectorArea(r float64, theta float64) float64 {
  result := theta / 360.0 * (PI * c.Power(r, 2));

  fmt.Printf("Sector Area = %f\n", result);

  return result;
};
