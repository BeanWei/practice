// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/service.tmpl
package internal

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

var _templateServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\xcd\xea\xdb\x30\x10\xc4\xcf\xcd\x53\x0c\x42\xa7\xc2\xdf\xbe\x17\x72\xe8\xa5\xc7\x50\x68\xe8\xb9\x8a\xb5\xb2\x45\x65\x29\xd8\x9b\x36\x41\xe8\xdd\xcb\xca\xce\x47\xf3\x71\x08\x42\xbb\x33\xfa\xcd\xe0\x9c\x61\xc9\xf9\x48\x50\x96\x0e\xa7\xbe\xad\xff\x0a\xa5\x6c\x36\x39\x7f\xe0\xaf\xe7\x01\x74\x66\x8a\x16\x1a\xea\xbb\xe9\x7e\x9b\x9e\xd4\xba\xad\xf0\x51\xca\xe6\x53\xce\x60\x1a\x8f\xc1\x30\x41\x0d\x64\x2c\x4d\x0a\x8d\x78\xe4\x0c\x51\x2e\x6e\x98\x4c\xec\x09\x3a\xe2\xcb\x16\xba\xd9\x25\x4b\xb3\x8c\x00\xa0\x6d\x91\x33\x74\x6c\x76\x66\x24\x94\x02\x3f\x83\x07\xc2\x98\x2c\x05\x50\x64\xcf\x17\xb8\x34\xd5\xcb\xff\x37\xe7\x6e\xa0\xd1\x34\xd5\x86\x2f\xc7\x97\x31\x4f\xa7\x8e\x91\xeb\x5c\x7e\x12\x6b\x25\x71\x95\x24\x36\xdf\x3c\x05\x7b\x63\x79\x5c\xd4\x6c\xfa\xba\xe4\x9a\x1f\xd5\x68\x6f\x7a\x94\x92\xf3\xd2\x8c\x8c\xe7\x25\xce\xd7\x18\x13\x1b\xf6\x29\xce\xab\xdf\x5b\x85\x8f\x96\xce\xab\x4e\xbb\x2b\xa5\x20\xcb\x4b\xdb\x5a\xdb\xad\xb5\x7b\x7d\xcf\x5c\xf7\xbe\x6b\x43\xad\x93\x17\xbb\x34\x8e\x14\x59\x49\xb0\x17\xcd\x3d\x42\xa5\x93\x6a\x72\x86\x77\x95\xc2\x87\x60\x0e\x81\x7e\x9a\x70\x12\x9c\xcf\x8f\x04\xda\x35\x7b\xa9\xf5\x26\x88\x89\xab\x19\xc5\xd9\xb3\xff\x23\x93\x5f\xd7\x00\xcb\x91\xc2\xbc\xa6\x7a\xe6\x4c\xa3\x67\xc9\xae\xa0\xdf\xe6\x93\x6c\x0f\x57\xcf\x5f\xd0\x72\xfc\x17\x00\x00\xff\xff\x06\x91\x7a\x47\xb5\x02\x00\x00")

func templateServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateServiceTmpl,
		"template/service.tmpl",
	)
}

func templateServiceTmpl() (*asset, error) {
	bytes, err := templateServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/service.tmpl", size: 693, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/service.tmpl": templateServiceTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"service.tmpl": &bintree{templateServiceTmpl, map[string]*bintree{}},
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
