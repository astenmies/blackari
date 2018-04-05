// Code generated by go-bindata.
// sources:
// gqlSchema/types/mutation.graphql
// gqlSchema/types/post.graphql
// gqlSchema/types/query.graphql
// gqlSchema/types/review.graphql
// gqlSchema/types/reviewInput.graphql
// gqlSchema/types/schema.graphql
// gqlSchema/types/user.graphql
// gqlSchema/types/userInput.graphql
// DO NOT EDIT!

package gqlSchema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _gqlschemaTypesMutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xf0\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xa8\xe6\x52\x50\x50\x50\x48\x2e\x4a\x4d\x2c\x49\x0d\x4a\x2d\xcb\x4c\x2d\xd7\x28\xc8\x2f\x2e\x09\xce\x29\x4d\xb7\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x51\x28\x02\xcb\x58\x29\x40\x54\x78\xe6\x15\x94\x96\x28\x6a\xc2\xb8\x48\x26\x84\x16\xa7\x16\x69\x94\x16\xa7\x16\xe5\x25\xe6\xa6\x22\xe9\x2f\x48\x2c\x2e\x2e\xcf\x2f\x4a\x81\x0b\x69\x5a\x29\x80\xd4\x72\xd5\x02\x02\x00\x00\xff\xff\x6f\x3e\x76\xa4\x8e\x00\x00\x00")

func gqlschemaTypesMutationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesMutationGraphql,
		"gqlSchema/types/mutation.graphql",
	)
}

func gqlschemaTypesMutationGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesMutationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/mutation.graphql", size: 142, mode: os.FileMode(436), modTime: time.Unix(1522965026, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesPostGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\xc8\x2f\x2e\x51\xa8\xe6\x52\x50\x50\x50\x50\x56\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xf3\x8a\x73\x4a\xd3\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x22\x25\x99\x25\x39\xa9\x08\xa1\x5a\x40\x00\x00\x00\xff\xff\x52\xb9\xd0\x74\x40\x00\x00\x00")

func gqlschemaTypesPostGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesPostGraphql,
		"gqlSchema/types/post.graphql",
	)
}

func gqlschemaTypesPostGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesPostGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/post.graphql", size: 64, mode: os.FileMode(436), modTime: time.Unix(1522953605, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesQueryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcb\x31\x0e\xc2\x30\x10\x44\xd1\xde\xa7\x98\x88\x26\x34\x1c\x20\xa7\x00\x42\x87\x28\x28\x86\xc4\x92\x65\xaf\x76\xd7\x20\x0b\x71\x77\x84\xe9\x52\xff\xf7\x77\xb8\xac\xc4\xa9\x52\x1b\xbc\x09\xa1\x14\xa5\x31\xbb\xe1\x9e\x12\xca\x03\xbe\x12\xcc\xae\x0d\x52\x62\x76\x3b\x84\x0e\xff\xcf\x3b\x00\x80\x14\xf3\xd1\x52\x5d\x26\xcc\xae\x31\x2f\xc3\x7e\xc2\xb1\x98\xf7\xaa\x7c\x46\xbe\x6c\xfc\xa9\x79\x83\xae\xe7\x1e\x6f\x43\xf8\x7c\x03\x00\x00\xff\xff\x2f\xdd\xb0\xf5\x8b\x00\x00\x00")

func gqlschemaTypesQueryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesQueryGraphql,
		"gqlSchema/types/query.graphql",
	)
}

func gqlschemaTypesQueryGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesQueryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/query.graphql", size: 139, mode: os.FileMode(436), modTime: time.Unix(1521940364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesReviewGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcd\xb1\xd2\x82\x30\x10\x04\xe0\x3e\x4f\xb1\xff\xd0\xfe\x16\x16\x36\xb4\x56\xb6\xe8\x0b\x04\x5d\x20\x45\x12\xe6\xee\xc0\x61\x1c\xdf\xdd\x11\x88\xe5\xdd\x77\xb7\x5b\xa1\xe1\x28\x54\x26\x53\x78\x08\xe7\xc0\x27\xba\x2c\xbf\xc1\xd9\x32\x12\xcd\x06\x2f\x07\x00\x15\xc2\xa3\xc6\xd5\x24\xa4\xfe\x6f\xdf\xdc\x06\x22\x4d\xb1\xa5\x20\x77\x50\xf3\xa2\xb0\x21\x68\x89\xec\xfd\xcc\x7f\x1c\x0f\xa7\xf5\x7e\xf5\x1a\x97\x64\xfb\xfb\x39\xc7\xc8\x64\xf0\x6d\x9e\xac\x34\x7f\xe9\xbe\x81\x97\xa5\x34\xba\xf7\x27\x00\x00\xff\xff\x09\xe0\x91\xa9\xb5\x00\x00\x00")

func gqlschemaTypesReviewGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesReviewGraphql,
		"gqlSchema/types/review.graphql",
	)
}

func gqlschemaTypesReviewGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesReviewGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/review.graphql", size: 181, mode: os.FileMode(436), modTime: time.Unix(1522954703, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesReviewinputGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\xc1\xaa\xc2\x40\x0c\x45\xf7\xfd\x8a\x0b\xdd\xbe\x07\x6e\xdc\x74\xeb\xaa\x5b\xf5\x07\xd2\x12\xda\x88\x93\x0c\x93\xd4\x41\xc4\x7f\x17\xa7\xee\xce\x21\x27\xdc\x1e\xd7\x95\x21\x9a\xb7\x80\x4d\x37\x9e\x03\xce\x1a\xa8\x2b\x2b\xdc\x12\x9b\x32\xc4\x31\x17\xa6\x10\x5d\x40\x50\xae\x28\xfc\x10\xae\xdd\xfe\x77\x6e\x32\x36\x7e\x75\x00\xd0\xe3\xf0\x7f\x84\x07\x15\x6f\xde\x68\xc0\xa8\xf1\x3b\x9f\x2c\xa5\xef\x0c\x4d\xb6\x05\x62\x65\x64\xf3\xf8\x83\xe5\x10\x53\xba\xb7\x6c\xde\x23\x2a\xcf\x01\x97\x28\xa2\x4b\xf7\xfe\x04\x00\x00\xff\xff\x41\x7f\xeb\x55\xb0\x00\x00\x00")

func gqlschemaTypesReviewinputGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesReviewinputGraphql,
		"gqlSchema/types/reviewInput.graphql",
	)
}

func gqlschemaTypesReviewinputGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesReviewinputGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/reviewInput.graphql", size: 176, mode: os.FileMode(436), modTime: time.Unix(1521845195, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x60\x81\xdc\xd2\x92\xc4\x92\xcc\xfc\x3c\x2b\x05\x5f\x28\x8b\xab\x16\x10\x00\x00\xff\xff\x8e\x43\x79\x00\x32\x00\x00\x00")

func gqlschemaTypesSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesSchemaGraphql,
		"gqlSchema/types/schema.graphql",
	)
}

func gqlschemaTypesSchemaGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/schema.graphql", size: 50, mode: os.FileMode(436), modTime: time.Unix(1521835442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8d\x3b\x0e\xc2\x30\x10\x44\xfb\x9c\x62\x50\x5a\x28\x28\x68\x68\xb9\x01\x9f\x03\x6c\xc4\x24\x71\x61\x3b\xda\x5d\x27\x42\x88\xbb\x23\x50\x9c\x72\x3e\x7a\xaf\xc5\x95\x93\xd2\x98\xdc\x20\x50\xce\x81\x0b\xfa\xac\x5b\x68\xfc\x35\x11\x0f\xa3\xe2\xdd\x00\x40\x8b\xf0\x3c\xe3\xe6\x1a\xd2\xb0\x5b\x9b\xfb\x48\xa4\x12\x3b\x2a\x72\x0f\x73\x51\x83\x8f\xc1\x2a\x70\x90\x99\x7b\x1c\x0f\xa7\xff\xbf\x18\x35\x49\x64\xa5\xac\x90\x4b\x8e\x91\xc9\x21\x5d\x2e\x5e\xed\xbf\x69\x12\xb3\x25\xeb\x66\x6d\x3e\xdf\x00\x00\x00\xff\xff\xbc\x04\xa4\xdb\xb7\x00\x00\x00")

func gqlschemaTypesUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesUserGraphql,
		"gqlSchema/types/user.graphql",
	)
}

func gqlschemaTypesUserGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/user.graphql", size: 183, mode: os.FileMode(436), modTime: time.Unix(1522961960, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesUserinputGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8d\xb1\xaa\xc2\x40\x10\x45\xfb\x7c\xc5\x85\xb4\xef\x81\x8d\x8d\xad\x95\xb5\xfa\x01\x9b\x78\x49\x56\xdc\x99\x65\x66\xe2\x16\xe2\xbf\x8b\x51\xec\xee\x85\x73\x38\x3d\x4e\x33\x91\xa5\x2e\x01\x1d\xae\x1c\x03\x4e\x09\xb4\x99\x02\xd7\x42\x15\x22\x3b\x46\x63\x8a\x2c\x13\x12\x84\x0d\xc6\x7b\x66\xeb\x3e\xde\xd9\x69\x87\x75\x3d\x3a\x00\xe8\xb1\xf9\xdf\xc2\x23\x99\xaf\x7f\x71\x9a\xa4\xc2\x1d\x8e\x61\x59\xa6\x2f\xb4\xd7\x52\xde\xa9\x34\xe8\x12\x88\x99\xa8\xea\xf1\x07\xad\x91\x55\xd2\x6d\xc5\x6a\x72\x6f\x6a\x97\x9f\xfb\x7c\x05\x00\x00\xff\xff\x0a\x56\x2b\xb7\xb2\x00\x00\x00")

func gqlschemaTypesUserinputGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gqlschemaTypesUserinputGraphql,
		"gqlSchema/types/userInput.graphql",
	)
}

func gqlschemaTypesUserinputGraphql() (*asset, error) {
	bytes, err := gqlschemaTypesUserinputGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gqlSchema/types/userInput.graphql", size: 178, mode: os.FileMode(436), modTime: time.Unix(1522961972, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"gqlSchema/types/mutation.graphql": gqlschemaTypesMutationGraphql,
	"gqlSchema/types/post.graphql": gqlschemaTypesPostGraphql,
	"gqlSchema/types/query.graphql": gqlschemaTypesQueryGraphql,
	"gqlSchema/types/review.graphql": gqlschemaTypesReviewGraphql,
	"gqlSchema/types/reviewInput.graphql": gqlschemaTypesReviewinputGraphql,
	"gqlSchema/types/schema.graphql": gqlschemaTypesSchemaGraphql,
	"gqlSchema/types/user.graphql": gqlschemaTypesUserGraphql,
	"gqlSchema/types/userInput.graphql": gqlschemaTypesUserinputGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"gqlSchema": &bintree{nil, map[string]*bintree{
		"types": &bintree{nil, map[string]*bintree{
			"mutation.graphql": &bintree{gqlschemaTypesMutationGraphql, map[string]*bintree{}},
			"post.graphql": &bintree{gqlschemaTypesPostGraphql, map[string]*bintree{}},
			"query.graphql": &bintree{gqlschemaTypesQueryGraphql, map[string]*bintree{}},
			"review.graphql": &bintree{gqlschemaTypesReviewGraphql, map[string]*bintree{}},
			"reviewInput.graphql": &bintree{gqlschemaTypesReviewinputGraphql, map[string]*bintree{}},
			"schema.graphql": &bintree{gqlschemaTypesSchemaGraphql, map[string]*bintree{}},
			"user.graphql": &bintree{gqlschemaTypesUserGraphql, map[string]*bintree{}},
			"userInput.graphql": &bintree{gqlschemaTypesUserinputGraphql, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

