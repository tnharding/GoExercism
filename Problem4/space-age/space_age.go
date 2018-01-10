//Package space contains a function to compute your age in years on different planets
//in earth's solar system
package space

//Planet type represents a planet by name
type Planet string

const secondsPerEarthYear = 31557600

//Age will compute your age on a certain planet given the number
//of seconds you have been alive.
func Age(seconds float64, planet Planet) float64 {

	switch planet {
	case "Earth":
		return seconds / secondsPerEarthYear

	case "Mercury":
		return seconds / (secondsPerEarthYear * 0.2408467)

	case "Venus":
		return seconds / (secondsPerEarthYear * 0.61519726)

	case "Mars":
		return seconds / (secondsPerEarthYear * 1.8808158)

	case "Jupiter":
		return seconds / (secondsPerEarthYear * 11.862615)

	case "Saturn":
		return seconds / (secondsPerEarthYear * 29.447498)

	case "Uranus":
		return seconds / (secondsPerEarthYear * 84.016846)

	case "Neptune":
		return seconds / (secondsPerEarthYear * 164.79132)

	}
	return 0.0
}
