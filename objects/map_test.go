package objects

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

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
	assert.True(t, l.Has("request", "url"))
	assert.False(t, l.Has("request.method"))
	assert.False(t, l.Has("nothing"))

}

func TestGet(t *testing.T) {

	var l Map = Map{"request": Map{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, "http://www.stretchr.com/", l.Get("request.url"))
	assert.Nil(t, l.Get("something.that.doesnt.exist"))

}

func TestGet_WithNativeMap(t *testing.T) {

	var l Map = Map{"request": map[string]interface{}{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, "http://www.stretchr.com/", l.Get("request.url"))

}
