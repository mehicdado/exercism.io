package space

//Planet custom string type for planet name
type Planet string

//Age will calculate how many given planet years is someone old,
//based on the age given in seconds and planet name
func Age(ageInSec float64, planet Planet) float64 {
	var numOfYears float64
	var earthYearInSec float64 = 31557600.0

	orbitalPeriods := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	numOfYears = ageInSec / (orbitalPeriods[planet] * earthYearInSec)

	return numOfYears
}
