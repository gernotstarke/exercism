package flatten


func Flatten(list interface{}) []interface{} {
	// I don't quite understand the following
	// (took it from one of the community solutions)
	var flat = []interface{}{}


	for _, element := range list.([]interface{}) {
		switch t := element.(type) {
		case int:
			flat = append( flat, t )
		case []interface{}:
			// again, I'm not sure about the option
			flat = append(flat, Flatten(t)...)
		}
	}

	return flat
}
