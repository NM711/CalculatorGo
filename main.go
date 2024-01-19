package main

import (
	"fmt"
	"multipurpose_calc/lib/calculators"
)

func randomCalculations() {
	calc := calculators.Calculator{}

	combinatronics := calc.CombinatronicsCalc()
	solids := calc.GeometrySolidsCalc()

	/* Assume you find a safe code that is 8 digits long. The numbers 0-9 may be used for any digit. How
	   many possible codes are there? Show your work and leave your answer in either exponential or
	   factorial form.
	*/

	combinatronics.PermuatationWithRepetition(10, 8)

	/* How many ways can we list the names Bobby, Billy, Mickey, Sally, Betty, Jimmy, Mary, and Fred?
	   Show your work and leave your answer in either exponential or factorial form.
	*/

	err := combinatronics.PermuatationWithoutRepetition(8, 8)

	if err != nil {
		fmt.Println(err.Error())
	}

	/* Part 1: How many possible phone numbers can we make out of XXX – XXXX, where X can be any
	   number 0-9? Show your work and leave your answer in either exponential or factorial form.
	*/

	combinatronics.PermuatationWithRepetition(10, 7)

	// Part 2 the same without repetition
	err2 := combinatronics.PermuatationWithoutRepetition(10, 7)

	if err2 != nil {
		fmt.Println(err.Error())
	}

	/* Given a four digit number XXXX, are there more permutations if repetitionis allowed or not allowed?
	   Explain your answer
	*/

	combinatronics.PermuatationWithRepetition(10, 4)

	combinatronics.PermuatationWithoutRepetition(10, 4)

	combinatronics.CombinationsWithoutRepitition(6, 3)

	combinatronics.CombinationsWithRepition(7, 3)

	combinatronics.CombinationsWithoutRepitition(5, 3)

	e, v := solids.CheckSegmentType(calculators.SegmentType(calculators.Diameter), 1028)

	fmt.Println(v)

	if e != nil {
		fmt.Println(err.Error())
	}

	volume := solids.CylidnerVolume(v, 325)

	// output should be = 325 or be roundable to 325

	solids.CylinderHeight(v, volume)

	_, r := solids.CheckSegmentType(calculators.SegmentType(calculators.Diameter), 5345)
	solids.SphereVolume(r)

	calc.CircularTrigCalc().TangentRatio(123, 213213)

	/* A sector with a central angle measure of
	    7rad/4

	   (in radians) has a radius of 16cm
	*/

	deg := calc.RadToDeg(7)

	theta := deg / 4

	calc.CircularTrigCalc().SectorArea(16, theta)

	// A sector with a central angle measure of 175 deg has a radius of 12 cm, what is the area of the sector?

	calc.CircularTrigCalc().SectorArea(12, 175);

  combinatronics.CombinationsWithoutRepitition(9,4);
  combinatronics.CombinationsWithoutRepitition(6, 3);

  combinatronics.PermuatationWithRepetition(9, 3);
  combinatronics.CombinationsWithoutRepitition(5, 2);

  combinatronics.PermuatationWithoutRepetition(6, 3);

  combinatronics.CombinationsWithoutRepitition(13, 6);
  combinatronics.CombinationsWithRepition(13, 6)
  combinatronics.PermuatationWithoutRepetition(13, 6)
  combinatronics.PermuatationWithRepetition(36, 12);

  /*
    The roulette wheel contains numbers 1–36. Half the numbers are red,
    and half the numbers are black. The wheel also contains two green
    spaces marked 0 and 00. A ball spins around the wheel and randomly
    lands on a number.

    What is the probability of the ball landing on the number 21?
  */

  combinatronics.Probability(1, 36); // 2.78%

  combinatronics.Probability(3, 8);

  d := calculators.DoublePoint{
    X1: 123,
    X2: 32,
    Y1: 20,
    Y2: 10,
  }

  calc.CoordinateCalc().DoublePointSlope(d);

  calc.CoordinateCalc().Midpoint(d);
};

func main () {
  randomCalculations();
};
