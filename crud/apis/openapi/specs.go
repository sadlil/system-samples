// Code generated for package openapi by go-bindata DO NOT EDIT. (@generated)
// sources:
// apis/openapi/gen/crudapi.swagger.json
package openapi

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

var _apisOpenapiGenCrudapiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x6f\xdb\x38\x10\xbe\xfb\x57\x0c\xb4\x7b\xd8\x05\xd2\x28\xed\x16\x05\x9a\x53\xb3\x71\xb6\x30\x50\xc4\x46\xec\x9c\xb6\x45\x40\x8b\x23\x9b\x2d\x45\x2a\x7c\x38\xf5\x16\xfe\xef\x0b\x52\x96\x2d\xd1\x92\xed\xd4\x7d\xb8\x40\x7c\x93\x38\x1c\x7d\x33\xfc\xe6\x41\xd2\x5f\x3a\x00\x91\x7e\x20\x93\x09\xaa\xe8\x1c\xa2\x17\xa7\x67\xd1\x89\x7b\xc7\x44\x2a\xa3\x73\x70\xe3\x00\x91\x61\x86\xa3\x1b\x1f\x92\x2c\xe7\x08\x97\x37\xb7\x5d\xb8\x18\xf4\xb4\x17\x06\x88\x66\xa8\x34\x93\xc2\x89\xcc\xce\x4e\x9f\x2f\xb5\x00\x44\x89\x14\x86\x24\x66\xa5\x0a\x20\xc2\x8c\x30\xee\x24\x3f\xfd\xc7\x94\xc9\xde\x4c\xdc\xf3\x69\x22\xb3\xc8\x4b\x2c\x3a\x00\x0b\x8f\xc1\x90\x89\x8e\xce\xe1\x5f\xff\x7a\x35\x5d\x90\xcc\x43\x19\x49\x2a\x87\xa8\x66\x2c\xc1\xf5\xc4\x0f\x7e\xe2\x54\x6a\xf7\xc5\x68\x6a\x4c\x7e\x1e\xc7\x5c\x26\x84\xbb\x77\xe7\xaf\x5e\xbf\x7e\x5d\xd8\xa7\x93\x29\x66\xb8\x56\xef\x65\xa3\x95\x86\x44\x0a\x6d\x6b\xe3\x24\xcf\x39\x4b\x88\x61\x52\xc4\x1f\xb5\x14\x6b\xd9\x5c\x49\x6a\x93\x3d\x65\x89\x99\xea\xb5\x5f\x63\x92\xb3\x78\xf6\x3c\x36\x92\xca\xaa\x8b\x26\x58\xf5\x98\x83\x6b\xb3\x8c\xa8\xb9\x33\xea\x1d\xd3\xc6\xd9\x0e\x37\x83\x4b\x50\x68\x14\xc3\x19\x6a\x20\xc0\x99\x36\x20\x53\xf0\x83\x72\xfc\x11\x13\xa3\x4f\x97\xeb\xe0\x95\xc8\x1c\x95\x07\xd5\xa3\x81\x03\xef\x4a\xa5\x55\x71\x85\x3a\x97\x42\xa3\xae\x41\x01\x88\x5e\x9c\x9d\x05\xaf\x00\x22\x8a\x3a\x51\x2c\x37\x4b\x16\x5c\x80\xb6\x49\x82\x5a\xa7\x96\x43\xa9\xa9\x8a\xa6\x30\xcb\xad\x02\xd9\x50\x06\x10\xfd\xae\x30\x75\x7a\x7e\x8b\x29\xa6\x4c\x30\xa7\x57\xc7\xb3\xe7\x25\xd0\x9b\xa5\xca\xa8\x36\x71\x51\x79\x5a\x54\xbf\x15\x51\x4c\x89\xe5\x66\x37\x6e\x01\x56\xe0\xe7\x1c\x13\x83\x14\x50\x29\xa9\xbe\x1d\x7c\x95\x27\x43\x43\x8c\xd5\x5b\x50\x77\x1a\xf0\x47\x39\x51\x24\x43\x83\x6a\x4d\xb2\xe2\x17\x18\x53\xc6\x06\x67\x19\x33\x21\x5a\xe6\x0d\xbc\xb7\xa8\xe6\xe1\x90\xc2\x7b\xcb\x14\x3a\x5a\xa4\x84\x6b\x0c\x86\xcd\x3c\xf7\x6a\xb5\x51\x4c\x4c\xc2\xc9\xa9\x54\x19\xf1\xf1\xc6\x84\x79\xf5\x32\x6a\x5b\x83\x16\xac\x32\x4d\x35\xfe\x2a\x60\xb5\x54\xa6\xaf\x28\xaa\xef\x86\xf7\xb1\x80\x0a\x3e\xfd\x20\xef\xa1\xb0\x59\xc0\x40\xff\x7e\xd4\xef\xf6\xef\x86\xa3\x8b\xd1\xed\xf0\xee\xf6\x7a\x38\xb8\xba\xec\xfd\xd3\xbb\xea\x06\xd3\x03\xc1\xc1\xd5\x75\xb7\x77\xfd\x76\xbb\x50\xef\xfa\x6e\x70\xd3\x7f\x7b\x73\x35\x1c\x6e\x17\xec\xf6\xaf\xaf\x76\x48\x5c\xbd\xbb\x1a\xed\x02\x75\xf1\xf7\xc5\xb5\x53\xd5\xad\x87\xe8\x87\x93\x30\x6b\x94\xe9\xa4\xd5\xf6\x47\x2e\x64\xae\x98\x54\xcc\x6c\xac\xd7\x77\x21\x56\xa7\xc1\xae\x7a\xb9\x2d\x5d\x13\x16\x59\x3f\xa9\x13\x98\x15\xe5\x45\xc1\x6d\x2e\x58\x97\x0a\x89\xc1\x55\xc9\x4a\xfc\xa3\x2b\x58\x02\x1f\xaa\xc5\x6a\xdf\x5a\xb5\xd6\x77\xf4\xd5\x6a\x0d\xf5\xa9\x5e\xf9\x5f\x0b\xf9\xc7\x92\xb6\x10\xbf\x69\xa4\xc2\x7b\xa3\x6c\x48\xfb\xc3\x56\xe9\xde\xa2\x36\xfb\x98\x7b\x58\xe4\x74\x2a\x1e\xab\x75\x81\xf1\x17\x46\x17\xfb\xb6\x82\x6f\xb1\xb9\x13\xd4\x4c\x4c\x38\x56\x63\x0b\xc6\x44\x23\x05\x29\x80\x19\x0d\xbd\xee\xbe\xb1\xb6\xfc\xc2\xd1\x07\xda\x12\xe7\x53\x94\xf9\x5f\x4b\x94\x31\xda\x1c\x63\x6e\x53\xf2\xb8\x18\xfb\x81\xa5\x85\x22\x47\x83\xad\x21\x30\x9a\x22\x74\xbd\xc8\x2a\x12\x8a\x19\x2e\x0e\xbe\x41\x00\xac\x75\x1f\x4d\x0c\x94\xce\x2f\x0c\xdb\x6c\x67\x72\xe5\xec\x31\xac\x40\xb7\x78\x0a\x85\xca\x67\x7f\xe9\x50\xc8\x6d\x7b\x29\xb8\xcd\x69\xb5\xc9\xb2\xfe\x51\x03\x11\x80\x9f\x99\x36\x4c\x4c\xbe\xa6\xd5\x5a\x6b\x3d\x1a\xf6\xb7\x55\x80\x35\xd4\xa7\x22\xe0\x7f\x3f\x93\xf9\x7b\xec\x77\x7e\x4c\xcb\xf7\xb8\x5c\x19\x8c\x7a\xef\xce\xb9\x24\xb4\x71\x70\x0b\x19\x7d\xc4\x6c\x4c\x58\x74\xb6\x3d\x7f\xbf\x36\x73\x75\xa0\x5a\x01\xb9\x3e\x83\xcc\x95\x34\x72\x6c\xd3\x0b\x31\xaf\xf6\x9d\x2d\x9e\x6b\xf3\x58\xf4\x66\x39\xa1\x96\x10\xda\x28\xb2\xd8\xc8\x6d\x84\x52\x0f\x8c\xf0\x41\x43\xf9\x2a\x3b\xe5\x75\x94\x1c\x80\x34\x91\xb4\x15\x28\x13\x06\x27\xc1\xe9\x4e\xfd\xcc\xe8\xaf\x17\x51\x63\x1c\x66\xa8\x35\x99\xec\xef\x81\xca\x54\x8a\x86\x30\xbe\x91\x4d\xcb\xa9\x44\x29\x52\x0f\x88\x88\x19\xcc\x36\x19\xbb\x9d\xed\x2d\x64\xad\x2e\x7f\x73\xb2\xa9\x93\xa9\x3c\xf0\x6f\xd8\x35\x1d\xb0\x26\xc1\x01\xf8\x16\xbc\x61\x70\xed\x03\x6f\x59\x0f\x8e\x0a\x5f\xb8\x63\x39\x2a\x70\x1b\xa7\xec\x07\xa2\xfb\x49\xdc\xde\x4c\xc4\xbb\x69\x3d\x0a\xae\x62\x1e\x6b\x2e\x0b\x8b\xc5\xb6\x23\xd5\xb0\xc5\xb0\x96\x51\xb0\x82\xdd\x5b\xe4\x73\x50\x98\x2b\xd4\x28\x8c\xeb\xe3\xe6\xb5\xdd\x0c\x13\x60\xa6\x08\xda\x48\x45\x26\x78\x0a\x6e\x1b\x34\x23\xdc\x22\xc8\xf4\xbd\x60\x14\x1e\x18\xe7\x30\x46\x98\xa0\x70\xdd\x1d\x52\x18\xcf\x8b\x29\x73\x6d\x30\x83\x87\x29\xe3\xe8\x5f\xb8\xf5\x01\xa6\x61\x8c\xae\x49\x4c\xbc\xf0\x69\x73\xa2\x5a\xd6\xee\xaf\xb5\xee\x56\xa3\x82\x5c\xc9\x19\xa3\x48\x81\x32\x9d\x73\x32\x07\xa7\x14\x64\xba\xc2\xd2\xf2\xed\xba\xae\x47\xe7\xd8\xd5\xf1\xea\x81\xf0\x3d\xbd\x90\x42\xa9\xaf\x44\x3e\xea\x77\xfb\x27\x90\x10\x01\x52\xf0\xb9\xf3\xbc\x46\x03\x04\x0a\xbd\xab\xb5\x81\xc1\x1f\x67\xcf\x5e\xfe\xd9\x62\x63\x71\x34\x4a\x2f\xc2\x5e\x74\x2b\xca\x75\x75\x72\x0d\xf0\x33\xc3\x32\xdc\x6a\x46\x91\x18\x29\x10\x03\x4e\x56\x1b\x92\xe5\x27\x55\xae\x3d\x4c\xb1\x60\x17\x0a\xa3\xe6\x55\x6a\x60\x95\x48\x56\xa3\x6a\x5d\x2b\x42\x39\x13\x07\x73\xa5\x74\x76\xa9\x0f\x74\x8e\x09\x4b\x19\x52\x17\x00\xd4\x16\xdb\x96\x13\x60\x82\xb2\xc4\xef\x77\x3c\x89\x58\x86\x0e\xe4\xca\x0e\x43\xf4\xa7\xf7\x42\x4f\xa5\xe5\xd4\xad\x4c\x22\xb3\xdc\xed\xea\xdb\x68\xae\xc3\x16\x03\x76\xa5\x98\xe1\xe6\x2d\x4c\x68\x50\x21\xd2\x46\xf4\x8d\x7e\x28\x98\xed\x63\x7f\xd9\x62\xd4\xf2\x42\x7d\x57\xd7\x90\xcb\xda\xfb\xa5\x60\x15\x36\x2f\x77\xf6\xb8\xd6\xd9\x71\xa1\xb3\xc7\x55\xce\xb6\x4b\x9c\x1d\xd7\x37\xdb\x2f\x6e\x3e\x54\x7c\xb9\xf3\xb2\x66\x9b\xdf\x97\x2b\xe7\xdc\xb3\x64\x64\xc1\xb3\x5c\x6a\xcd\xc6\xdc\x65\x61\xcf\x3d\x99\xee\x58\x8e\x86\x2d\xea\x31\x54\xfd\xd5\xf6\x00\x3f\x1b\x54\x82\xf0\xae\x4c\x2a\xfb\x03\xab\x78\xf9\x5f\x0a\x7d\x1e\xc7\x13\x66\xa6\x76\x7c\x9a\xc8\x2c\xd6\x2e\x2a\x79\xac\xfd\x7f\x42\x74\x6c\x14\x62\x9c\x11\x26\xe2\x44\x59\xea\x3e\xb7\xe8\x2c\x3a\xff\x07\x00\x00\xff\xff\x46\x72\xdc\xbb\x62\x22\x00\x00")

func apisOpenapiGenCrudapiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apisOpenapiGenCrudapiSwaggerJson,
		"apis/openapi/gen/crudapi.swagger.json",
	)
}

func apisOpenapiGenCrudapiSwaggerJson() (*asset, error) {
	bytes, err := apisOpenapiGenCrudapiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "apis/openapi/gen/crudapi.swagger.json", size: 8802, mode: os.FileMode(420), modTime: time.Unix(1679238772, 0)}
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
	"apis/openapi/gen/crudapi.swagger.json": apisOpenapiGenCrudapiSwaggerJson,
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
	"apis": &bintree{nil, map[string]*bintree{
		"openapi": &bintree{nil, map[string]*bintree{
			"gen": &bintree{nil, map[string]*bintree{
				"crudapi.swagger.json": &bintree{apisOpenapiGenCrudapiSwaggerJson, map[string]*bintree{}},
			}},
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
