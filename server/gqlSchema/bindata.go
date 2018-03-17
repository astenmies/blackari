// Code generated by go-bindata.
// sources:
// gqlSchema/types/post.graphql
// gqlSchema/types/query.graphql
// gqlSchema/types/schema.graphql
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

var _gqlschemaTypesPostGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\xc8\x2f\x2e\x51\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\x8b\x73\x4a\xd3\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x22\x25\x99\x25\x39\xa9\x08\xa1\x5a\x40\x00\x00\x00\xff\xff\x98\xcb\xa3\x23\x3e\x00\x00\x00")

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

	info := bindataFileInfo{name: "gqlSchema/types/post.graphql", size: 62, mode: os.FileMode(436), modTime: time.Unix(1521313133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesQueryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcb\xb1\x09\x02\x41\x10\x05\xd0\x7c\xab\xf8\x62\xa2\x89\x05\x5c\x15\x8a\x36\x60\xf0\xbd\x3b\x58\x66\x86\x99\xbf\xc1\x22\xf6\x2e\x68\xfe\xde\x11\x8f\x8d\xb8\x0d\xe6\x84\x66\x10\xc9\x48\x16\x4d\x85\x67\xef\xf0\x17\xb4\x11\x34\xe5\x44\xf8\x6e\xaa\x4b\xfb\xc1\xff\x79\x37\x00\x08\x2f\x9d\xaa\x8f\x75\xc1\x5d\xb9\xdb\x7a\x38\x2f\xb8\x7a\xa9\x7d\xbe\x01\x00\x00\xff\xff\xe0\xc2\xee\x50\x61\x00\x00\x00")

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

	info := bindataFileInfo{name: "gqlSchema/types/query.graphql", size: 97, mode: os.FileMode(436), modTime: time.Unix(1521313127, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gqlschemaTypesSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\xb5\x80\x00\x00\x00\xff\xff\x54\xe0\x78\x3a\x1b\x00\x00\x00")

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

	info := bindataFileInfo{name: "gqlSchema/types/schema.graphql", size: 27, mode: os.FileMode(436), modTime: time.Unix(1521313182, 0)}
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
	"gqlSchema/types/post.graphql": gqlschemaTypesPostGraphql,
	"gqlSchema/types/query.graphql": gqlschemaTypesQueryGraphql,
	"gqlSchema/types/schema.graphql": gqlschemaTypesSchemaGraphql,
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
			"post.graphql": &bintree{gqlschemaTypesPostGraphql, map[string]*bintree{}},
			"query.graphql": &bintree{gqlschemaTypesQueryGraphql, map[string]*bintree{}},
			"schema.graphql": &bintree{gqlschemaTypesSchemaGraphql, map[string]*bintree{}},
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

