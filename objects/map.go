package objects

import (
	"fmt"
	"strings"
)

// Map is a map[string]interface{} with additional helpful functionality.
//
// You can use Map functionality on any map[string]interface{} using the following
// format:
//
//     data := map[string]interface{}{"name": "Stew"}
//     objects.Map(data).Get("name")
//     // returns "Stew"
type Map map[string]interface{}

// NewMap creates a new map.
//
// The arguments follow a key, value pattern.
//
// Panics if:
//
// * any key arugment is non-string
// * there are an odd number of arguments
//
// Example
//
//     m := objects.NewMap("name", "Mat", "age", 29, "subobj", objects.NewMap("active", true))
//
//     // creates a Map equivalent to
//     m := map[string]interface{}{"name": "Mat", "age": 29, "subobj": map[string]interface{}{"active": true}}
func NewMap(keyAndValuePairs ...interface{}) Map {

	newMap := make(Map)
	keyAndValuePairsLen := len(keyAndValuePairs)

	if keyAndValuePairsLen%2 != 0 {
		panic("NewMap must have an even number of arguments following the 'key, value' pattern.")
	}

	for i := 0; i < keyAndValuePairsLen; i = i + 2 {

		key := keyAndValuePairs[i]
		value := keyAndValuePairs[i+1]

		// make sure the key is a string
		keyString, keyStringOK := key.(string)
		if !keyStringOK {
			panic(fmt.Sprintf("NewMap must follow 'string, interface{}' pattern.  %s is not a valid key.", keyString))
		}

		newMap[keyString] = value

	}

	return newMap
}

// Get gets the value from the map.  Supports deep nesting of other maps,
// For example:
//
//     m = Map{"name":Map{"First": "Mat", "Last": "Ryer"}}
//     
//     m.Get("name", "Last")
//     // returns "Ryer"
//
//     //... or dot notaion (not as quick)
//     m.Get("name.Last")
//     // returns "Ryer"
func (d Map) Get(path ...string) interface{} {

	var segs []string
	if len(path) == 1 {
		segs = strings.Split(path[0], ".")

	} else {
		segs = path
	}

	obj := d

	for fieldIndex, field := range segs {

		if fieldIndex == len(segs)-1 {
			return obj[field]
		}

		switch obj[field].(type) {
		case Map:
			obj = obj[field].(Map)
		case map[string]interface{}:
			obj = Map(obj[field].(map[string]interface{}))
		}

	}

	return obj

}

// Exclude returns a new Map with the keys in the specified []string
// excluded.
func (d Map) Exclude(exclude []string) Map {

	excluded := make(Map)
	for k, v := range d {
		var shouldInclude bool = true
		for _, toExclude := range exclude {
			if k == toExclude {
				shouldInclude = false
				break
			}
		}
		if shouldInclude {
			excluded[k] = v
		}
	}

	return excluded
}

// Copy creates a shallow copy of the Map.
func (d Map) Copy() Map {
	copied := make(Map)
	for k, v := range d {
		copied[k] = v
	}
	return copied
}

// Merge blends the specified map with this map and returns the result.
//
// Keys that appear in both will be selected from the specified map.
func (d Map) Merge(merge Map) Map {

	merged := d.Copy()

	for k, v := range merge {
		merged[k] = v
	}

	return merged

}

// Has gets whether the Map has the specified field or not. Supports deep nesting of other maps.
//
// For example:
//     m := map[string]interface{}{"parent": map[string]interface{}{"childname": "Luke"}}
//     m.Has("parent.childname")
//     // return true
//
//     // or the more efficient:
//     m.Has("parent", "childname")
//     // also returns true
func (d Map) Has(path ...string) bool {
	return d.Get(path...) != nil
}
