package objects

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	pathSeparator string = "."
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
// Panics
//
// Panics if any key arugment is non-string or if there are an odd number of arguments.
//
// Example
//
// To easily create Maps:
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

// NewMapFromJSON creates a new map from a JSON string representation
func NewMapFromJSON(data string) (Map, error) {

	var unmarshalled map[string]interface{}

	err := json.Unmarshal([]byte(data), &unmarshalled)

	if err != nil {
		return nil, err
	}

	return Map(unmarshalled), nil

}

// NewMapFromBase64String creates a new map from a Base64 string representation
func NewMapFromBase64String(data string) (Map, error) {

	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))

	decoded, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil, err
	}

	return NewMapFromJSON(string(decoded))

}

// Get gets the value from the map.  Supports deep nesting of other maps,
// For example:
//
//     m = Map{"name":Map{"First": "Mat", "Last": "Ryer"}}
//
//     m.Get("name.Last")
//     // returns "Ryer"
func (d Map) Get(keypath string) interface{} {

	var segs []string = strings.Split(keypath, pathSeparator)

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

// GetString gets a value as a string from the map. It simply calls Get() and casts the result for you.
// If the object is not a string or the object is not found, an error is returned.
func (d Map) GetString(keypath string) (string, error) {
	object := d.Get(keypath)
	if object == nil {
		return "", errors.New("Requested object does not exist")
	}
	if objectAsString, isString := object.(string); isString == false {
		return "", errors.New("Requested object is not a string")
	} else {
		return objectAsString, nil
	}
}

// GetMap gets another Map from this one, or panics if the object is missing or not a Map.
func (d Map) GetMap(keypath string) Map {
	return d.Get(keypath).(Map)
}

// GetWithDefault gets the value at the specified keypath, or returns the defaultValue if
// none could be found.
func (d Map) GetOrDefault(keypath string, defaultValue interface{}) interface{} {
	obj := d.Get(keypath)
	if obj == nil {
		return defaultValue
	}
	return obj
}

// GetWithDefault gets the string value at the specified keypath,
// or returns the defaultValue if none could be found.  Will panic if the
// object is there but of the wrong type.
func (d Map) GetStringOrDefault(keypath, defaultValue string) string {
	obj := d.Get(keypath)
	if obj == nil {
		return defaultValue
	}
	return obj.(string)
}

// Set sets a value in the map.  Supports dot syntax to set deep values.
//
// For example,
//
//     m.Set("name.first", "Mat")
//
// The above code sets the 'first' field on the 'name' object in the m Map.
//
// If objects are nil along the way, Set creates new Map objects as needed.
func (d Map) Set(keypath string, value interface{}) Map {

	var segs []string
	segs = strings.Split(keypath, pathSeparator)

	obj := d

	for fieldIndex, field := range segs {

		if fieldIndex == len(segs)-1 {
			obj[field] = value
		}

		if _, exists := obj[field]; !exists {
			obj[field] = make(Map)
			obj = obj[field].(Map)
		} else {
			switch obj[field].(type) {
			case Map:
				obj = obj[field].(Map)
			case map[string]interface{}:
				obj = Map(obj[field].(map[string]interface{}))
			}
		}

	}

	// chain
	return d
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

// Merge blends the specified map with a copy of this map and returns the result.
//
// Keys that appear in both will be selected from the specified map.
func (d Map) Merge(merge Map) Map {
	return d.Copy().MergeHere(merge)
}

// Merge blends the specified map with this map and returns the current map.
//
// Keys that appear in both will be selected from the specified map.  The original map
// will be modified.
func (d Map) MergeHere(merge Map) Map {

	for k, v := range merge {
		d[k] = v
	}

	return d

}

// Has gets whether the Map has the specified field or not. Supports deep nesting of other maps.
//
// For example:
//     m := map[string]interface{}{"parent": map[string]interface{}{"childname": "Luke"}}
//     m.Has("parent.childname")
//     // return true
func (d Map) Has(path string) bool {
	return d.Get(path) != nil
}

// MSI is a shortcut method to get the current map as a
// normal map[string]interface{}.
func (d Map) MSI() map[string]interface{} {
	return map[string]interface{}(d)
}

// JSON converts the map to a JSON string
func (d Map) JSON() (string, error) {

	result, err := json.Marshal(d)

	return string(result), err

}

// Base64 converts the map to a base64 string
func (d Map) Base64() (string, error) {

	var buf bytes.Buffer

	jsonData, err := d.JSON()
	if err != nil {
		return "", err
	}

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(jsonData))
	encoder.Close()

	return buf.String(), nil

}
