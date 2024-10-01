package gson

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

const gson = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func TestNew(t *testing.T) {
    data := New()
    assert.Equal(t, "{}", data.String())
}

func TestNewArray(t *testing.T) {
    data := NewArray()
    assert.Equal(t, "[]", data.String())
}

func TestParse(t *testing.T) {
    data := Parse(gson)
    assert.Equal(t, gson, data.String())
}

func TestGson_Get(t *testing.T) {
    value := Parse(gson).Get("name.last")
    assert.Equal(t, "Prichard", value.String())
}

func TestGson_Set(t *testing.T) {
    data := Parse(gson)
    _, err := data.Set("name.last", "Jack")
    assert.Nil(t, err)
    value := data.Get("name.last")
    fmt.Println("name.last", "is", value)
    assert.Equal(t, "Jack", value.String())
}

func TestGson_GetGson(t *testing.T) {
    data, err := Parse(gson).GetGson("name").SetRaw("sub", gson)
    assert.Nil(t, err)
    data.Print()
    assert.Equal(t, "Janet", data.GetGson("sub").Get("name.first").String())
}
