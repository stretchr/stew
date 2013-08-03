package objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var signatureTestKey = "e1zJJGjCJfLAR1b4dDqg0PY33731D8gM"

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

func TestMSI(t *testing.T) {

	m := NewMap("name", "Mat", "age", 29, "bool", true)

	var msi map[string]interface{} = m.MSI()

	assert.Equal(t, "Mat", msi["name"])
	assert.Equal(t, 29, msi["age"])
	assert.Equal(t, true, msi["bool"])

}

func TestMergeHere(t *testing.T) {

	d := make(Map)
	d["name"] = "Mat"

	d1 := make(Map)
	d1["name"] = "Tyler"
	d1["location"] = "UT"

	merged := d.MergeHere(d1)

	assert.Equal(t, d, merged, "With MergeHere, it should return the first modified map")

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

	// test some fail cases
	assert.Nil(t, l.Get("something.that.doesnt.exist"))
	assert.Nil(t, l.Get("request.url.somethingelse"))
	assert.Nil(t, l.Get("request.somethingelse"))

}

func TestGetOrDefault(t *testing.T) {

	var defaultValue string = "Default"
	var l Map = Map{"request": Map{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, defaultValue, l.GetOrDefault("request.nope", defaultValue))
	assert.Equal(t, "http://www.stretchr.com/", l.GetOrDefault("request.url", defaultValue))

}

func TestGetString(t *testing.T) {

	var l Map = Map{"request": Map{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, l.GetString("request.url"), "http://www.stretchr.com/")

}

func TestGetStringOrDefault(t *testing.T) {

	var l Map = Map{"request": Map{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, l.GetStringOrDefault("request.url", "default"), "http://www.stretchr.com/")
	assert.Equal(t, l.GetStringOrDefault("request.nope", "default"), "default")

}

func TestGet_WithNativeMap(t *testing.T) {

	var l Map = Map{"request": map[string]interface{}{"url": "http://www.stretchr.com/"}}

	assert.Equal(t, "http://www.stretchr.com/", l.Get("request.url"))

}

func TestSet_Simple(t *testing.T) {
	// https://github.com/stretchr/stew/issues/2

	var m Map = make(Map)
	assert.Equal(t, m, m.Set("name", "Tyler"))

	assert.Equal(t, "Tyler", m["name"])

}

func TestSet_Deep(t *testing.T) {
	// https://github.com/stretchr/stew/issues/2

	var m Map = make(Map)
	assert.Equal(t, m, m.Set("personal.info.name.first", "Tyler"))

	assert.Equal(t, "Tyler", m.Get("personal.info.name.first"))

	nameObj := m.Get("personal.info.name")
	if assert.NotNil(t, nameObj) {
		assert.Equal(t, "Tyler", nameObj.(Map)["first"])
	}

}

func Test_GetMap(t *testing.T) {

	var parent Map = make(Map)
	var child Map = make(Map)
	child.Set("name", "child")

	parent.Set("child", child)

	var gottenChild Map = parent.GetMap("child")
	assert.Equal(t, "child", gottenChild.Get("name"))

}

func TestMapJSON(t *testing.T) {

	m := make(Map)

	m.Set("name", "tyler")

	json, err := m.JSON()

	if assert.NoError(t, err) {
		assert.Equal(t, json, "{\"name\":\"tyler\"}")
	}

}

func TestMapNewMapFromJSON(t *testing.T) {

	m, err := NewMapFromJSON("{\"name\":\"tyler\"}")

	if assert.NotNil(t, m) && assert.NoError(t, err) {
		assert.Equal(t, m.Get("name").(string), "tyler")
	}

}

func TestMapBase64(t *testing.T) {

	m := make(Map)

	m.Set("name", "tyler")

	b64, err := m.Base64()

	if assert.NoError(t, err) {
		assert.Equal(t, b64, "eyJuYW1lIjoidHlsZXIifQ==")
	}

}

func TestMapSignedBase64(t *testing.T) {

	m := make(Map)

	m.Set("name", "tyler")

	b64, err := m.SignedBase64(signatureTestKey)

	if assert.NoError(t, err) {
		assert.Equal(t, b64, "eyJuYW1lIjoidHlsZXIifQ==_125052af5002afcf68f5b83089756c62cc139b97")
	}

}

func TestNewMapFromBase64String(t *testing.T) {

	m, err := NewMapFromBase64String("eyJuYW1lIjoidHlsZXIifQ==")

	if assert.NotNil(t, m) && assert.NoError(t, err) {
		assert.Equal(t, m.Get("name").(string), "tyler")
	}

}

func TestNewMapFromSignedBase64String(t *testing.T) {

	// malformed string
	m, err := NewMapFromSignedBase64String("eyJuYW1lIjoidHlsZXIifQ==125052af5002afcf68f5b83089756c62cc139b97", signatureTestKey)
	if assert.Error(t, err) {
		assert.Nil(t, m)
	}

	// altered signature
	m, err = NewMapFromSignedBase64String("eyJuYW1lIjoidHlsZXIifQ==_125052af5002afcf68f5b83089756c62cc139b97BREAK", signatureTestKey)
	if assert.Error(t, err) {
		assert.Nil(t, m)
	}

	// altered data
	m, err = NewMapFromSignedBase64String("eyJuYW1lIjoidHlXIifQ==_125052af5002afcf68f5b83089756c62cc139b97", signatureTestKey)
	if assert.Error(t, err) {
		assert.Nil(t, m)
	}

	// correct string
	m, err = NewMapFromSignedBase64String("eyJuYW1lIjoidHlsZXIifQ==_125052af5002afcf68f5b83089756c62cc139b97", signatureTestKey)

	if assert.NotNil(t, m) && assert.NoError(t, err) {
		assert.Equal(t, m.Get("name").(string), "tyler")
	}

}

func TestMapHash(t *testing.T) {

	m := make(Map)

	m.Set("name", "tyler")

	hash, err := m.Hash()

	if assert.NoError(t, err) {
		assert.Equal(t, hash, "4100f62944bafb39f3cd36a08fe7094482b69207")
	}

}

func TestMapHashWithKey(t *testing.T) {

	m := make(Map)

	m.Set("name", "tyler")

	hash, err := m.HashWithKey(signatureTestKey)

	if assert.NoError(t, err) {
		assert.Equal(t, hash, "125052af5002afcf68f5b83089756c62cc139b97")
	}

}
