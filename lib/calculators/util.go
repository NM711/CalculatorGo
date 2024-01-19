package calculators;

func (c Calculator) Factorial (n int) float64 {
  factorial := 1.0;
  for i := n; i >= 1; i-- {
    factorial = factorial * float64(i);
  };

  return factorial;
};

func (c Calculator) Power (v float64, n int) float64 {
  result := 1.0;

  for i := 0; i < n; i++ {
    result = result * v
  };

  return result;
};

func (c Calculator) RadToDeg(rad float64) float64 {
  return rad * 180 / PI; 
};

const PI float64 = 3.14159265358979323846264338327950288419716939937510582097494459;
