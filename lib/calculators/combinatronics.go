package calculators

import (
	"errors"
	"fmt"
)

type CombinatronicsCalc struct {Calculator};

func (c Calculator) CombinatronicsCalc() CombinatronicsCalc {
  return CombinatronicsCalc{c};
};

// Probability Formula = P = (f / p) * 100

// f = Favorable Cases
// p = Possible Cases

func (c CombinatronicsCalc) Probability(f float64, p float64) float64 {
  result := (f / p) * 100;
  fmt.Printf("Probability = %f%s\n", result, "%");
  return result;
};

// Permuatation Without Repetition Formula = P(n,r) = n!/(n-r)!

// n = total objects
// s = total selected objects

func (c CombinatronicsCalc) PermuatationWithoutRepetition(totalObjs float64, selectedObjs float64) error {
  if selectedObjs > totalObjs {
    return errors.New(`Selected objscts "r" cannot be greater than the total number of objects "n"!`);
  };

  value := c.Factorial(int(totalObjs)) / c.Factorial(int(totalObjs) - int(selectedObjs));
  fmt.Println(value);
  return nil;
};

// Permutation With Repetition Formula = P(n, r) = n^r;

// n = total objects
// s = total selected objects

func (c CombinatronicsCalc) PermuatationWithRepetition(totalObjs float64, selectedObjs float64) {
  result := c.Power(totalObjs, int(selectedObjs));
  fmt.Printf("%f\n", result);
};

// Combinations With Repition Formula = C(n, r) = (n + r - 1)! / (n - 1)!r!

func (c CombinatronicsCalc) CombinationsWithRepition(n float64, r float64) {
  result := c.Factorial(int(n) + int(r) - 1) / (c.Factorial(int(n) - 1) * c.Factorial(int(r)))
  fmt.Println(result);
};

// Combination Without Repition Formula = C(n, r) = n! / r!(n - r)!

// n = Total Elements
// r = Select Sample Size

func (c CombinatronicsCalc) CombinationsWithoutRepitition(n float64, r float64) error {

    if (r <= 0) {
      return errors.New(`Sample Size "r" cannot be less than or equal to 0!`);
    };

    result := c.Factorial(int(n)) / (c.Factorial(int(r)) * c.Factorial(int(n - r)));

    fmt.Println(result);

    return nil;
};
