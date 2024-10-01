// Package gson provides searching for json strings and setting json values
//
// It's just a simple wrapper for the [gjson] and [sjson]
//
// [gjson]: https://github.com/tidwall/gjson
// [sjson]: https://github.com/tidwall/sjson
package gson

import (
    "encoding/json"
    "fmt"

    "github.com/tidwall/gjson"
    "github.com/tidwall/sjson"
)

// Gson Inherits from [gjson.Result]
//
// Note that when using the Get method, the returned result is [gjson.Result] instead of [Gson]
//
// You may need to use the [Gson.GetGson] method
type Gson struct {
    gjson.Result
}

// New Create an empty json object
func New() *Gson {
    return &Gson{gjson.Parse("{}")}
}

// NewArray Create an empty json array
func NewArray() *Gson {
    return &Gson{gjson.Parse("[]")}
}

// Parse parses the json and returns a result.
//
// This function expects that the json is well-formed, and does not validate.
// Invalid json will not panic, but it may return back unexpected results.
// If you are consuming JSON from an unpredictable source then you may want to
// use the Valid function first.
func Parse(json string) *Gson {
    return &Gson{gjson.Parse(json)}
}

// ParseBytes parses the json and returns a result.
// If working with bytes, this method preferred over Parse(string(data))
func ParseBytes(json []byte) *Gson {
    return &Gson{gjson.ParseBytes(json)}
}

// MarshalJSON Serialize gson to byte slice
func (gson *Gson) MarshalJSON() ([]byte, error) {
    return []byte(gson.String()), nil
}

// UnmarshalJSON Deserialize data to gson
func (gson *Gson) UnmarshalJSON(data []byte) error {
    gson.Result = gjson.ParseBytes(data)
    return nil
}

// GetGson searches result for the specified path. The result should be a JSON array or object.
func (gson *Gson) GetGson(path string) *Gson {
    return &Gson{gson.Get(path)}
}

/*
Set sets a json value for the specified path.
A path is in dot syntax, such as "name.last" or "age".
A value should be a simple JSON value and not a JSON string.
This function expects that the json is well-formed, and does not validate.
Invalid json will not panic, but it may return back unexpected results.
An error is returned if the path is not valid.

A path is a series of keys separated by a dot.

 {
   "name": {"first": "Tom", "last": "Anderson"},
   "age":37,
   "children": ["Sara","Alex","Jack"],
   "friends": [
     {"first": "James", "last": "Murphy"},
     {"first": "Roger", "last": "Craig"}
   ]
 }
 "name.last"          >> "Anderson"
 "age"                >> 37
 "children.1"         >> "Alex"
*/
func (gson *Gson) Set(path string, value interface{}) (json *Gson, err error) {
    var raw string
    raw, err = sjson.Set(gson.Raw, path, value)
    if err == nil {
        gson.Result = gjson.Parse(raw)
    }
    return gson, err
}

// SetRaw sets a raw json value for the specified path.
// This function works the same as Set except that the value is set as a
// raw block of json. This allows for setting premarshalled json objects.
func (gson *Gson) SetRaw(path, value string) (json *Gson, err error) {
    var raw string
    raw, err = sjson.SetRaw(gson.Raw, path, value)
    if err == nil {
        gson.Result = gjson.Parse(raw)
    }
    return gson, err
}

// Delete deletes a value from json for the specified path.
func (gson *Gson) Delete(path string) (json *Gson, err error) {
    var raw string
    raw, err = sjson.Delete(gson.Raw, path)
    if err == nil {
        gson.Result = gjson.Parse(raw)
    }
    return gson, err
}

// Print Pretty print for debug purpose
func (gson *Gson) Print() {
    jsonData, _ := json.MarshalIndent(gson, "", "  ")
    fmt.Println(string(jsonData))
}
