// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package templates generated by go-bindata.// sources:
// scripts/create.tpl.sh
package templates

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _scriptsCreateTplSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x51\x6b\xc2\x30\x14\x46\xdf\xf3\x2b\xae\x99\x0f\xdb\x43\xcc\xb3\x8e\x0e\x02\x2d\x54\xcc\x54\x1a\xcb\xd8\xc6\x28\xb5\xb9\x2e\xc1\xd9\x4a\xd2\x32\x9c\xfa\xdf\x47\xc1\xa2\x63\x7d\xcd\xfd\xce\xe1\x90\xbb\x01\x5f\xdb\x92\xaf\x73\x6f\x80\x61\x43\x48\x38\x4d\x32\xa5\xe2\x80\x72\x53\xed\x90\x7f\xe6\x4e\x63\x89\x8e\x8f\xbc\x37\x94\x2c\xc5\x2a\xce\x44\xba\x8a\x17\xc9\xf4\x2d\x0a\xb3\x59\xf4\xaa\x02\x3a\xbc\x40\x3c\x6f\x6a\x53\x39\xfb\x83\x3a\xdb\xe2\xc1\x5f\xf6\x2a\x0d\x17\x51\xa2\x02\xca\xb1\x2e\xb8\x6f\x74\x85\xce\x8f\x34\x1f\x8f\x59\xa7\x67\x8d\x47\x47\x49\xaa\xa2\x64\x2e\x9e\xa3\x80\x76\x07\x4a\x88\xd5\x30\xec\x0e\x70\x3a\x41\x3b\xcd\xf5\xed\x23\xdb\xa5\xc4\x6e\xe0\x1d\x06\xc0\x36\x40\x87\x7d\x95\x14\x3e\x1e\xa1\x36\x58\x12\x80\xdd\x56\x5b\x07\x6c\x0f\x5d\x38\x01\x28\xf6\x2d\x7b\x3c\xc2\x68\x9f\xd7\x66\xd9\xac\xbf\x6c\xa1\x54\x3c\xc3\x03\x9c\xcf\xd0\xab\x6c\x29\x53\x7d\x97\xd7\x92\xc9\xb5\xa9\x9f\xd8\xd8\xff\xa1\x97\xef\xb9\x0d\xc4\xc2\x54\x40\xaf\x32\x21\x65\x70\x2f\xa4\x7c\x80\xf9\x62\x29\x94\x7a\x09\x27\x42\x4a\x0a\x4f\xf0\x47\xd1\xea\x7f\x03\x00\x00\xff\xff\xf0\x0d\x0f\x58\xd1\x01\x00\x00")

func scriptsCreateTplShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsCreateTplSh,
		"scripts/create.tpl.sh",
	)
}

func scriptsCreateTplSh() (*asset, error) {
	bytes, err := scriptsCreateTplShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/create.tpl.sh", size: 465, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"scripts/create.tpl.sh": scriptsCreateTplSh,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"scripts": &bintree{nil, map[string]*bintree{
		"create.tpl.sh": &bintree{scriptsCreateTplSh, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
