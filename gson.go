// Package gson provides searching for json strings and setting json values
// It's just a simple wrapper for the [gjson]: https://github.com/tidwall/gjson and [sjson]: https://github.com/tidwall/sjson
package gson

import (
    "encoding/json"
    "fmt"

    "github.com/tidwall/gjson"
    "github.com/tidwall/sjson"
)

// Gson Note that when using the Get method, the returned result is gjson.Result instead of Gson
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

/*
Set sets a json value for the specified path.
A path is in dot syntax, such as "name.last" or "age".
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
func (gson *Gson) Set(path string, value interface{}) (err error) {
    var raw string
    raw, err = sjson.Set(gson.Raw, path, value)
    gson.Result = gjson.Parse(raw)
    return
}

// SetRaw sets a raw json value for the specified path.
// This function works the same as Set except that the value is set as a
// raw block of json. This allows for setting premarshalled json objects.
func (gson *Gson) SetRaw(path, value string) (err error) {
    var raw string
    raw, err = sjson.SetRaw(gson.Raw, path, value)
    gson.Result = gjson.Parse(raw)
    return
}

// Delete deletes a value from json for the specified path.
func (gson *Gson) Delete(path string) (err error) {
    var raw string
    raw, err = sjson.Delete(gson.Raw, path)
    gson.Result = gjson.Parse(raw)
    return
}

// Print Pretty print for debug purpose
func (gson *Gson) Print() {
    jsonData, _ := json.MarshalIndent(gson, "", "  ")
    fmt.Println(string(jsonData))
}
