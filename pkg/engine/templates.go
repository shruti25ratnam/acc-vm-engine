// Code generated for package engine by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../parts/tvm/base.t
// ../../parts/tvm/outputs.t
// ../../parts/tvm/params.t
// ../../parts/tvm/resources.t
// ../../parts/tvm/vars.t
package engine

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

// Mode return file modify time
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

var _tvmBaseT = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x4d\x4a\x04\x31\x10\x85\xf7\x7d\x8a\x22\xba\x74\x92\x6e\xc1\x4d\x9f\x43\xdc\x88\x8b\x32\x16\xce\x48\xe7\x87\xaa\x4a\x83\x36\xb9\xbb\x24\xd3\xa8\xe0\x0c\xd9\xe5\xbd\xef\x7d\xb5\x0d\x00\xe6\x56\xfc\x91\x02\x9a\x19\xcc\x51\x35\xcb\xec\xdc\xf9\xc7\x06\x8c\xf8\x4e\x81\xa2\x5a\xfc\x2a\x4c\xd6\xa7\xb0\x67\xe2\xee\xc7\xe9\xe1\x30\x4e\x87\x71\x72\x6f\x94\x97\xf4\xd9\x7a\x8f\x14\xf2\x82\x4a\xf6\x43\x52\xbc\x31\x77\x6d\xdf\xa7\xa8\x14\xf5\x89\x58\x4e\x29\x36\xcd\x64\xc7\xf6\xce\x71\x46\xc6\x40\x4a\x2c\x66\x86\x76\x10\xc0\xb6\xe9\xbe\x03\x46\xd7\xe0\x7a\x45\xac\x1a\xb0\xb5\x0e\x00\xb5\x83\x2b\xf2\x09\x5f\x17\xba\xce\xad\xc8\xff\x28\x26\x49\x85\x7d\xa7\x9e\x2f\x52\x3f\x8d\x5f\xf4\xa5\xa3\xa9\x68\x2e\x7a\x5d\xb7\xe7\x7f\x8c\x43\x1d\xbe\x03\x00\x00\xff\xff\x7a\x3e\x98\xca\x63\x01\x00\x00")

func tvmBaseTBytes() ([]byte, error) {
	return bindataRead(
		_tvmBaseT,
		"tvm/base.t",
	)
}

func tvmBaseT() (*asset, error) {
	bytes, err := tvmBaseTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tvm/base.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tvmOutputsT = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x50\x50\x4a\x4c\xc9\xcd\xcc\x0b\x2d\x4e\x2d\xca\x4b\xcc\x4d\x55\xb2\x52\xa8\xe6\x52\x50\x50\x50\x50\x2a\xa9\x2c\x00\xf1\x94\x8a\x4b\x8a\x32\xf3\xd2\x95\x74\x20\xa2\x65\x89\x39\xa5\x60\xe1\xe8\x82\xc4\xa2\xc4\xdc\xd4\x92\xd4\xa2\x62\x0d\x75\x14\x23\xd4\x35\x63\x95\xb8\x14\x14\x6a\xb9\x00\x01\x00\x00\xff\xff\xe6\xc7\xec\x89\x5c\x00\x00\x00")

func tvmOutputsTBytes() ([]byte, error) {
	return bindataRead(
		_tvmOutputsT,
		"tvm/outputs.t",
	)
}

func tvmOutputsT() (*asset, error) {
	bytes, err := tvmOutputsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tvm/outputs.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tvmParamsT = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x98\x4d\x6f\xe3\x36\x10\x86\xef\xfd\x15\x03\x5d\x36\x01\xdc\x1c\x16\x45\x0f\xb9\xa5\xc8\x22\x28\xe2\x7c\x20\x8a\x7d\x09\x82\x62\x22\x8d\x2d\x22\x12\x29\x70\x86\x76\xb2\x86\xff\x7b\x41\xc9\xb2\x65\xaf\x3f\x24\x39\xf5\xa1\x39\x05\x14\x39\xf3\xbc\x2f\x47\x43\xca\x00\x00\x41\x6a\x22\x14\x65\x74\x70\x09\xb3\xdf\xa0\xf8\x0b\xe4\x33\xa7\xe0\x12\x02\x16\xab\xf4\x38\xe8\x55\xe3\x31\x8d\xd0\xa5\x32\xc4\xd4\x15\xcf\x5f\x2c\xb1\x71\x36\xa2\x1b\x6b\x5c\x7e\x76\x7e\x51\x05\x7b\x5d\xae\x99\xcd\x6e\x48\xae\xd2\xd4\x4c\x29\xee\x2f\x9e\xf2\x7c\x5e\x45\xcc\x48\x30\x46\xc1\x5a\xf6\x22\x0f\x47\x56\xe5\x0b\xac\xa0\x5a\x07\x23\x63\x01\xd3\x14\xaa\xb4\x7c\x11\x2c\x16\x95\x01\xe7\x65\xd6\x60\x92\xdd\x63\x46\x4d\x14\x35\xca\xef\x83\x81\x19\x81\x24\x04\xc3\xbb\x9d\x39\x43\xf5\xb3\x49\xce\xba\x23\xc3\x3b\xbf\xa8\xad\x1f\x7e\xcd\x41\x1e\x74\x92\x90\x16\x55\x5a\xf7\x5c\xb2\xb4\xde\xe1\x1c\x99\xa7\xc6\xc6\xab\x19\xb8\x20\xf7\x33\x38\xb8\x84\x97\x15\xe6\x2f\x93\x01\x02\xe6\xe4\xd1\xbd\xa5\x2a\xba\xa5\xcf\x0a\xf4\xb5\x9d\xfd\x9e\xdd\xcb\x5d\x57\x04\x62\xc0\x31\x81\xd1\xd0\x57\xda\x7d\xc0\x44\x59\x71\x98\x42\x86\x51\xa2\x34\xed\xb2\x25\xce\x94\x1e\x30\x59\xdd\xb0\x42\x36\x1d\xc1\x9f\xce\x92\x63\xb2\x2d\x8b\xa8\xca\x59\x14\x71\xb1\x73\x0b\xde\xbb\xc3\xbc\x8f\x95\xb1\xdb\x78\x29\x72\x96\x0e\x51\xb7\x84\xad\x12\xb6\x84\x35\xfc\x77\x86\x63\x2a\x36\x9c\x13\xb2\x5d\xfc\x6d\x49\xfa\x10\x82\xf2\x39\x21\xaf\x92\xee\x67\x7b\x18\x8d\x4e\xcb\x65\x7c\xc2\xfd\x4c\xe1\xed\xe0\x94\x44\xe1\xed\x60\x3f\xcf\x90\x2c\x77\x3c\x11\x52\x14\x62\xe9\x4a\x36\x29\x13\x1f\x70\xab\x68\xfd\x03\x9b\x9e\x9e\x6f\xf0\xd4\xdf\xc9\x76\xad\xf8\xbd\x61\x97\x2d\x4e\x80\x87\xe5\x8a\xb6\xed\xbf\xea\x87\x19\x6a\x1c\x53\x0c\xb1\xe2\x77\xdf\x0d\x23\x4b\x28\xbb\x5e\xcd\x89\x26\xb9\xa7\xe9\x83\xfd\xf1\xa1\x58\x3c\x4d\x07\xf7\x34\x4d\x1b\x1d\x04\xf5\x79\x00\x01\x55\x29\xbb\xf5\xff\x6b\x12\xb2\x99\xd2\xc4\x30\x4d\x48\x12\xb2\x60\x2c\x68\x23\x80\xa0\x69\xba\x6c\xfc\x9a\x64\x6a\xec\x3b\x70\x62\x5c\x1a\xc3\x1b\x41\x6e\xcd\x44\xf9\x8a\xa2\x78\x8f\x2b\x1d\x0f\x82\x97\xc8\xe8\x08\xe5\x6c\xf3\x0e\xe4\x9b\x7c\x0f\xbe\xfd\xee\x63\x7f\x3b\x7f\x3d\xe2\xae\xb1\x29\xec\x0c\xd3\x3c\x41\xed\x32\xb2\x2a\xea\x41\xf2\x99\x27\xa4\x7b\xe0\x74\x4c\x96\x23\x63\xa9\x07\x39\x59\x65\xe2\xf3\x55\x11\xac\xb2\x2b\xdd\x27\x3d\x96\x24\xb8\x84\xef\xab\x51\xfc\x58\x8e\xfe\xf9\xc7\x2f\xde\x3c\xd5\xb5\x75\x36\x6a\x9b\x43\xc7\xf8\x52\xc5\x83\xb1\x0f\xb8\x3c\xa3\xaa\x32\xdb\xf4\x6d\xcf\x1b\x71\x15\xc7\x96\x98\xbb\xa8\x9a\xcd\x2e\x86\x9a\xe4\xd1\x9a\x91\x4a\xa9\xf8\x7f\x11\x6d\x3e\x6f\xa9\x6d\x78\xff\xe3\x19\xb0\x5c\x0c\x9c\x63\x44\xdb\x89\xd9\xbd\xfd\x87\xf5\x5a\x46\x3f\xae\x62\xcb\x18\x3b\x0c\x2f\x1f\x7e\x9d\xe5\x61\x3d\x5e\x6b\xd3\x43\x12\xae\x31\x1f\xbc\x4f\x97\xf7\xac\xbf\x8c\x91\x2e\xe8\x62\x1d\x35\x6a\x9d\x6b\x13\xfd\xc0\x08\x53\xa6\x8e\x9d\x33\x2c\x98\xc1\x43\x03\x93\x14\x2f\xc7\xc1\xef\x98\xe7\xc7\xbb\x53\x2a\x3c\x4a\xa0\x87\x6d\xac\xec\xcd\x18\xb9\x56\x38\xd6\x86\x45\x45\x9d\x2a\xb0\x64\x3d\xb5\xca\x62\xff\xe2\x15\x79\x63\xc5\xb5\x35\xa1\x18\x8b\x63\xba\x8a\x22\xe3\xf4\xff\xf5\x26\xc0\xa5\x48\xc0\x52\xe5\xf6\x9b\x40\x6b\xaf\x3a\xf6\x5b\x6d\x34\x1d\xd1\x4a\xd7\xf6\x7b\x5d\x57\x5b\x05\x5d\x3f\xfe\x43\x41\x1d\xa3\x8d\xff\xe9\x3f\x85\x8d\x76\x7b\xeb\x82\xfa\xf8\xcd\x53\x78\xe4\x8f\x00\x5f\x68\xcb\xad\xd2\xdb\xbf\xa8\x0f\xda\x52\x44\x69\xe8\xc8\xfa\xdc\xd5\xd0\xf0\x7b\x47\x27\x3c\x76\x4b\x27\xfe\x0d\x00\x00\xff\xff\x8b\x18\xaa\x17\xe9\x13\x00\x00")

func tvmParamsTBytes() ([]byte, error) {
	return bindataRead(
		_tvmParamsT,
		"tvm/params.t",
	)
}

func tvmParamsT() (*asset, error) {
	bytes, err := tvmParamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tvm/params.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tvmResourcesT = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x4b\x6f\xe3\x36\x10\xbe\xe7\x57\x10\x6a\x01\x29\x80\x2c\x3b\x01\xfa\xd8\xdc\xb6\x9b\x60\x6b\xb4\xce\x1a\xab\x34\x97\x20\x07\x9a\x1a\xcb\xdc\x48\xa4\x96\xa4\x9c\xf5\x1a\xfe\xef\x05\xf5\x7e\x50\x96\xdd\x6c\x93\x83\x6d\x91\x33\x1c\x7e\xdf\xcc\x37\x63\x23\x84\xd0\xfe\x02\x65\x7f\x16\x4e\xe8\x23\x08\x49\x39\xb3\x6e\x90\x75\x3d\xbb\xfa\x7d\x32\xfb\x65\x32\xbb\xb2\xdc\x72\x07\xc3\x31\xe8\xb5\x27\xc2\x19\xc1\xca\xb1\x3f\x83\xe4\xa9\x20\xf0\x51\xf0\x34\xb9\x85\x24\xe2\xbb\x18\x98\x9a\xd8\x2e\x4a\x19\xfd\x9a\x82\xaf\x04\x65\xa1\x13\x54\x4b\xce\xa5\xa7\xbd\x5c\x5e\x3e\xd7\x6e\xd5\x2e\xc9\xdc\x2e\x28\x11\x5c\xf2\xb5\xf2\x4a\xbf\x72\x5a\x5b\xca\xda\x20\x11\x3c\x01\xa1\x28\x48\xeb\xa6\x8a\x1f\x21\x2b\xe6\x41\xe6\x68\xce\x88\x00\x6d\x83\xa3\xca\x48\x9f\x03\x71\x12\x61\x05\x2d\x23\x84\xac\x9f\x25\xd9\x40\x8c\xb5\xe5\x46\xa9\x44\xde\x4c\xa7\xf9\x13\x2f\xc6\x0c\x87\x99\x27\x0f\x7f\x4f\x05\x78\x84\xc7\xc5\x9a\x9c\x5e\xcf\xae\x34\x3c\x93\xd9\x55\x23\xca\x87\xe2\x0c\xef\x8b\xe4\xec\xa7\xc6\xe9\x08\x59\x84\x33\x05\x4c\x35\x40\xbe\xf2\x66\xfa\xbf\xbd\x4d\x94\x97\xb7\x6e\xd0\xd3\x73\xb5\x72\xb8\x68\xbe\x1e\x72\x93\xfd\x30\x86\xf7\xa0\x5e\xb9\x78\x99\x26\xe9\x2a\xa2\x64\xbe\x7c\x1f\x04\x02\xa4\x84\x06\x8e\x3d\xc6\xdf\x4d\x66\xd7\x66\xc6\xb7\x58\x50\xbc\x8a\x40\x3a\x76\xc7\xe1\x3d\x8e\xc1\x6e\xd2\x19\x71\x82\x55\xe1\xf3\x29\xc1\x02\xc7\xa0\x40\x48\xc7\x2e\x17\x5a\xbb\x3b\x5c\xee\xf7\x74\x8d\xfe\xc4\xf2\xf6\xde\xd7\x8e\x91\x77\x38\xd4\x04\x06\x4c\xfa\xa0\x14\x65\xa1\xec\x72\x18\xf0\x18\x53\xa6\x4d\xfe\xc6\x2b\x88\x7a\x67\x6f\xe3\x32\xce\x1a\x50\xf7\x62\xbf\x07\x16\x34\x4f\xa8\xee\x16\x95\xc1\x2e\x40\x6d\x78\xa0\xfd\xdd\xee\x18\x8e\x29\xb1\xfe\x23\x0f\x2c\x7f\xf5\x81\xa4\x82\xaa\x5d\x56\x32\x6f\xe6\x82\xc9\xf0\xc7\xe2\x5f\x21\x21\x8b\x38\x3f\xa7\x51\x9e\x89\x0d\xb4\xf7\xfb\x8f\xa0\xfc\xe6\x06\xe4\x3d\x2e\x96\x82\xaf\x69\x04\xde\x92\x0b\x25\x1b\xa0\x3e\x1f\x05\x8c\x70\x16\xd0\x2a\x5c\xf8\x9a\xe2\x48\x3a\x2d\xe6\x18\xa8\x7b\x78\xfd\x24\xee\xbe\x51\xa9\xb9\xb7\x2f\x5d\x64\x33\x78\xb5\x8f\x6b\x48\x89\xfb\x96\x0a\x95\xe2\xa8\xf8\x38\x82\xf8\x3b\x33\xe2\xbd\x78\xfe\x1f\xcc\x71\x5e\x51\x7e\x82\x49\x4f\xa4\x8a\xb5\xa5\x80\x35\xfd\xd6\x63\x04\x19\x82\x2c\xea\xb3\x95\xf3\x35\x1b\x15\x13\x39\xdb\xe9\x8a\x81\xea\xf1\xdc\x3e\xc1\x08\x47\x6e\xd9\x05\xe4\xf8\x45\x4d\x57\x1a\x70\x6c\xbe\x44\xad\x86\xed\xf7\xc7\x53\x6d\xb4\x36\xe7\x4c\x81\x58\x63\x32\xa6\x91\xbf\x8d\xd7\x25\x25\x6f\xcc\x91\x00\x12\x60\x81\xfc\xc4\x5a\x9c\x9c\x23\xc4\xa6\x9c\x30\xed\x31\xcb\x49\xb1\xe1\x79\x34\x69\x69\xf2\x81\xb3\x35\x0d\x53\x91\x5d\xe2\xc4\x1c\x2a\xad\xee\x75\xcb\x3e\x2b\x6b\x12\x41\xb7\x58\xc1\x71\x8d\x76\xbb\x56\x79\x36\x19\xfc\xe9\x1b\x04\x5d\xfa\x34\x54\x7e\x66\x31\x0f\x7a\xa9\xd7\xaa\x9c\x32\xa8\x9c\x8a\xa4\xa0\xe2\xe8\x39\x65\x7f\x9f\x07\x8e\x7d\x42\xbb\xb6\xdd\x31\xca\x0d\x01\x9e\x51\x2b\x08\xb9\x96\xa9\x37\x75\x05\xa8\x8f\x12\x93\x61\x07\x9e\xe3\x33\x8a\x51\xea\x1b\xfe\x02\x8a\x43\xc6\xa5\xa2\x44\xfa\x8a\x0b\x1c\xc2\x7b\x92\xd7\xc5\x49\x82\x5f\xd8\x4c\x65\x69\x4b\x78\xda\x1a\x1b\x4d\xa5\xfc\xeb\xb8\xe0\x9b\xc2\xca\x5c\xbf\xb1\xc2\x5f\x28\x0b\x4e\x3f\xec\x2f\xca\x82\x96\xb9\x7c\x49\xdb\x85\x78\x5e\xf0\x0f\xbb\xa4\x59\xe7\xa7\x2a\xe6\x07\x1e\x27\xa9\x82\xb2\xab\x2e\x30\xd9\x50\x76\x4c\x2f\xaf\x67\x27\x81\x5c\xcf\x67\x3f\x44\x2f\xcb\x31\xf2\x43\x2a\x15\x8f\xe7\x31\x0e\xa1\x39\xe6\x35\x1e\x5b\x86\x29\xb0\xfa\x7e\x73\x42\xa7\x98\xda\x2e\x32\x09\xbf\x49\x42\x15\xd6\x83\x6b\xf1\xa9\xc1\x1c\x11\x90\x5d\xc8\xcf\x74\xc1\x42\xfa\xc2\x36\x26\x64\xb2\x8d\x27\xc0\x42\xca\xc0\xae\x89\x1a\x15\xe4\x0d\x16\xc1\x2b\x16\x50\xcc\x63\xdd\x3a\xde\xc6\x3e\xfd\x6e\x22\x40\x3f\xee\x0e\xc8\x95\x55\x39\x0e\x0e\x38\x4d\x61\x4d\x07\x46\xf3\xd2\x18\xfe\xe0\x5c\xdd\x31\x8d\x53\x3f\xed\xeb\x1d\x86\x29\x62\xfb\xb0\x5c\x0c\x19\x3e\x3e\x2c\x17\x1d\x85\x3e\x18\x2f\xc0\x65\x1d\x7a\x4b\xc5\xaa\x85\x4e\x5f\x2c\x64\x64\xc0\xaa\xbd\xda\x31\x2d\x92\x64\x00\xab\xfe\xb0\xd1\x9d\xe5\x7a\x3d\xef\xa4\xce\xd1\xf3\x3b\x9a\x99\x5d\xbc\x86\xa6\xc3\x86\x88\x0c\xdc\x69\xc5\xb9\xba\xad\x77\xf5\x53\x00\x1a\xfc\x19\x86\xfc\x8e\x7d\xa6\xf8\x4a\xa4\xa6\x99\xb2\x40\xfe\x1f\x41\x33\x67\x74\xed\x9c\xe5\xcf\x45\x02\xd6\x20\x80\x11\x70\x1a\x78\x9e\xa4\x99\xad\x9f\x3b\x72\x30\x5d\x64\x8f\x76\x20\xdb\x45\x67\xf4\x13\xed\x32\xff\xf1\x45\x7f\xfd\xb3\x2f\xbd\x44\xd0\x18\x8b\xdd\x1d\x0b\x12\x4e\x99\x92\x4f\xf6\x2a\xe2\x2b\xfb\xd9\x45\x5f\x24\x67\x8e\xcd\xd2\x28\xea\xb2\x7a\x18\x6a\xc9\x17\xff\x06\x00\x00\xff\xff\x02\x6a\xeb\x05\xf3\x11\x00\x00")

func tvmResourcesTBytes() ([]byte, error) {
	return bindataRead(
		_tvmResourcesT,
		"tvm/resources.t",
	)
}

func tvmResourcesT() (*asset, error) {
	bytes, err := tvmResourcesTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tvm/resources.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tvmVarsT = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x41\x6f\xe2\x3c\x10\xbd\xf7\x57\x58\xb9\x18\x24\xbe\x72\xef\x0d\xb5\xdf\xd7\x0f\x75\xb7\x65\xcb\xb6\x7b\xa8\x7a\x30\xce\x24\xcc\x36\x19\x67\x3d\x36\xb4\x42\xfd\xef\xab\x38\x09\x18\x12\xa4\x9e\x20\x99\x79\x6f\x9e\xe7\x3d\x47\x08\x21\x12\x42\x7d\xaf\x4a\x48\xae\x44\xf2\xa2\x0d\x69\xe5\x46\x95\xb2\xaa\x04\x07\x96\x47\x72\x53\xd6\x55\x39\x9e\x08\xf9\x0f\xa1\x96\xe3\xd7\x64\x72\x11\x80\x95\x5f\x15\xa8\xe7\x8b\x59\x9a\x5a\x60\xfe\x2a\x09\x56\x11\x07\x71\xfe\xe5\xe1\x9c\x1f\x03\xe7\x69\x80\x59\x60\xe3\xad\x86\x79\x3a\xea\xfe\xde\x5a\xe3\xab\xd1\xf8\x92\x54\x09\x13\x21\xbf\xa3\xb6\x86\x4d\xe6\x2e\xef\xc1\x6d\x8d\x7d\x9b\x52\xf3\xbb\x04\xed\x2d\xba\x8f\xd0\xcf\x72\x22\x36\xca\xa2\x5a\x15\xc0\x23\xd9\x2a\x93\xe3\xc3\xcc\x0d\x81\x5b\xfa\x15\x81\xeb\x8f\x3e\x52\x4d\xe0\x1e\x63\x29\xfb\x43\xf4\x95\x6c\xd0\x3a\xaf\x8a\xf6\x91\xa7\x1c\xf8\x79\x2a\x27\xe2\x94\xb2\x63\x89\xdf\x37\xed\xa7\x42\x0b\x24\xff\x7e\x6d\x28\xc3\xdc\x5b\xe5\xd0\x50\x72\x25\x76\xa1\x26\x44\x92\x22\xd7\x67\x5c\x28\xe6\xad\xb1\xe9\xcc\xbb\x35\x90\x43\xdd\x35\x26\xce\x7a\x68\xa9\x84\x48\x98\xd7\x35\x7a\x77\x0b\xee\x5b\xcd\xbb\x08\xb6\xdf\xc1\x07\x7f\x7e\x86\x9e\xcf\x76\x2a\x23\xe5\x05\xfc\xf0\xc6\x05\x3f\x65\xa7\x66\x8b\x94\x9a\x2d\x9f\xd5\x53\x59\xb3\x41\x46\x43\xcf\xe5\x2c\x07\x72\x7b\x09\x47\xec\x29\xaa\x9c\x0c\x3b\xd4\xbc\x74\xc6\xaa\x1c\x66\xba\x13\xfc\x82\xd9\x08\xfe\x78\x55\xf0\x91\x0f\x2b\x63\xdc\xcd\x01\x16\x1c\xc8\x54\xc1\x8d\x17\x64\xaa\x93\x25\x0f\x8d\xd0\xc6\x93\xbb\x87\xed\x83\xfd\xf7\x1d\xd9\x21\xe5\xf1\xa2\x0d\x2f\xac\xc9\xb0\x80\xf8\x3c\xda\x94\x95\x77\x60\xa9\x0b\xf6\x50\xa2\x5f\x0f\x0b\x56\x69\x89\xf4\xc4\x67\x00\x47\xd5\x3e\xae\x73\x71\x18\xd7\x55\x1b\xdc\x6e\x87\x99\x98\x73\x70\x51\x5c\xb6\xee\x9d\x8b\xcb\xb9\xa5\xaa\xa3\xb8\xfc\xfc\xa8\x9a\x6d\x56\xfb\x49\x13\xf1\x9b\x0d\x8d\x24\xf9\xa2\xa8\x9f\xa2\x3b\xd5\x1f\x14\x96\x79\xb1\xdb\x41\xc1\x70\x10\x74\x26\x31\xc9\x4b\xc4\x35\xd4\x23\x5b\x32\x4a\x4f\xa3\xd9\xd8\x39\x60\x16\x96\x2a\x87\x47\xc8\xc0\x02\xe9\xa6\x12\xd6\xf4\xbf\xe2\x6b\xcf\xce\x94\xf3\xba\x61\x2f\xad\x46\xf4\xee\x7e\x74\xad\xaf\x1b\xef\xa7\x81\x96\xe5\x44\x46\x24\x72\xe0\xa8\xdd\x67\x94\xd7\x60\x7b\x16\x1a\x0e\xb8\x45\xd7\x10\xbb\x5f\x87\x2f\xcb\xce\x83\x1e\xea\xe2\x09\x80\xdf\xfc\xb9\xf6\xe5\xdd\xd3\x49\xf3\x06\x2c\x77\x6b\x1f\x00\x3c\x37\xe5\xde\xc6\xf7\x3b\x17\xcd\x27\xe3\x46\x39\x75\x83\xfc\xc6\x71\xe0\x0c\xd7\xaf\x22\x1b\xea\x5b\xa3\xf4\x1a\x29\xaf\xe7\x3d\x82\x4a\x7f\x59\x74\x10\xeb\xd1\x16\x94\x83\x87\xaa\xcb\xc2\x7f\xb6\x5d\x6b\xdc\x54\x2a\x52\x39\xa4\x3d\xf6\x43\x04\xda\x1b\x5d\xe7\x76\xe0\x68\x35\xb0\x89\xf4\x6b\xb2\x07\xef\x0f\xd6\x44\xea\xe2\x6f\x00\x00\x00\xff\xff\xe9\x27\x63\xc3\x36\x07\x00\x00")

func tvmVarsTBytes() ([]byte, error) {
	return bindataRead(
		_tvmVarsT,
		"tvm/vars.t",
	)
}

func tvmVarsT() (*asset, error) {
	bytes, err := tvmVarsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tvm/vars.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"tvm/base.t":      tvmBaseT,
	"tvm/outputs.t":   tvmOutputsT,
	"tvm/params.t":    tvmParamsT,
	"tvm/resources.t": tvmResourcesT,
	"tvm/vars.t":      tvmVarsT,
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
	"tvm": {nil, map[string]*bintree{
		"base.t":      {tvmBaseT, map[string]*bintree{}},
		"outputs.t":   {tvmOutputsT, map[string]*bintree{}},
		"params.t":    {tvmParamsT, map[string]*bintree{}},
		"resources.t": {tvmResourcesT, map[string]*bintree{}},
		"vars.t":      {tvmVarsT, map[string]*bintree{}},
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