package space

type Planet string

func Age(seconds float64, planet Planet) float64 {

	planetsMap := map[string]float64{
		"Earth":   31.69,
		"Mercury": 280.88,
		"Venus":   9.78,
		"Mars":    35.88,
		"Jupiter": 2.41,
		"Saturn":  2.15,
		"Uranus":  0.46,
		"Neptune": 0.35,
	}

	return planetsMap[string(planet)]
}
