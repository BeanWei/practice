// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/gf.tmpl
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

var _templateGfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4b\x6f\xdb\xb8\x13\x3f\xc7\x80\xbf\x03\xff\x82\x5b\x48\x81\x43\xe1\x7f\xf5\xa2\x87\x6e\xec\xcd\x1a\xe8\xb6\xde\xd8\xed\xd9\xac\x35\x52\xb8\x91\x28\x99\xa2\xd3\x64\x0d\x7d\xf7\x05\x29\xc9\x7a\x98\x52\xdd\xc6\x4d\x94\x87\x2e\xb6\xc8\x99\xe1\xcc\x8f\xf3\xe0\x50\x11\x59\x5d\x13\x0f\x10\x30\xc1\x21\x16\xfd\x5e\xbf\x47\x83\x28\xe4\x02\x99\xfd\xde\x89\x01\x4c\x38\x10\x84\x36\x30\x61\xd4\xde\x25\xb9\x21\xe9\x4f\x0c\x8f\x8a\xab\xcd\x57\xbc\x0a\x03\xdb\x0b\x3d\xd7\xf6\x5c\xdb\xe5\x24\x00\xdb\x33\x1a\xa6\x19\x08\xdb\xbb\x12\x22\x52\x12\x10\x42\x68\xbb\x3d\x43\x03\x12\xd1\xf7\x8c\x85\x82\x08\x1a\x32\x34\x7a\x87\x06\xb8\x78\x8f\xf1\x84\x89\xcb\xc9\x7c\xf1\x7e\x36\x45\x49\x52\xb0\x71\xc2\x3c\x40\x83\x40\x31\x54\x44\xe0\xbf\xa8\xe3\xf8\xf0\x8d\x70\x88\x77\x2c\xf2\x31\xb6\x5b\x34\x08\xf0\xec\xda\x9b\x11\x71\x85\x92\xc4\x28\xc4\x01\x73\x14\xad\x25\x55\xb3\x6d\x24\x14\x2e\xa5\x85\x20\xd5\x6c\xe2\x78\x25\xa9\xb6\x8d\xa4\x4c\xc0\x1f\x49\x00\x28\x49\xb2\xb7\xc5\x5d\x54\x1f\xba\x04\x37\x1f\x49\xc5\xe6\x0b\xaa\xd5\x3e\xd0\x58\x48\xba\x9c\xe4\x12\xd6\x1b\x88\x05\xc2\xfd\x9e\xb8\x8b\xa0\x71\x3e\x16\x7c\xb3\x12\x68\xdb\xef\x9d\xcc\x88\x07\x73\xfa\x2f\x20\x44\x99\x90\xba\x2d\xff\x89\x43\x36\x32\xa2\x6c\x7c\x18\x06\x54\x40\x10\x89\x3b\x63\x99\x91\x2f\xc2\x6b\x60\x1a\x72\x35\x5e\xa3\x9f\x4b\xef\x50\x4f\x2c\x38\x65\x5e\x4e\x1f\x87\x5c\x54\x49\x6b\x1b\xe4\xa6\xb8\xfd\x41\xc1\x77\xaa\xdb\xa1\x36\x5f\x10\x4f\x11\xb8\x78\xae\x6c\x59\x10\x0f\x25\xc9\x76\x8b\xbe\x51\x71\xa5\xa6\xe3\x7d\x97\x48\x85\x69\x39\x28\x73\xe0\x36\xe3\x1b\xec\x20\x97\xd8\xc9\x95\xde\x21\x9c\xbe\xa5\xe8\xef\xfe\x54\x95\xa2\x2e\x62\xa1\x50\x3a\x01\x8b\xa9\xa0\x37\x50\xa1\x29\xd1\x0d\x48\xd5\x73\x5d\x9d\xeb\x2a\x75\xf7\x24\xe8\xa5\xe0\x39\x10\xbe\xba\x22\x5f\xfd\xfd\x35\x0b\xae\x02\xaf\x5c\xb6\x1c\x94\x92\x5c\xfc\x91\xfa\xbe\x64\xff\x42\xfc\x8d\x94\x71\x5a\x36\x77\xe0\x2a\xe7\x94\x0c\xcb\x1c\x94\x24\x59\xea\x35\xab\x63\xb3\x1b\xf7\x63\xbd\x72\xbf\x5a\x31\x9d\x52\xf5\xb1\xf2\x7b\x1e\x5e\x17\xd0\x1a\x5d\x0d\xd3\x45\x70\x49\xb9\xd3\xb1\x32\x0f\x4f\xc7\x3b\x45\x77\x5e\xda\xea\x9c\x98\x4a\x65\xa4\x51\x38\xfb\xcd\xf0\x5b\xe6\xf2\xca\x7e\xbc\x2c\xf9\x64\xae\xfe\x39\x07\x22\xa0\xcd\x82\x66\x8a\xaa\x11\x4f\x3e\x2e\xef\x11\x6f\x9a\x58\x53\xb0\x35\x86\x9a\x86\xe1\x0b\xf1\xa9\x43\x44\x5b\x68\x9e\xe5\x16\x45\x9c\x32\xe1\xa2\xe5\x9b\x18\xdd\x8c\x8c\x37\xb1\xb1\x4c\x67\x0e\x96\xd7\x14\x82\xc7\x8b\xb4\x94\x41\x93\xeb\xca\x21\x58\x72\xd8\x34\xe5\x9f\x19\x4b\xcd\x06\xb5\xc6\xa7\x2e\x61\x74\xcf\x84\x43\x52\xc9\xe7\xc8\xf9\x4e\x2c\x36\x53\xbc\xc6\x62\x4b\x68\x29\xd8\x5e\x63\xf1\x61\x62\xb1\x58\x69\x1a\x04\x1b\x3d\xec\xdd\x33\xeb\x90\xf8\x1c\x83\x0f\xed\xf1\xd9\x4c\xd1\x8d\x82\xef\x6e\xd8\x0a\x7d\x84\x6f\x65\x0d\xe7\xc0\x6f\xe8\x0a\xfe\x24\xcc\xf1\x81\x9b\x2b\x9f\x02\x13\xe8\x14\x98\xc0\xe7\xea\xff\x10\x71\x88\xa3\x6c\x1e\x49\x11\x26\x47\xa7\xaa\xd9\xc2\x99\x79\x8a\x64\xe3\xa7\x6c\xb2\x8f\xc3\x97\xea\xdd\xb2\x54\xff\xa0\x52\x88\x87\xe5\x4a\xc0\x4d\xab\xde\xa1\xcd\x38\xb8\xf4\x56\xd2\x18\x36\x89\xa8\x51\xd9\x04\x15\x99\x95\x16\x2c\x23\xdf\x4b\x65\x85\xa4\xbd\xae\xad\xc6\x52\xed\x91\x76\x02\xa8\x23\x37\x63\x46\xe2\x15\xf1\xa5\x36\x51\xfa\x6f\xb7\x4d\x12\x5b\xd9\x9e\x14\x6c\x31\xfe\x9d\x32\xa7\xe8\x09\xcd\x52\x43\x78\x31\x59\x8c\x24\xce\x85\x5a\x49\x62\x6f\xb7\x28\x66\xe4\x1a\x0a\xf8\x8d\x61\xd5\x8c\x9f\xe8\x3d\x4b\x50\x11\xe6\x20\x33\xe4\xc8\x84\xb5\x6c\x47\x2f\x61\xad\xfc\xcb\x30\xac\xbd\x21\x9f\xc6\xc2\xb0\x2c\x64\x32\xb9\x1c\x3e\x0f\x9d\x94\xae\x29\x0d\x65\x24\x49\x32\xfc\xf1\x83\xb3\x7c\xb4\x8e\x63\xe5\x11\x91\x3f\x37\x84\x23\x0e\xeb\xa6\xae\xb4\x4a\x4c\x5d\x04\x9c\x4b\xa4\x38\x9e\x11\x1e\x83\xf9\x96\xc3\xda\xfa\x4d\x8d\xfe\xef\x1d\x62\xd4\xaf\xcb\x97\x4f\xc9\x9d\x4d\x3e\x44\x6f\xab\x3e\xab\x61\x90\xcf\x84\xf3\x90\x4b\xe4\x46\xf9\xd5\x06\x56\x43\x33\xc2\x49\x00\x02\xb8\xf4\x84\x61\x0b\xef\x48\xfd\x05\xce\x35\x44\x89\x55\x1d\xdb\x39\x58\x49\xe5\x61\x6e\x6b\x1a\xa1\xb8\x8c\x0e\xde\x17\xf9\xf7\x06\xf8\x9d\x69\x69\x66\x3e\xd0\x80\x0a\x93\xc3\x1a\xe7\x7d\xbd\x8e\xea\x93\xeb\xc6\x20\xcc\x1d\x5d\xda\xd0\x9f\xa1\xff\x5b\xe8\x14\x7d\x8f\xf9\xbd\xef\x9b\x1c\x9f\x87\x4c\xc0\xad\x30\x2d\x4b\xbb\x6d\x0f\xb4\x41\xd2\x93\x8e\xb8\x2f\xf7\xd2\x73\x4c\x04\x19\x65\xdb\xb9\x3f\x3b\x8d\xa5\xae\x23\x24\xf8\x06\x6a\xd3\x65\x45\xf2\x00\xb4\x8e\x98\x85\xec\x2d\x75\x1e\x2f\x15\x79\xd0\xe1\x4c\xa4\xef\xe0\x5f\x6a\x22\xa2\x4e\x6a\x65\x86\x4a\xa5\x64\x26\x89\x69\x50\xc7\xb0\xee\x9b\xba\x2e\x40\x94\xb3\xc7\x10\x51\xe7\x51\x33\xc8\x05\x74\x2c\x81\xec\xa7\x8f\xc6\xfc\x70\xd2\x96\x1c\x66\x9f\xe6\xdd\x3e\xa3\xac\xd4\xdd\x53\x77\x73\x43\xe3\xdd\xd8\x4b\x4d\x0f\x3f\x1e\xec\x29\x84\xda\x83\xca\x41\xb7\x18\x75\x86\xfb\x5f\x9c\x37\x4b\xfb\xce\xa5\x9e\x8e\x79\xbf\x9b\x6d\xe7\x93\xcf\x1c\x44\xce\xa4\x6b\x94\xd5\x91\x4c\x37\xa1\xc3\xb0\xae\x53\xe3\xfd\xba\x46\x89\xe3\x2f\xde\x70\xcb\x71\x28\xcd\x41\x06\xfc\x2c\xf2\xbf\x02\xf5\x83\x11\x3f\x26\xda\x07\xa1\xd8\x78\xdf\xd4\x3c\x37\x27\x37\xd0\x99\x53\x7d\x9a\x36\x9e\x69\x59\xfe\xfc\x04\xce\xec\x1b\x75\x15\xdd\xdd\xd2\xdc\x78\x55\xfe\x52\x4b\xf3\x43\x9c\xdc\x53\xd0\x3f\x31\x98\x8e\x4d\xea\x3c\x8d\x9a\xde\xfe\x71\x40\xc7\xfc\x5a\xd3\xcb\x8b\x1f\xab\xa6\x1f\xf2\xdd\xa0\xce\xf8\x5a\xe7\x9f\x7b\x9d\x4f\x33\xca\xf3\xac\xf3\xe3\xc9\x87\xc9\x62\xd2\xfd\x52\xef\xa8\xaf\x5a\xdd\x2d\xf5\x8d\x5f\xdd\x5e\x4b\xfd\xa1\xa5\xfe\x87\xaa\x7c\x8a\x77\x5b\x95\x9f\xdc\xc2\xaa\x33\x29\x24\x55\xf7\x71\x52\x48\x43\x62\x48\xfa\xbd\xff\x02\x00\x00\xff\xff\x8b\x19\x0c\x56\x28\x2b\x00\x00")

func templateGfTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateGfTmpl,
		"template/gf.tmpl",
	)
}

func templateGfTmpl() (*asset, error) {
	bytes, err := templateGfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/gf.tmpl", size: 11048, mode: os.FileMode(438), modTime: time.Unix(1, 0)}
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
	"template/gf.tmpl": templateGfTmpl,
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
		"gf.tmpl": &bintree{templateGfTmpl, map[string]*bintree{}},
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
