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

var _templateGfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x6f\xdb\x38\x12\x7f\xb6\x01\x7f\x07\x9e\xe0\x5d\x48\x81\x23\xe3\xf6\xd1\x87\x3c\xe4\x62\x6f\xcf\x40\xaf\xf5\xc6\xe9\xf6\xa1\x28\xae\xac\x35\x92\x79\x91\x28\x87\xa2\x93\xf4\x0c\x7d\xf7\x03\xa9\x3f\xa6\x24\x4a\x56\x12\xa7\x71\x36\xf6\x4b\x22\x72\x66\x38\xf3\xe3\xfc\xa3\xa8\x15\x5e\x5c\x63\x0f\x10\x50\xce\x20\xe2\xbd\x6e\xaf\xbb\xd9\xa0\xfe\x02\x07\xe0\x7f\xc0\x01\xa0\xd1\x19\x92\x0f\xa8\x6f\xcb\xe7\x38\x4e\x28\xfc\xf0\x0e\x58\x46\x21\x1f\x54\x8a\x5e\x97\x04\xab\x90\x71\x64\xf6\xba\x1d\x03\x28\x77\x20\x08\x87\xe9\x22\x46\xaf\x8b\x10\x42\x42\x0a\x50\x3e\xbb\xf6\x84\x08\xce\x48\x70\x49\xbc\x25\x9f\x73\x21\x69\xbe\x58\x42\x80\x91\x31\x8c\xe4\x3f\x86\x94\x2a\xb8\x0c\x85\x2d\x8e\x0d\xdd\xe0\xb0\xa8\x9f\xa4\x12\x6a\x78\x84\x2f\xd7\xdf\xed\x45\x18\x0c\xbd\xd0\x73\x87\x9e\x3b\x74\x19\x0e\x60\xe8\x19\x35\xd3\x14\xf8\xd0\x5b\x72\xbe\x92\x12\x12\xa5\x4f\x51\x1f\xaf\xc8\x39\xa5\x21\xc7\x9c\x84\x54\x28\xdf\xb7\xb7\xcf\x91\x3d\xa1\xfc\x72\x32\xbf\x3a\x9f\x4d\x73\xad\x05\x1b\xc3\xd4\x03\xd4\x0f\x24\x43\x41\x84\xfd\x6f\xe2\x38\x3e\xdc\x61\x06\x51\xce\x92\xdb\x15\xd8\xb3\x6b\x6f\x86\xf9\x72\x6b\xaf\x10\x07\xd4\x91\xb4\x96\x50\x6d\x38\x94\x70\x66\x1b\x80\xec\x5e\x97\xff\x58\x41\x71\x30\xe2\x6c\xbd\xe0\x68\x93\xc8\x98\x8e\x93\xd9\xe9\xd8\xbe\x12\xa4\x71\x2c\x9e\xef\x08\x5f\x96\xcc\xf9\x9d\x80\xef\x44\xf6\x5c\x72\x5f\x61\xcf\x26\x62\xe1\x6f\x9b\x0d\xb2\xd3\xbf\xe0\x47\x90\xfe\x2b\xe5\xe5\xa4\xd9\x7c\xaa\x6a\x09\x09\x37\x81\x2e\x91\x5f\xb0\x5b\xa2\xcc\xb1\x74\x8c\xbe\x5b\x90\x97\xeb\xc8\xb1\x17\x55\xb1\x2f\x2b\xab\x70\x10\xea\xc0\x7d\xca\xd7\x77\x33\x54\x84\xce\x62\xa5\x33\x69\x4d\xae\x6c\x49\xeb\xcc\x5f\x33\x5d\xe4\x32\x29\x64\xc4\x95\xe2\x88\xef\xe3\xef\x3e\xfc\x89\xfd\xb5\x90\x7b\xa2\x8a\xea\xbb\x2a\xc6\xc4\x45\x34\xe4\x52\x18\xd0\x88\x70\x72\x9b\xa3\xc7\xb7\x98\xa5\x98\xfe\x37\x0a\xe9\xc8\x38\x35\x74\x38\xd6\xe0\x0a\x09\x2c\x13\xc7\x03\x0d\xac\xb8\xe8\xb9\xa0\x73\xdd\xcc\xbc\x22\x2b\x71\x51\xc8\x90\x49\xa2\x0f\xc4\x57\xe5\x58\xea\x83\x7d\xc1\x00\x73\x81\x44\x61\xf4\xd3\xca\x49\x47\x55\xa9\xe5\xbd\x5e\x31\x42\xb9\x8b\x52\xa3\x7f\x89\x06\x61\x40\x38\x04\x2b\xfe\xc3\xf8\x26\x54\xcd\xf3\x4b\x59\x80\xd8\x03\x65\xb5\x3f\xb1\x4f\x1c\xcc\xab\xb4\x85\x05\xb7\xeb\xfd\x12\xa1\x5b\xb1\x9e\x58\x45\xcc\xb4\x92\x55\xde\x01\xc5\x4f\x56\x38\x5a\x60\x5f\xd1\x58\xd9\x75\xb0\x3f\x51\x72\xb3\x86\x38\x46\x5f\xbe\xe6\x9b\x7a\x02\x94\xdb\x32\x8b\x49\x4f\xc9\x63\x56\xf5\x8a\xe2\x6e\xd4\x39\x43\x9c\xe6\x83\xf7\x24\xe2\x6a\xf8\x5f\xc2\xcd\x1a\x22\x9e\xa7\x86\xba\xf9\x6d\x96\xe8\xcc\xb0\x07\x73\xf2\x3f\x40\x88\x50\x2e\xd6\x49\x37\x66\x95\x8e\xab\xdb\x93\x92\x5f\x85\xd7\x40\x35\xe4\x72\xbc\x44\x3f\x17\x05\x42\xfe\x22\xce\x08\xf5\x32\xfa\x28\x64\xbc\x48\xfa\xda\x13\xc7\x69\x4d\xcc\xef\x72\xe4\xd4\x8c\x56\x11\xda\x10\x0e\x73\xc0\x6c\xb1\xd4\x46\xdf\x7e\x93\x9b\xde\x5b\x77\x07\xcc\x69\x96\xee\x74\x4c\xcf\xab\x98\x4e\xa9\x36\xe1\xf5\x0e\x1a\xa3\xab\x66\xfa\xc5\x4b\x70\xa6\xbe\xcc\xd2\xd0\x64\x41\x3d\x45\xd1\x88\x57\x1f\x97\x4f\x88\x37\x4d\xac\x6d\xcb\x5f\xcb\xe0\x6c\xac\x55\x68\x8f\xf5\xaa\x29\x04\xf7\x17\x69\xfb\xea\x6f\x76\xc6\xa7\x2e\x61\x1c\x9e\x09\x4d\xa9\x04\x1d\x66\xdb\x76\x6c\xd0\xb4\x0d\xda\x5e\x9b\x33\xd9\x0e\x37\xa6\xdf\x7a\x8a\x63\xfa\x6d\x70\xac\xfa\x73\xc6\x63\x3c\x71\x9f\xde\xd8\xe4\x91\xe8\x20\x73\x57\x9d\xc6\x79\xfa\xdd\xae\x34\x0d\x82\x75\x5d\xf6\x38\x34\xb3\x5e\x4d\x4a\x3e\x9e\x99\x7f\x66\x4a\x1e\x83\x0f\xcd\x29\xb9\x9e\xe2\x30\xda\x7a\x77\x4d\x17\x08\x68\xe1\xe0\xf1\x1b\x83\xe2\x31\xdf\x2c\xbe\xde\x8d\x63\x94\xbf\x7a\xc8\x48\x2c\x74\x52\x78\x6f\xb8\xe9\x75\x3b\x0c\xf8\x9a\x51\xf4\xab\x3a\xb1\xd9\x02\x3c\x1d\x8f\x50\x59\xb0\x3d\x1d\x0f\x7a\xdd\x4e\xa7\x55\x85\xaa\x4b\x15\x1a\xb1\x3a\xb2\x81\x7e\xaf\x3b\x4d\xd0\x44\x15\x6c\xa2\x0a\x38\x11\xfa\xf2\x55\x83\xcf\x97\xaf\x3a\x84\x4a\xc2\x84\xad\x01\xbe\x06\xb3\x44\x3d\x40\x3e\xd0\xea\x4a\x96\xd5\xeb\x76\xdc\x90\xa1\xff\x0c\x2a\x36\x0b\x51\x09\x84\x55\x05\xc5\xda\x9a\xc5\xcf\x10\x5e\xad\x80\x3a\x66\x65\x6a\xf0\x70\x1f\x91\xba\xc5\x5b\x37\xa8\xc8\x54\x61\x2e\x4f\xfe\x56\x5a\x4e\xe3\x81\x25\xef\x2b\x03\x5e\xf0\xc0\xf2\xa4\xb4\xff\x55\xf8\x5f\x05\xb5\x32\x32\x35\xfe\x57\xf1\xbd\x1a\x7c\x1a\xfc\xaf\xcc\xf1\xf2\x3e\xf8\x60\x2f\x79\x92\x0f\x4e\xc7\x91\x29\xaa\xa2\x16\xcf\xf4\xed\xa3\x30\x82\x38\x2a\x6c\xc9\x44\x82\x95\xe4\x56\xf1\x11\x03\x5b\x4c\x12\xe1\x12\x07\x21\x23\xb7\x9c\x38\x51\x42\x6a\x4f\xc7\x25\x03\x88\x53\x50\xf9\x03\xdc\xa9\x8a\xcd\x81\xdd\x92\x05\xfc\x0b\x53\xc7\x07\x66\x2e\x7c\x02\x94\x27\xa1\x71\x21\xff\x97\x10\xae\xd2\x79\x24\x44\x98\x0c\x9d\xc8\xdb\x2a\x3b\xad\x4b\x92\x64\xed\x27\x6c\x02\x12\xfb\x52\x3e\x5b\x96\xd4\x54\x9a\xea\xd9\x62\x25\x60\xa6\x55\xbe\xe2\x9a\x31\x70\xc9\xbd\xa0\x31\x86\x78\x45\x8c\x42\xf5\x94\xcd\x43\xe1\x0e\x2b\x25\xaf\x76\x49\xb9\xa4\xca\xb5\x57\x89\x45\x8d\x1c\x45\x15\xe2\x88\x2a\x3a\x4b\xda\x03\xd1\xf0\xa4\x8d\x42\x56\x5f\x45\x28\x8a\xfd\xdb\xb2\x45\xf6\x3f\x09\x75\xb6\x97\x6a\xa6\x72\xa3\xf6\x6e\x72\x35\x12\x38\x6f\xd5\x4a\x2e\x0c\x23\x8a\xaf\x61\x0b\xbf\x51\x8a\xe8\x47\x5c\xde\x29\x50\x61\xea\x20\x53\x74\x7c\x70\x83\xfa\x81\xd8\x1e\xd9\x18\x18\x86\x55\x19\xf2\x49\xc4\x0d\xcb\x42\x26\x15\xcb\xd9\x17\xa1\x93\xd0\xd5\x1d\x19\x52\x92\x42\x0a\xaa\xeb\xb4\x74\x63\x5a\xc7\xb1\xd0\xa6\x28\xed\x16\x33\xc4\xe0\xa6\xee\xd2\xa0\x48\x4c\x5c\x04\x8c\xc9\xd8\xb0\x67\x98\x45\x60\xfe\xca\xe0\xc6\xfa\x87\x1c\xfd\xdb\x19\xa2\xc4\x2f\xcb\x17\x3f\xc5\x9d\x4d\x36\x90\x89\x5e\xf1\x59\x0d\x83\xf8\x4d\x18\x0b\x99\x40\x6e\x94\x5d\x61\xdb\x72\x68\x86\x19\x0e\x80\x03\x13\x9e\x30\x68\xe0\x1d\xc9\x7f\x81\x31\x0d\x51\x6c\x15\xc7\x72\x07\x53\x37\xa0\x90\x09\x07\x99\xe5\x49\xbc\x16\x72\xae\x5d\x5d\xe0\x8f\x35\xb0\x1f\xa6\xa5\x99\x69\x73\x08\x29\xd3\x3f\xf2\x40\x52\x16\xf3\x3c\xd7\x7c\xea\xef\x33\xe1\x4b\x5d\xc7\xaf\x45\x02\xed\x3c\x37\xd7\xcf\xbd\x27\x01\xe1\x26\x83\x1b\x3b\xbb\xca\xd2\xad\xf0\xd1\x75\x23\xe0\x66\x4e\x97\xdc\x61\x9d\xa2\xbf\x5b\xe8\x04\xed\x62\x3e\xf7\x7d\x93\xd9\x17\x21\xe5\x70\xcf\x4d\xcb\xd2\x86\xc2\x4f\x72\x7a\x11\x9d\xcf\xe6\xeb\x22\xbd\x8d\xce\x1e\xd5\x44\x5b\x15\x49\x0f\xb3\x79\x8c\x39\x1e\x25\x8c\x1a\xcd\xa7\x91\xb0\x7b\x84\x38\x5b\x43\x69\x5a\x35\x2a\x4b\x90\xd6\x1e\xab\xc4\x70\x43\x9c\x97\x2b\x15\x1e\x1c\x70\xa5\xd0\x5f\x80\xbd\xd5\x42\x41\x9c\xc4\xca\x14\x95\x42\x4b\x13\xc7\xa6\x41\x1c\xc3\x6a\x2e\x2d\x8f\xae\x2c\x9d\x4e\xa7\xf3\x79\x09\xd2\xbf\xc5\xff\x9d\xf2\xd7\x51\xf6\x74\x3c\xf9\xc3\x24\x8e\x35\x48\x08\x32\xa6\x63\x09\x2a\xab\xd5\xa6\x04\x09\xe4\x7e\x27\x2c\xe2\x07\x53\x15\xde\xc1\xcf\x2e\x0a\x2d\x5e\x28\xec\xa5\x22\x54\xeb\x41\x6d\xc2\xef\x34\x65\xfb\xd9\xc7\xf9\x61\x1f\x0a\x16\xf2\x2e\xfe\x70\x93\x7d\xed\xb7\x02\x6f\x35\xdf\x0b\xcf\x7c\x58\xc2\x4e\x20\xdc\x75\x16\x68\x7e\x81\xa5\x32\x3c\xfd\x43\xa2\x7a\x69\x3b\x3e\x72\xd0\x31\x57\xaf\x7a\x9a\xf9\xc4\x6f\x0e\x3c\x63\xd2\xbd\x73\x93\xfd\xba\x6e\xa2\x2e\x85\xab\x3a\xd5\x7e\x6f\xa4\x51\x62\xff\x8b\xd7\xd4\x91\xb6\x34\xad\x0c\x78\x2c\xf2\xcf\x81\x7a\x6b\xc4\xf7\x89\x76\x2b\x14\x1f\x71\xa8\x7c\x45\x7d\xd1\x2e\xff\xe8\x83\xe3\x41\xb2\xac\xf2\x62\xad\xe6\xce\x50\xa3\x4c\x7e\x31\xd7\x76\x63\xb7\xeb\xc5\xf1\x74\x9c\x6f\xab\x3a\x6c\x4f\xc7\x7b\x71\xa5\x73\xc7\xa9\xde\x0f\x4e\xc7\x51\xfe\x32\x5a\x37\x53\x55\xc7\xb2\x6d\xfb\x25\xbc\x6c\x8e\x6f\xe1\x60\x5a\xc8\xa4\x38\xed\xb1\x5a\x3e\x49\xd3\x3d\x37\x7f\x9f\x5e\xc1\x51\x7f\x2d\xbf\x06\x3a\xdc\x06\xb0\xf6\x6b\xa5\xb7\xda\x00\x3e\xfc\xc0\xff\xf0\x96\x31\x01\xfd\x23\x85\xe9\x58\x1c\xdf\x5f\x45\xe7\xd8\xee\x74\x8e\x8e\x9d\x63\xcd\xe2\xfb\xea\x1c\xdb\x7c\xba\x55\x66\x3c\x76\x93\x6f\xb1\x9b\x6c\x17\xb1\xc7\x6e\xf2\xd8\x4d\x96\x7f\x75\x15\x38\xa9\x5b\x7f\xcd\x6e\x72\x3c\x79\x3f\xb9\x9a\x1c\x7e\x43\xe9\xc8\x6f\x19\x0f\xb7\xa1\xac\xfd\xd6\xf2\xd8\x50\xb6\x6d\x28\x1f\xd4\x4b\x26\x78\x37\xf5\x92\x93\x7b\x58\x1c\x4c\x0a\x49\xd4\x7d\x99\x14\x52\x93\x18\xe2\x5e\xf7\xff\x01\x00\x00\xff\xff\x34\x37\xc4\x6e\x4d\x41\x00\x00")

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

	info := bindataFileInfo{name: "template/gf.tmpl", size: 16717, mode: os.FileMode(438), modTime: time.Unix(1, 0)}
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
