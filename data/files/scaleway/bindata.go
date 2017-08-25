// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package scaleway

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x96\x4d\x6f\xdb\x3c\x0c\x80\xef\xf9\x15\x84\xce\x7e\x0b\xdb\x75\xdd\x20\xb7\xb7\x29\xda\xcb\xba\x0d\x48\x31\x0c\x18\x7a\x50\x6c\x2e\xd3\x22\x4b\x9e\xa4\x64\x4d\x83\xfc\xf7\x41\x4a\x62\xc7\x1f\x72\xd0\xe6\x52\x57\xfc\x10\x1f\x2a\x14\xa9\xed\x08\x80\x08\x5a\x20\x99\x00\xd1\x19\xe5\xf8\x97\x6e\x48\x60\xa5\x28\xd6\x64\x02\x3f\x46\x00\x00\x24\xc7\xb5\x93\x02\x90\x3f\x94\x8c\x00\x5e\x9c\x8d\xc2\x05\x93\x42\x57\x76\x5b\xf7\x17\x80\x70\x99\x51\xc3\xa4\xb0\xdb\x7e\xa5\x8a\xe9\x00\x1e\x14\x15\x19\x1e\x76\xa9\x7c\xad\x41\x49\x55\x54\xcb\xdf\xa4\xc0\x7a\x47\x27\x72\x06\x87\xe5\x8b\xfb\xee\x02\x7f\xbc\xff\x0b\x6d\x50\xe5\xb4\x08\xe0\x33\x9a\x5f\xa8\x38\x15\xb9\xee\x0b\x4c\x0b\x3d\x1c\xd8\x19\x34\x03\x57\xb9\x33\xa1\x8d\xcd\xe8\x79\x53\x62\xcf\x09\xe8\xe5\xca\x86\xf8\x36\x8d\x66\x75\x88\x1c\x75\xa6\x58\x79\x24\x8d\xe1\x75\x9c\x42\x9a\xcc\x99\x81\xa9\x54\xa8\x03\x88\x1f\xef\xa0\xc0\x42\xaa\x4d\xed\x95\x51\x83\x0b\x2b\x99\x00\x99\x72\xb9\xca\x61\x86\x6a\x8d\xea\x24\xa7\xac\xb4\xd1\xe2\x3a\x47\x5a\x34\xd6\x39\xd3\x4b\x32\x81\x9b\xb0\xf7\xf4\x6a\xd6\x27\x2f\x6b\xd2\x65\x4d\x2e\x61\x4d\x5a\xac\x49\x9b\x35\x0a\xcf\xc1\x7e\xf2\xc2\xa6\x5d\xd8\xf1\x25\xb0\x69\x0b\x76\xdc\x86\x8d\x87\x61\xbf\xa7\xc9\x7f\x69\xf8\x78\xe7\x05\x8e\xc2\x2e\xb1\x75\xf8\x38\x72\x14\xb6\x98\xd3\xb0\x0d\x7d\x7b\x1e\x3a\x8a\x07\xa9\x7b\x0a\xd8\x79\x5c\x80\xdd\xae\xe1\x28\xee\x70\x47\xe1\x30\xf8\x34\xf6\x5f\xb9\x04\xee\x31\x67\x16\x27\x7f\x7f\x8d\xdc\x51\x85\x4f\x68\x28\x3f\x57\xcc\x9d\xfa\x18\xbe\x78\xd3\xd8\x7f\xef\xc6\x43\xc0\x51\xfa\x51\xe2\x71\xfb\x98\xd3\xf7\x22\xfb\x6f\xdf\x20\xf2\x75\xdd\xe1\x02\x88\x6f\x6c\xb1\xdc\x33\x85\x99\x81\xd9\xec\xfe\xc2\x24\xae\xbd\x0d\xaf\xea\xda\x99\xc2\x1c\x85\x61\x94\xf7\xf4\x6c\xb3\x29\xdd\x20\x9c\x9d\x0e\xc2\xe3\x76\x25\xa7\x9b\x07\xa9\x0a\x6a\xac\xc9\x4f\x86\x3c\xaf\xf5\x6e\xd9\x1c\x1e\xdb\xea\x3f\x00\xf2\x5b\xef\x8f\x46\xaa\x05\x15\xec\x6d\x3f\xab\x82\x53\x8b\xe3\x10\xf6\x5b\x70\x3a\x47\x6e\x4d\xbe\x78\x4d\x98\x28\x57\x0e\xcf\xe0\xab\x21\x95\x66\x17\x0c\x53\x19\xb9\x44\x0f\x4e\x8f\xaa\xe2\x78\xee\xea\x7c\x00\xcd\x19\x0a\x00\xd5\x0f\xb2\x5c\xcd\x51\x09\x34\x6e\x86\xee\xe9\x88\x6d\x0c\x8d\x77\xc5\x29\xf7\x51\xeb\x1a\xd0\xd5\xed\x55\x72\x42\xd0\x69\x50\x6d\xbd\x91\x92\xd7\x81\x0e\x42\x8b\x40\xf3\xa2\x76\xe8\x3b\xb8\xc3\x7b\xa8\xe1\x68\x5f\x45\x13\x30\x6a\x85\x9d\x4c\xf7\x5f\x9b\xef\x6e\xb4\xfb\x17\x00\x00\xff\xff\xa2\xc9\x5b\x8d\x66\x09\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 2406, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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