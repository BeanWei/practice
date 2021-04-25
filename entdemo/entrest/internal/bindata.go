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

var _templateGfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x73\xdb\xb8\x11\x7f\x16\x3f\x05\xca\x51\x32\xa4\x47\xa6\xae\x7d\x54\xc7\x9d\x71\x2d\x5d\xaa\x99\x34\xf1\x45\xce\xdd\x43\x26\x73\x41\xc4\x25\x8d\x9a\x04\x65\x00\x74\x9c\x6a\xf8\xdd\x3b\x00\xff\x98\x20\x41\x89\x72\xe4\x44\x6e\xc4\x17\x8b\xc0\xee\x62\xf7\xc7\xfd\xc3\x05\xe8\x15\x5e\xde\xe0\x10\x10\x50\xc1\x80\x0b\xcb\x5a\xaf\xd1\x70\x89\x63\x88\xde\xe0\x18\xd0\xe4\x0c\xa9\x1b\x34\xf4\xd4\x7d\x96\x29\x82\x28\xf9\x02\xac\x24\x50\x37\x35\x02\x8b\xc4\xab\x84\x09\xe4\x58\x03\x1b\xa8\xf0\x21\x4e\xc6\x85\x78\xdb\x42\x08\x21\x29\x01\xa8\xb8\xbc\x09\x25\xbb\x60\x24\x7e\x47\xc2\x6b\xb1\x10\x52\xca\x62\x79\x0d\x31\x46\xf6\x98\xab\x1f\xb6\x94\x28\x99\xec\x1a\x57\x96\xd9\x86\xb1\xb1\xae\x99\x24\xb2\x06\x76\x48\xc4\x75\xfa\xd9\x5b\x26\xf1\x38\x4c\xc2\x60\x1c\x06\xe3\x80\xe1\x18\xc6\xa1\x6d\x9e\xa5\x20\xc6\xe1\xb5\x10\xab\x62\x0d\x03\x89\x80\x7b\x31\x0e\xb9\x60\xb6\x55\x18\x74\x8a\x86\x78\x45\xce\x29\x4d\x04\x16\x24\xa1\xd2\xb0\xa1\xf7\x70\xcf\xbd\x19\x15\xef\x66\x8b\xab\xf3\xcb\x79\x69\x91\xe4\x62\x98\x86\x80\x86\xb1\xa2\xd7\x24\x78\xff\x26\xbe\x1f\xc1\x17\xcc\x80\x97\x1c\x95\xcd\xb1\x77\x79\x13\x5e\x62\x71\x5d\x41\x21\x85\x01\xf5\x25\xa5\x6b\x59\xe3\xb1\x02\xb9\x7c\x24\xc8\xb3\xc4\xd7\x15\xe8\x63\x5c\xb0\x74\x29\xd0\x5a\xb1\xcf\xa7\xf9\xe4\x7c\xea\x5d\x49\xca\x2c\x93\xf7\x5f\x88\xb8\x6e\x98\xf1\x2b\x81\xc8\xe7\xde\x42\x31\x5f\xe1\xd0\x23\x72\xcd\x4f\xeb\x35\xf2\x8a\xbf\x10\x71\x28\x7e\x2a\x79\x15\x69\x39\x9f\x6b\xd9\x40\x20\xc8\x11\xcb\xc5\xd7\xed\x55\xd8\x0a\xac\x5c\x65\x18\x68\xd2\x2a\x0d\x05\x0e\x79\x1b\xf1\xa6\xaa\x35\x0e\x42\x7d\xb8\x2f\xf8\x86\x41\x09\x89\xd4\x58\xae\x74\xa6\x6c\xa9\x54\xd5\x75\x2e\x1d\xb8\x54\x45\xad\x52\xe0\x45\x02\x25\x8d\x44\x11\xfe\x1c\xc1\xef\x38\x4a\xa5\xd8\x93\xba\xa4\x61\x50\x07\x98\x04\x88\x26\x42\x09\x03\xca\x89\x20\x77\x15\x74\xe2\x01\xb0\x02\xd0\xff\xf0\x84\x4e\xec\x53\xdb\x00\xa2\x19\x53\xc8\x31\x99\xf9\x21\x34\x20\x45\x2b\xcc\x97\x38\x42\x43\xa8\xd9\x5e\x2a\x03\xde\x7b\x4a\x6e\x53\xc8\x32\xf4\xe1\x63\xb5\xd6\x09\x50\xe1\xa9\x78\x53\x06\x54\x6e\x24\xb5\x59\x31\x42\x45\x80\x0a\x0d\x5f\xf0\x51\x12\x13\x01\xf1\x4a\x7c\xb5\x3f\xd5\xd6\xf8\xd4\xd4\x37\x53\x9e\xfa\x9a\x70\x51\xf7\xcc\x77\x70\x9b\x02\x17\xa5\xd3\x76\x4d\x57\xfe\x3b\xb8\xc4\x21\x2c\xc8\x7f\x01\x21\x42\x85\x5c\xa2\x50\x64\x55\x8c\xd7\xd5\xc9\xa9\xaf\x92\x1b\xa0\x06\x6a\x35\xae\x93\x2f\x64\x2a\x53\x17\x17\x8c\xd0\xb0\x24\xe7\x09\x13\x1a\xe5\xf3\x76\xe8\xd3\x0e\x5f\xac\x48\x6a\x64\x43\xac\xe7\xb9\xc0\x94\xe8\xca\xb8\xd0\x04\x98\x85\x78\x0b\xc0\x6c\x79\x2d\x63\xc6\xc4\xb0\xbf\x88\xab\x07\x96\x51\xaf\x06\x2c\xd5\x70\x1e\x81\x06\x96\x27\xd5\xca\xa0\x91\x21\xd8\xb5\x58\x7a\x05\x9b\x42\xa9\x63\xf6\x29\x2a\x01\xba\x9b\xd8\x0c\x6e\x53\xc2\xc0\xb7\xb7\x95\x85\x36\x71\xdd\xa6\x0b\x06\x58\xc0\x06\xb3\xba\x09\x34\xcb\x9e\x79\x70\x3e\x3e\xea\x0c\x11\xa7\x20\xeb\x0a\x38\x03\xfd\xef\x38\x22\x3e\x16\x1b\xe2\xf3\xb4\xb4\xa6\x2c\x06\x2f\xb8\x7c\xac\x2f\xb8\x2c\x01\x72\xa6\xaf\xb8\x8e\x38\xdc\x5f\xbc\xed\xa9\xee\x6e\x8b\x52\x43\xce\x38\x38\xfd\xbb\xd3\x09\xea\xfb\x2a\x71\xda\x72\x4c\x30\x39\xa6\xe4\x34\x38\x75\xc2\x90\x43\xf8\x1b\x12\xd5\xa5\xb8\x9a\xaf\xbc\x02\xf1\x07\x51\x6f\xba\x2d\xd8\xcb\x48\xed\xf9\xfe\xb1\x2d\x28\xb6\x79\xe5\x1e\x1c\xbc\xb3\xc8\x3c\xe6\x8d\xac\xe3\x6d\xac\x55\x4b\xb6\x15\x8d\xf7\x2b\x7f\x73\x82\xed\x26\x38\xd8\xd2\x81\x8e\x09\x5f\x73\x48\xf5\x08\x8f\x09\xff\xa9\x13\xfe\xc3\x3a\xf3\x38\x4e\x8d\x80\x1f\x9c\x4d\xc7\x22\x70\x2c\x02\x63\x34\x85\x08\x36\x16\x81\x6e\x82\x83\x2d\x02\x99\x65\x05\x29\x5d\x22\xa0\x5a\xdf\xf3\x37\x06\xfa\x96\x82\xa3\x6f\x7a\x66\x19\xaa\x76\x3b\x4a\x12\x17\x9d\x68\xbb\x67\x6b\x6b\xc0\x40\xa4\x8c\xa2\x97\xf5\xf1\x75\x85\xf7\x7c\x3a\x41\x4d\xb1\xde\x7c\x3a\xb2\x06\x83\x3e\x75\xa9\x2b\x53\x18\x84\x9a\xc8\x46\x5d\xcf\xbd\x67\x48\x7f\x53\x58\xef\x25\xb4\x2b\x0d\xf2\x98\x98\x95\xbb\xcd\xad\x10\x31\x16\x8d\x3a\x93\x11\x34\x65\xb1\xd7\x26\x1d\xed\xd0\x81\x0f\x3a\x3d\x8c\xb7\x5c\x8c\xb7\x7c\x8c\xa3\x0f\x1f\x0d\x6e\xf6\xe1\xa3\xc1\xd1\x1a\xb2\x24\x0e\x31\xbe\x01\xa7\x41\x3c\x42\x11\xd0\xf6\x42\xae\x6b\x0d\x82\x84\xa1\x3f\x47\x2d\x1c\xa4\xa4\xdc\x15\xda\xea\xad\xad\x81\x61\xe9\x33\x84\x57\x2b\xa0\xbe\xd3\x9a\x1a\xed\x1e\x67\x52\xb3\xac\x8a\xa4\x96\xc4\x0a\xde\xe6\xcc\x7c\xca\x1d\x99\xdc\x79\x13\x2e\x89\x5f\xb1\x4f\xb7\xb6\x06\xc4\xaf\x43\x95\x8f\xe7\x18\x29\xe6\x1a\x2e\xf2\xfe\x01\x8b\x5c\xb4\xb4\x5f\x4a\xa8\x2c\x26\x3e\xcf\x29\xbd\xf9\x54\x53\x9c\xf8\x0f\xaa\xbe\x81\x2f\x75\x85\x16\xc0\xee\xc8\x12\xfe\x85\xa9\x1f\x01\x73\x96\x11\x01\x2a\xf2\xf4\x72\xa1\x7e\x8f\xa4\x71\xab\x62\x1e\x49\x11\x0e\x43\x27\xea\x00\xc2\x2b\x12\xac\x22\x49\xa3\x9c\x4d\x42\xe1\xbd\x53\xf7\xae\x2b\x75\x54\x26\x86\x9e\x5c\x08\x98\xe3\xea\x27\x12\x97\x0c\x02\x72\x2f\x29\xec\x31\x5e\x11\xbb\x9e\xfe\x55\xe1\xd3\x8e\x1c\x0a\xea\x56\x69\xaf\xe4\xb4\x0e\x29\x74\x8e\x5a\x6c\xa8\x7b\xee\xfd\x93\x50\xff\xe1\x20\xc3\x79\x38\xc5\x78\x35\xbb\x9a\x48\x9c\x1e\x64\x9b\x0e\x70\x46\xa6\x9c\xb5\xc3\x61\x49\xcd\x54\x4c\x7d\xe4\xc8\x74\x04\xb7\x68\x18\x4b\x68\x55\x79\xb2\x6d\xb7\x35\x14\x11\x2e\x6c\xd7\x45\x0e\x95\xab\x79\x17\x89\x9f\xd3\x75\xe4\x9a\x82\x62\xb7\xdc\x51\x0e\x19\x1f\xb8\x8b\xd6\x9a\xa8\x3b\xcc\x10\x83\xdb\xae\xad\x70\x8d\x96\x04\x08\x18\x53\xbe\xec\x5d\x62\xc6\xc1\x79\xc9\xe0\xd6\xfd\xbb\x1a\xfd\xcb\x19\xa2\x24\x6a\x48\x97\x57\xcd\x07\x1d\x36\x42\x2f\x75\x47\x6b\xd3\xcb\x6b\xc6\x58\xc2\x24\x62\x93\xf2\xf4\xd0\x53\x43\x97\x98\xe1\x18\x04\x30\xf9\xf0\x47\xdd\xac\x13\xf5\x13\x18\x6b\xd3\x64\xae\x36\x54\xb8\x53\x1d\x74\x2d\x55\x8d\x4a\x93\xf3\xf0\xd2\x32\xaa\xd7\x12\xfe\x5b\x0a\xec\xab\xe3\xb6\x27\x7a\xd6\xc5\x3a\xf9\xa3\xeb\x63\x5d\x48\xbf\x3a\xa9\xd4\xee\xaa\x94\xe5\x25\xe7\x4d\xef\x91\x26\x73\xd1\xb6\xde\xab\x73\xea\x35\x89\x89\x70\x18\xdc\x7a\xe5\xd1\x8b\x41\xfc\xdb\x20\xe0\x20\x9c\x8a\x2c\x3f\x73\x39\x45\x7f\x75\xd1\x09\xda\xc2\x7b\x1e\x45\x0e\xf3\x2e\x12\x2a\xe0\x5e\x38\xae\x6b\xf2\xf0\xef\xe2\xcb\x32\xe2\x9e\xc2\x85\x65\x9e\x9a\x9c\x3d\xea\xb5\xc1\x6d\x0a\xda\xc9\xd8\x29\x16\x78\x92\xf3\xb5\x75\x9e\x73\x69\xef\x04\x09\x96\x82\x3e\x5b\xb3\xa6\x48\x73\xee\x7e\x52\xfc\x78\x4d\xfc\x1f\x95\xe7\x43\x38\xcc\x34\x6f\x3e\xa6\xf9\xd9\xb3\xfc\x23\x93\xfc\x60\x30\x18\xfc\x71\x0d\xd2\x33\xe5\xcf\x41\xd3\x01\xbd\xf9\x74\xf6\x9b\x4a\x52\xf3\xa9\x3b\x52\x34\x05\xd7\x61\x97\x83\x0d\x6d\x53\x79\x3d\x4d\x31\x90\xd0\xfc\x4a\x18\x17\x87\x91\xa1\x5f\xc1\x77\x4c\xd0\x3d\x3a\x9a\x3d\x64\xe7\x56\x6e\x36\x66\xdf\xc1\x86\xd4\x7b\xf9\x76\x71\xb8\xaf\xd7\x4b\x75\x5c\x7b\x90\x99\xb7\xf3\x24\xf9\xe7\x4b\xbe\xa7\x68\x98\xaa\xfd\xc6\x62\x6f\x6a\x72\x86\x64\x5f\x64\x7a\xfc\x7d\xf6\xb2\x50\x47\x86\x7c\xd4\xb7\x23\x95\xb3\x6d\xc9\x91\xd5\x29\xb7\x2b\x17\xca\xb7\x4f\x37\xa6\x6f\xcd\xe2\xb2\xef\x6f\x0c\xb7\xf6\xda\xfa\xbe\x39\x6f\x18\x26\x01\xa2\x80\x9c\x08\xa8\xbe\x98\x8b\x7e\x31\x2d\xe0\xa7\xab\x88\x2c\xb1\x80\x8b\x24\xa5\x62\xb7\xc2\x88\x36\x75\x40\xf2\x2a\x8a\xa5\x69\x0a\xe9\x0f\x3c\xcd\x9f\xb8\x86\xce\x86\x72\x84\x8a\x70\xd6\xaa\xaf\x1c\x48\x03\x94\x65\x45\x11\xae\xee\x5d\xb3\x5b\x6f\x00\xb2\xbc\x3a\xec\x52\x58\x75\xd7\x2c\xd4\xa3\x6e\xa1\x47\x86\x31\xda\x10\xca\x17\xd7\xb0\xbc\x99\x96\x0f\xb4\xdb\xe8\x6d\xf1\x8c\xda\x31\xad\x86\xaa\x73\x29\xb8\x27\x5c\x70\x85\x02\xfa\x07\xfa\xe5\x3b\x99\xf7\x84\x96\x75\x05\x57\xb3\x02\xef\x16\x1f\x79\x09\xd8\xb2\x3d\xb0\x35\xd5\xa1\x7d\x7d\x2d\xd7\x2d\x6c\xf3\x37\x3c\x26\xde\xf6\xe9\xe2\x46\x36\x79\x2d\x40\x94\x3c\xa6\x63\x86\x2a\x64\x9b\x13\x1d\x31\x58\xd7\xa8\xeb\xa3\x3a\x83\x0a\x7b\x5f\xba\x3b\x79\xf4\x20\xe9\xa3\xfc\x23\x21\x7f\x02\xb8\xfb\x42\xbd\x47\x98\xfb\xe0\xb7\xf3\x9e\xd3\xb3\xef\xc6\x94\x46\xe0\x87\x79\xa5\xec\x77\x94\xd5\x50\xa4\x3a\x03\xee\xf9\x30\x1f\x56\xcb\xb2\xf9\xb4\x7a\x94\xf5\x61\xd9\xf5\x7e\xbb\xf3\x9c\xfb\x7e\xfb\x1c\x7a\x3e\xe5\xd5\x21\x91\x69\xa6\xad\x8c\xeb\x79\xde\xf7\xf6\xab\x05\xbe\x83\xc3\xe8\x64\xf3\xca\xf3\x44\xcd\xec\x8f\x69\x43\xdf\x1f\xf8\x0e\x60\xaa\x3e\x6c\x3b\xc8\x56\xb4\xf3\x9b\xbb\x63\x2b\xfa\x1c\x5b\xd1\xea\xfb\xbb\x63\x2b\xaa\x5d\xdf\xd6\x8a\xb6\xb6\x72\xdf\xd4\xf7\x72\xbb\xf9\x8e\x2d\xec\xb1\x85\xed\x6d\xde\xb3\x6b\x61\xf3\xd2\xf1\x96\x42\xf1\xce\x67\x7c\xc5\x3b\xb4\x6e\x76\xe3\x07\xca\x26\xde\x63\x37\xbb\xdf\x6e\xb6\xc7\xa7\xcb\x4d\xbe\x63\x87\x7b\xec\x70\x9b\x1a\x1d\x3b\xdc\x63\x87\xdb\x55\x49\xf3\xc2\xf4\xff\xd4\xe1\x4e\x67\xaf\x67\x57\xb3\xc3\x6e\x72\x7d\xf5\xe1\xfe\x41\x36\xb9\x9d\xff\x53\xf0\xb3\x35\xb9\xc5\x97\xca\x21\x17\xcc\x5b\xac\x22\x22\xce\xa9\x7f\xc5\x48\xbc\x58\xe1\x25\x14\x2f\x71\x23\x64\x8f\x6c\x5d\xca\x9f\xbb\xbd\x18\xe6\x70\x9b\xba\xad\x0d\x9d\x96\xa1\xcb\x9a\x53\x87\xf8\x5c\xa6\xb2\xb6\x9d\x06\xe1\xb3\x7b\x58\x1e\x46\x02\xca\x01\xf8\x01\x09\xc8\x94\x58\x32\xeb\x7f\x01\x00\x00\xff\xff\x3b\x77\x39\x15\x7d\x47\x00\x00")

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

	info := bindataFileInfo{name: "template/gf.tmpl", size: 18301, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
