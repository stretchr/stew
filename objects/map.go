package objects

import (
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

// Get gets the value from the map.  Supports deep nesting of other maps,
// For example:
//
//     m = Map{"name":Map{"First": "Mat", "Last": "Ryer"}}
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
