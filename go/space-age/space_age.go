package space

type Planet string

// const planets ( "Mercury", 	"Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune" )

func Age(seconds float64, planet Planet) float64 {
	var secsInYear float64
	secsInYear = 365.25 * 24 * 60 * 60
	var planetyears = map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 84.016846}

	return (seconds / secsInYear) / planetyears[planet]

}
