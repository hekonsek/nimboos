// Code generated by go-bindata.
// sources:
// data/env-create.yml
// data/env-remove.yml
// DO NOT EDIT!

package main

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

var _dataEnvCreateYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8a\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x3f\x80\x29\xea\x32\x97\x91\x61\xf8\x6d\x4a\x87\x19\xe9\x04\x37\xe2\xdd\x25\xe2\xee\xbd\xc7\xab\xe8\x91\x23\x1b\x2c\x54\x6c\x72\x01\x34\xdc\xa9\x63\x0f\xff\xf7\x02\x6c\x32\x3a\xcf\xc7\x2a\x3a\xef\x55\x2c\x59\x80\x33\x8c\xd9\x0a\x00\x54\xbc\x7f\xda\xd0\x79\x84\x27\x8f\x85\x7a\xaf\xaf\xa7\xd6\xdc\x7d\x33\x5e\xae\xcb\x0d\x9f\x6f\x00\x00\x00\xff\xff\xb5\x70\x01\x63\x71\x00\x00\x00")

func dataEnvCreateYmlBytes() ([]byte, error) {
	return bindataRead(
		_dataEnvCreateYml,
		"data/env-create.yml",
	)
}

func dataEnvCreateYml() (*asset, error) {
	bytes, err := dataEnvCreateYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/env-create.yml", size: 113, mode: os.FileMode(436), modTime: time.Unix(1519984153, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataEnvRemoveYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8a\x41\x0a\xc3\x30\x0c\x04\xef\x7e\xc5\x3e\xa0\x0a\xa1\xc7\x7c\xa6\x18\xb1\x89\x43\x54\xa9\x58\x26\x97\xd2\xbf\x97\x94\xde\x66\x86\x11\xb4\xc8\x91\x0b\x2c\xb4\xda\xc5\x05\xd0\x70\xa7\x8e\x3d\xfc\xdf\x0b\xb0\xd5\xd1\xd8\x1f\x6b\xd5\xeb\x5e\xab\x25\x0b\xd0\xc3\x98\x4b\x01\x00\xc1\xfb\xa7\x0b\x1a\x8f\xf0\xe4\x31\x51\xef\x72\xbe\x54\x72\xf7\xcd\x28\x9d\xcf\x38\x79\x9b\xa7\x19\x9f\x6f\x00\x00\x00\xff\xff\xd5\x11\x49\x2c\x78\x00\x00\x00")

func dataEnvRemoveYmlBytes() ([]byte, error) {
	return bindataRead(
		_dataEnvRemoveYml,
		"data/env-remove.yml",
	)
}

func dataEnvRemoveYml() (*asset, error) {
	bytes, err := dataEnvRemoveYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/env-remove.yml", size: 120, mode: os.FileMode(436), modTime: time.Unix(1519984065, 0)}
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
	"data/env-create.yml": dataEnvCreateYml,
	"data/env-remove.yml": dataEnvRemoveYml,
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
	"data": &bintree{nil, map[string]*bintree{
		"env-create.yml": &bintree{dataEnvCreateYml, map[string]*bintree{}},
		"env-remove.yml": &bintree{dataEnvRemoveYml, map[string]*bintree{}},
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

