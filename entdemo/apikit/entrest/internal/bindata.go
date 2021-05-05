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

var _templateServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1b\x5d\x6f\xdb\x38\xf2\xd9\xfe\x15\xb3\x82\x5b\x48\x81\x2b\xf7\xee\xd1\x8b\x1c\x90\x8b\xdd\xae\x81\x5e\x93\xad\xd3\xdd\x87\xa2\xd8\xaa\xd6\x48\xe1\x55\x96\x14\x8a\xee\xc7\x19\xfa\xef\x07\x92\x92\x2c\x89\xa4\x2d\xe7\xa3\x75\xb6\xd6\x4b\x62\x72\x66\x38\x33\x9c\x0f\xce\x50\x5a\xaf\xe1\x0b\x61\xd7\x30\x70\x7f\x43\xcf\x47\x0a\x79\xbe\x5e\x83\x2b\xff\x60\x94\x21\xe4\xf9\x68\x04\xe7\x89\x8f\x10\x62\x8c\xd4\x63\xe8\xc3\xc7\x6f\x80\x31\x5b\x0c\x61\x72\x01\xaf\x2f\xae\x60\x3a\x99\x5d\xb9\x1c\x21\xf6\x21\xcf\xfb\xfd\xd4\x5b\x7c\xf2\x42\xe4\x40\x14\x33\xd6\xef\xaf\xd7\x30\x58\x78\x4b\x8c\x5e\x7b\x4b\x84\xf1\x29\x88\x1f\x30\x70\xc5\x6f\x8e\x42\x96\x69\x42\x19\xd8\x7d\x00\x00\x0b\x63\xe6\xe3\x32\x19\x15\x04\x2c\x39\xca\xa9\xb8\xe7\x49\x1c\x90\xd0\xbd\x2c\x96\xc8\xf3\xad\x93\x23\x31\xdc\x15\x38\xa5\xe8\x93\x85\xc7\xd0\xea\x4b\xb8\x90\xb0\xeb\xd5\x47\x77\x91\x2c\x47\x61\x12\x06\xa3\x30\x18\x05\xd4\x5b\xe2\x28\xb4\x8c\x00\x31\xb2\x51\x78\xcd\x58\x6a\x06\x61\xf8\x95\x8d\xc2\x8c\xd1\x62\x9d\xf5\xfa\x19\x0c\xbc\x94\x9c\xc5\x71\xc2\x3c\x46\x92\x98\xeb\x68\xe0\x6e\x7e\x67\xee\x34\x66\x6f\xa6\xf3\xab\xb3\xcb\x19\xd7\x57\x89\x45\xbd\x38\x44\x18\x2c\x05\x7c\x83\x82\xfb\x1f\xe2\xfb\x11\x7e\xf1\x28\x66\x25\x46\x25\xfb\xd2\xbd\xfc\x14\x5e\x7a\xec\xba\x52\x09\x27\x56\x6c\x9f\xd3\xef\x8f\x46\x20\x54\x54\x6c\x0f\xb8\x7d\xf6\x2d\xc5\xe6\x58\xc6\xe8\x6a\xc1\x60\x2d\xd0\x67\x13\x39\x39\x9b\xb8\x57\x1c\x32\xcf\x61\x63\x5a\x75\x31\x5e\x10\x8c\xfc\xcc\x9d\x0b\xe4\x2b\x2f\x74\x09\x5f\xf3\x43\x61\x74\x1f\x36\x56\xf7\xa1\xa4\x57\x81\x96\xf3\x92\xcb\x96\x06\x02\xa9\x31\x49\xbe\x2e\xaf\xd0\x2d\xf3\x42\x31\x1f\x34\xa8\x55\x1c\x32\x2f\xcc\x54\x8d\xb7\x59\xad\x61\x90\xd8\xc7\xaf\x05\xde\x20\x28\x55\xc2\x39\xe6\x2b\x9d\x56\x0e\x24\x58\x6d\xf2\x2c\x79\xda\xb0\x22\x56\x29\xf4\x45\x02\x41\x8d\x44\x91\xf7\x31\xc2\x3f\xbc\x68\xc5\xc9\x9e\xd4\x29\x0d\x82\xba\x82\x49\x00\x71\xc2\x04\x31\x8c\x33\xc2\xc8\xe7\x4a\x75\x6c\xa3\xb0\x42\xa1\xff\xcd\x92\x78\x6c\x3d\xb3\x34\x4a\xd4\xeb\x14\xa5\x4e\xa6\x7e\x88\x2d\x95\x42\xea\x65\x0b\x2f\x82\x01\xd6\x64\x2f\x99\x41\xf7\x6d\x4c\x6e\x56\x98\xe7\xf0\xee\x7d\xb5\xd6\x09\xc6\x8c\xc7\x88\x01\x0a\x01\x2a\x33\xe2\xdc\xa4\x94\xc4\x2c\x80\x82\xc3\x27\xd9\x30\x59\x12\x86\xcb\x94\x7d\xb3\x3e\xd4\xd6\xf8\xd0\xe6\x37\x17\x96\xfa\x8a\x64\xac\x6e\x99\x6f\xf0\x66\x85\x19\x2b\x8d\xd6\x34\x5d\xd9\x6f\xef\xd2\x0b\x71\x4e\xfe\x87\x00\x24\x66\x7c\x89\x82\x91\xb4\x18\xaf\xb3\x23\xa1\xaf\x92\x4f\x18\x6b\xa0\xc5\x78\x13\x7c\xce\xc3\x9a\x78\x32\x46\x49\x1c\x96\xe0\x59\x42\x59\x03\xf2\x71\x1b\xf4\x33\x83\x2d\x56\x20\x9b\x38\xd7\x0c\x72\x81\x2e\xca\x95\x4e\xa1\x60\x93\x00\x12\x0a\x36\xc9\x5e\x93\xa8\x4e\xca\xa9\xff\x70\x7f\x5f\x21\xfd\xc6\x5d\xa8\x45\xc1\xe8\x7b\x27\x4d\xcf\xaa\x3b\x90\xc2\x82\x46\x74\xbd\x37\x49\xeb\x7c\x89\xdb\x8c\xd3\x30\xfb\x10\xb1\x15\x3e\x8f\x2d\x8a\x37\x2b\x42\xd1\xb7\x76\x05\x5a\x15\xb8\x2e\xd3\x39\x45\x8f\xe1\x16\xb1\xcc\x00\x0d\xc9\x1e\xb9\xb9\xdf\xd9\x94\xeb\x36\x2b\x54\xc6\x6d\xb6\x0d\x6b\x80\xff\xc3\x8b\x88\xef\x31\x2d\x78\x43\x6f\xa7\x55\x78\x7d\x92\xf1\x6d\x7d\x92\xf1\xa0\xca\x67\xba\x92\xd3\x58\xfd\x56\x6f\xfa\x51\x99\x6c\x97\x97\x4a\xec\xfe\x41\xf3\x6f\x0e\x27\xd0\x35\x39\xab\x31\x16\x75\x86\xc9\x31\x35\x46\xdd\x21\xbc\xbe\x44\xf6\x27\x11\x67\x47\x35\xba\x17\x9e\xda\x31\xa3\xef\x72\x8a\x5d\x56\x79\x0f\x06\x6e\x30\xee\xdb\x9d\x71\x0c\xe7\x1b\x25\x97\xec\x4a\x1a\x6f\x53\x7f\x7b\x80\x35\x03\x1c\x6c\xea\x80\x63\xc0\x6f\x18\xa4\xd8\xc2\x63\xc0\x7f\xe8\x80\xbf\x59\x67\xb6\x5c\xae\xb4\x0a\x3f\x38\x99\x8e\x49\xe0\x98\x04\x46\x30\xc1\x08\xb7\x26\x01\x33\xc0\xc1\x26\x81\xbc\xdf\x0f\x56\xf1\x02\x30\x6e\xd4\x3d\xff\xa4\xd8\x2c\xd2\xed\x66\xc7\x90\x97\x68\x65\xff\xa0\x04\x71\x64\xd5\x56\xe9\x78\xdd\xef\x51\x64\x2b\x1a\xc3\xd3\xfa\xf8\xba\xd2\xf7\x6c\x32\x86\x36\x59\x77\x36\x19\xf6\x7b\xbd\x2e\x79\xc9\x14\x29\x34\x44\x75\x60\x43\xd3\xbe\x77\x74\xe9\x3b\xb9\xf5\xbd\xb8\x76\xc5\x81\xf4\x89\x69\xd9\xcb\x55\x5c\x44\x9b\x34\xea\x48\x5a\xa5\x09\x89\x5d\x15\x74\xb8\x47\x05\xde\x33\x5a\x58\xa6\x98\x58\xa6\xd8\x58\x06\xef\xde\x6b\xcc\xec\xdd\x7b\x8d\xa1\xb5\x68\x71\x3d\x2c\xbd\x4f\x68\xb7\x80\x87\x10\x61\xac\x2e\xe4\x38\xfd\x5e\x90\x50\xf8\x6b\xa8\xe8\x81\x53\x92\xa6\xa0\xb2\xb7\xee\xf7\x34\x4b\x9f\x82\x97\xa6\x18\xfb\xb6\x32\x35\xdc\xdf\xcf\x38\x67\x79\xe5\x49\x0a\xc5\x4a\xbd\xed\x99\xd9\x24\xb3\x79\x70\xcf\xda\xea\xe2\xfa\x2b\x3a\x5f\xeb\x7e\x8f\xf8\x75\x55\xc9\x71\xa9\x23\x81\x5c\xd3\x0b\xff\xbd\xd1\x85\x24\xcd\xe5\xe7\x14\x2a\x89\x89\x9f\x49\x48\x77\x36\x69\x30\x4e\xfc\x0d\xab\x61\xb3\xc7\x72\x9e\x44\xab\x65\xfc\x22\xa1\xaf\x48\xc6\x8a\x80\x69\x07\xc2\x4b\x25\x3f\x0e\x6c\xf8\xcd\xbe\x10\xb6\xb8\x06\x39\xbd\x5f\xd3\xe2\x91\xf4\xc3\x16\x5e\x86\xf2\x52\xa0\x3a\x16\x5b\x63\xed\x01\xb3\x50\x6d\xeb\x46\x45\x4a\xaf\x8b\x78\xb7\xef\x9c\xf5\x7c\x0c\xbc\x55\xc4\xc6\xc2\xde\xc5\xaa\x96\x55\x77\xee\xd7\xf8\xa5\xbe\xa5\x73\xa4\x9f\xc9\x02\x7f\xf3\x62\x3f\x42\x6a\x2f\x22\x82\x31\x93\x19\xe3\x5c\xfc\x3f\xe4\xf6\x9a\x16\xf3\xc0\x49\xd8\x14\x4e\xc4\x2d\x8d\x5b\x98\x80\x00\x59\x45\x12\x8d\x5b\xb7\xfb\x46\xfc\x76\x1c\x61\x07\x7c\x6f\x42\x97\x2f\x84\xd4\x76\x9a\xd7\x36\x97\x14\x03\xf2\x95\x43\x58\x23\x2f\x25\x56\x5d\x22\x71\x96\x69\xdc\xcb\x14\xd0\xca\x69\xad\xa2\xa3\xdc\xe4\x34\x31\x6a\x7a\x12\xbf\x33\xf7\xdf\x24\xf6\x37\xb7\x3d\xf6\xe6\xaa\xe7\xe5\xf4\x6a\xcc\xf5\xb4\xa1\xad\xb9\x11\x1b\xea\xb2\xd0\x1e\x17\x4a\x35\x49\xbd\xd8\x07\x9b\xdb\x23\xde\xc0\x60\xc9\x35\x2b\x0e\x1c\x96\xe5\x28\x43\x11\xc9\x98\xe5\x38\x60\xc7\x7c\x35\x57\x5c\x30\x72\x38\x43\xf6\x28\x20\xf6\xcb\x06\xe5\x90\x76\xbf\x1d\x58\x37\x48\x7d\xf6\x28\x50\xbc\x31\x5d\x17\x34\x60\x49\x00\x48\xa9\x88\x4e\xee\xa5\x47\x33\xb4\x9f\x52\xbc\x71\x7e\x15\xa3\xbf\x9c\x42\x4c\xa2\x16\x75\xe9\x3f\x95\x09\xda\x74\x08\x4f\x9b\x76\xa6\xc2\xf3\x67\x4a\x69\x42\xb9\xc6\xc6\xe5\x65\xaa\x2b\x86\x2e\x3d\xea\x2d\x91\x21\xe5\x7b\x3f\x34\xa3\x8e\xc5\xbf\x48\xa9\x0a\x93\x3b\x8d\xa1\xc2\x9a\xea\x4a\xaf\xe7\x05\x11\x42\xc4\xcd\xad\x70\xa8\x46\x8a\x94\xf1\xa5\xf4\x8a\x9a\x92\x28\xde\xb8\xe2\xfe\xe3\x97\x53\xb0\x2c\x8d\x4a\xb2\x84\x32\x4e\x64\x28\xfe\xbb\x48\xf9\x02\xa5\x9c\x2f\x91\x71\xdc\x33\x1a\xda\x25\x1d\x47\x21\x40\x02\x19\x9b\xcb\xb3\x48\x97\x58\x5f\x2e\xea\xfc\x5a\xc3\x35\x71\x58\x2c\x52\xb0\x77\xba\x61\xef\x6c\x7e\x6e\x00\x37\xaa\xcf\xbd\xa0\x3e\x52\xdb\x88\xc4\x1f\xae\xdb\xb3\x6c\x61\x57\x9c\x39\xfa\xdd\xe5\x8f\xaa\x0e\xfe\xe4\x55\xdd\xab\x32\x3d\x99\x3e\x1c\xd7\x13\xbc\x1b\xdb\xaa\x81\x6e\xb3\xcf\x2f\xd7\x48\xf1\x32\xf2\x16\x78\x9d\x44\x3e\xd2\xda\xa9\xa2\xba\xc5\x77\x9b\x27\xb1\xe7\x8e\x12\x3b\xba\x9c\xf9\xa1\x7b\x3e\xaf\x83\xdf\x3e\xaf\xb7\x16\xbd\x63\x7e\x2f\x9f\xc2\x1d\xb5\xbd\x0e\x63\xc8\xd2\xb0\xa3\xac\x7b\x91\x9a\x04\xa8\x3f\xca\x7e\x55\x67\xb8\xf6\xcc\x50\x39\x61\xe8\x78\x16\x69\x4d\xc7\x89\x7d\x62\x12\xd3\xd1\x1b\x5e\x5d\x3e\xdd\x15\x81\x09\x76\x90\xa4\xe2\x54\xc0\xff\x0c\x82\xce\x58\xc5\xd1\x46\xc6\xba\xa2\x2d\xc9\x49\x3c\xff\x51\x5a\xac\x18\xaa\x6a\x82\x3b\x69\x50\xdf\x36\x2c\x1f\xb3\xa5\x1b\xfb\x8d\x86\xce\x8c\x76\x98\x04\xa2\x96\x68\x2b\xc3\x81\x7f\xc1\x73\x8d\x79\xeb\xe3\xdd\x9f\x1c\x5b\xa1\xe1\xba\xee\xd6\x84\xc9\x12\xe6\x45\xc3\xf2\x5c\xa0\xa7\x7c\x9e\xac\x62\x66\x53\xf7\x3c\x89\x19\x7e\x65\x76\x4b\x9f\xc5\xb1\xe2\xbb\x1c\x20\x78\x42\xfc\x1e\xe7\x86\x6c\x87\x4a\xb4\x1b\xde\xa1\x15\x52\x07\xbf\x75\x4b\xa4\x4e\x64\x8f\x48\x6b\x6a\x8e\x94\x0f\x9f\xd7\xb5\x0e\x6d\x47\x15\x17\x6e\x65\xfe\xfc\x79\x45\x96\x84\x89\xa3\x51\xf9\xfe\x8a\x86\xfc\x45\x10\x64\xc8\xec\x0a\x4c\xbe\xb8\xf2\x0c\xfe\xe1\xc0\x09\xec\xc0\x3d\x8b\xa2\xbf\xb7\xad\xf2\x42\x46\x1e\x38\xf7\xee\x14\x39\x6d\x42\x7b\x09\x3b\xf1\x98\x37\x96\x78\x2a\xcf\x57\x3c\x90\x8c\xcb\x80\xa2\x4c\xcf\x32\xae\x8e\x31\x30\xba\xc2\xe6\x6c\x4d\xd8\xa2\x4c\x72\xee\xa5\x42\x1c\xad\x89\xff\xa3\xca\xc4\x10\x0f\xb3\x4a\xd4\xbf\xb7\xf3\xb3\x17\x89\x55\xac\xd7\x15\x89\x0a\xed\xa2\x6a\x74\xfb\xbd\x5e\xaf\x27\x93\xae\xf8\xb7\xd7\x3e\xb9\xcc\x26\xd3\xdf\x45\x04\x9b\x4d\x9c\xa1\x00\x29\x90\x0e\x3b\x57\x6c\x69\xa3\x97\xcf\xc3\x64\x0a\xae\x9a\x17\x84\x66\x07\x72\xd4\x78\x89\xdf\x31\x7a\x77\xe8\x70\xdf\x43\xe8\x56\x02\xb7\x36\xf6\xf6\xb6\x04\xde\xcb\x8b\xf9\xc1\xf6\xe6\x16\xe2\xed\xbd\x83\x8c\xbb\xc6\x17\x0b\x7f\xc6\xd0\xbb\x12\xb7\xcf\x45\xd7\x62\x7c\x0a\x11\xc9\x98\x6e\xf7\xf7\xe9\x72\xdc\xad\x6d\xd1\xb6\xb5\x1d\x11\xb2\x7a\xe7\xd1\xe1\x0b\xc9\xcb\xf4\xad\xc1\xbb\x21\x71\x59\xfb\xb6\x86\xb7\xdf\x43\xc0\xde\x35\x65\xd9\xfb\x41\xb0\x23\x8c\x9b\x8b\x39\xfa\xb2\xdd\x5f\xa5\x91\xe8\x40\x89\x82\x6f\xbf\xac\x08\xf5\xcc\xa8\x9b\x2c\x32\xa5\x6e\x0a\x9a\x1b\xbe\x92\x3b\xde\xd0\xce\x8e\x26\x83\xae\x69\xb0\x0a\x20\xcf\x8b\x0c\x5c\xfd\xde\xd2\xdf\xdb\xd1\x06\x30\x88\xb5\xa3\x36\x86\x0e\x49\x0b\x6e\xe9\xc4\xb0\xc5\x91\xcf\xaf\x71\xf1\x69\x52\xee\xa7\x59\xe8\x5d\xde\x0c\xaa\x47\x43\xa3\x59\xdb\xb4\x19\x43\xc3\xe2\x21\x24\x7c\x40\xe1\x4c\xee\xd5\xce\xc0\xfb\x79\x88\xcc\x01\x3a\xf7\xd8\x27\xd8\xc1\x7d\xbc\xd6\xb7\x9d\xd8\xf6\x77\xba\x75\xb8\xea\xdb\x66\x3b\x5b\x82\x73\x64\x25\x8e\xae\x65\x67\x1b\x7b\x79\x7a\x37\xac\x73\xd4\xb5\x1b\x3a\x97\xc5\xd0\xbd\x2e\xbd\xbd\x8d\xb8\x03\xa4\x0b\xf3\xb7\x54\xf9\x03\xa8\xbb\xab\xaa\xef\x51\xcd\x5d\xf4\xb7\x77\x43\xea\xd1\x57\x63\x82\x23\xf4\x43\x99\x2b\xbb\xbd\xda\xd4\x62\xa4\x7a\x27\xb0\xe3\x66\x6e\x56\xcb\xf3\xd9\xa4\xda\xca\xfa\x30\xaf\x7a\xef\x6e\x3c\x67\xbe\xaf\xbe\x97\x38\x9b\x64\xd5\x4b\x43\xba\x19\x95\x19\xc7\x75\xdd\xef\x6d\x57\x73\xef\x33\x1e\x46\x25\x2b\x33\xcf\x03\x15\xb3\x3f\xa6\x0c\x7d\x7b\xd8\xfd\xbf\x95\xf8\xce\xe1\x20\x4b\x51\xe3\x27\x18\xc7\x52\xf4\x11\x96\xa2\xd5\xd7\x18\xc7\x52\xb4\xf1\xdc\xad\x14\x6d\xf7\x71\x5f\xd7\x1b\xb9\x66\xb4\x63\x05\x7b\xac\x60\xf7\x90\xf0\xd1\x55\xb0\x32\x75\x5c\xc4\x58\x1c\xf9\xb4\x27\xbc\x43\x2b\x66\xb7\x7e\xaf\xa6\xc3\x3d\x16\xb3\xf7\x5b\xcc\x76\xf8\x92\xad\x8d\x77\x2c\x70\x8f\x05\x6e\x9b\xa3\x63\x81\x7b\x2c\x70\x4d\x99\x54\x26\xa6\xbf\x53\x81\x3b\x99\xbe\x9a\x5e\x4d\x0f\xba\xc6\xf5\xc5\x67\x9c\x07\x59\xe3\x1a\xbf\x30\xfd\xd9\x6a\xdc\xe2\xbb\xb5\x30\x63\xd4\x9d\xa7\x11\x61\x67\xb1\x7f\x45\xc9\x72\x9e\x7a\x0b\x2c\xce\x70\x43\xb0\x86\x56\x93\xca\x5f\xfb\x9d\x0b\xa5\xba\x75\xd5\xd6\x96\x4a\x4b\xad\xb2\x66\xb1\x4d\xfc\x8c\x07\x32\x55\x4c\x0d\xed\xe9\x57\x5c\x1c\x46\xf8\x91\xf2\xff\x80\xf0\xa3\x0b\x2b\x79\xff\xff\x01\x00\x00\xff\xff\xc0\x47\xd5\x0f\x1d\x51\x00\x00")

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

	info := bindataFileInfo{name: "template/service.tmpl", size: 20765, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
