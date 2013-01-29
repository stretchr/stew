package objects

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestNewMap(t *testing.T) {

	m := NewMap("name", "Mat", "age", 29, "bool", true)

	assert.Equal(t, "Mat", m["name"])
	assert.Equal(t, 29, m["age"])
	assert.Equal(t, true, m["bool"])

	assert.Panics(t, func() {
		NewMap(1, "Mat", "age", 29, "bool", true)
	}, "Non string key should panic")
	assert.Panics(t, func() {
		NewMap("name", "Mat", "age", 29, "bool")
	}, "Wrong number of arguments should panic")

}

func TestCopy(t *testing.T) {

	d1 := make(Map)
	d1["name"] = "Tyler"
	d1["location"] = "UT"

	d2 := d1.Copy()
	d2["name"] = "Mat"

	assert.Equal(t, d1["name"], "Tyler")
	assert.Equal(t, d2["name"], "Mat")

}

func TestMerge(t *testing.T) {

	d := make(Map)
	d["name"] = "Mat"

	d1 := make(Map)
	d1["name"] = "Tyler"
	d1["location"] = "UT"

	merged := d.Merge(d1)

	assert.Equal(t, merged["name"], d1["name"])
	assert.Equal(t, merged["location"], d1["location"])
	assert.Nil(t, d["location"])

}

func TestExclude(t *testing.T) {

	d := make(Map)
	d["name"] = "Mat"
	d["age"] = 29
	d["secret"] = "ABC"

	excluded := d.Exclude([]string{"secret"})

	assert.Equal(t, d["name"], excluded["name"])
	assert.Equal(t, d["age"], excluded["age"])
	assert.False(t, excluded.Has("secret"), "secret should be excluded")

}

func TestHas(t *testing.T) {

	d := make(Map)
	d["name"] = "Mat"

	assert.True(t, d.Has("name"))
	assert.False(t, d.Has("nope"))

}

func TestHas_WithDeepNesting(t *testing.T) {

	var l Map = Map{"request": Map{"url": "http://www.stretchr.com/"}}

	assert.True(t, l.Has("request.url"))
	assert.False(t, l.Has("request.method"))
	assert.False(t, l.Has("nothing"))

}

func TestGet(t *testing.T) {

	var l Map = Map{"request": Map{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, "http://www.stretchr.com/", l.Get("request.url"))
	assert.Nil(t, l.Get("something.that.doesnt.exist"))

}

func TestSafeGet(t *testing.T) {

	var defaultValue string = "Default"
	var l Map = Map{"request": Map{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, defaultValue, l.GetWithDefault("request.nope", defaultValue))
	assert.Equal(t, "http://www.stretchr.com/", l.GetWithDefault("request.url", defaultValue))

}

func TestGet_WithNativeMap(t *testing.T) {

	var l Map = Map{"request": map[string]interface{}{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, "http://www.stretchr.com/", l.Get("request.url"))

}

func TestSet_Simple(t *testing.T) {
	// https://github.com/stretchrcom/stew/issues/2

	var m Map = make(Map)
	assert.Equal(t, m, m.Set("name", "Tyler"))

	assert.Equal(t, "Tyler", m["name"])

}

func TestSet_Deep(t *testing.T) {
	// https://github.com/stretchrcom/stew/issues/2

	var m Map = make(Map)
	assert.Equal(t, m, m.Set("personal.info.name.first", "Tyler"))

	assert.Equal(t, "Tyler", m.Get("personal.info.name.first"))

	nameObj := m.Get("personal.info.name")
	if assert.NotNil(t, nameObj) {
		assert.Equal(t, "Tyler", nameObj.(Map)["first"])
	}

}
