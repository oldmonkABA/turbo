// Code generated by go-bindata.
// sources:
// configs/interleaver.json
// DO NOT EDIT!

package turbo

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

var _configsInterleaverJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x9a\xbd\x8e\x65\x35\x10\x84\xe3\x99\xa7\x40\x1b\x6f\x70\xba\xdb\xbf\x84\x04\xbc\x04\x22\x42\x44\xac\xe0\x01\x10\xef\xce\x45\x04\x53\x35\x3a\xdf\x5e\xdf\x6c\xb4\xb2\xd7\xae\xd3\x76\x55\x57\xf9\xfe\xf2\xfe\xf6\xf7\xfb\xdb\xdb\x97\x9f\xbe\xfd\xf5\xdb\x1f\xdf\x7e\xff\xf3\xcb\x8f\x3f\xb4\xeb\xeb\x7f\xff\xf2\x73\x3c\xfe\xae\xff\xff\xcc\xc7\x9f\x71\xbd\xbf\xfd\xf3\xf5\x66\xf8\xfa\x18\x3e\x65\x78\xde\x0f\xef\xe3\x63\x78\xec\x8f\xf1\x0d\xc6\x8f\x76\xff\xdf\x8f\xfb\xe1\x33\xef\x87\xaf\xfb\xe1\x4b\xb0\x46\x7c\x8c\x4f\x00\xbb\x04\x6c\x97\xe1\xb0\xf9\x3d\xe0\xbf\x6f\xf7\xe3\xe3\xba\x47\x9b\x80\x36\x42\xe0\x36\x59\x60\xd1\x02\xa9\x80\x2f\x29\xef\x06\xc4\x91\x02\x39\x04\x73\x01\xe6\x28\x01\x2d\x05\x2e\xda\x52\x13\xcc\xa1\x35\xbb\xa0\x68\xd1\xf3\x7e\x05\x1a\x3f\x04\x74\x86\x9e\x51\x02\x3d\x14\xf4\x65\x1f\x96\x60\xcc\x71\xbf\x08\x8d\x5f\x02\xbb\x0b\xec\x46\xb5\xde\x82\x3a\xa5\x72\x0d\x60\xe7\xa5\xb5\x96\x09\x1d\x50\xe7\x25\xa8\x53\xb6\xd4\xa1\xd4\x19\x70\xbe\x0b\x30\x64\x36\x58\x81\x26\x94\x80\x5e\x72\xfa\x3a\x81\x56\xf6\x4a\x39\x1c\x83\x40\x2b\x7f\x95\x7c\xa5\x41\xa0\x8d\xc1\x0e\x2e\x44\x0e\x3a\xe0\x9b\x40\x28\x8b\xd9\x9e\x68\xc2\xa2\x6b\x9d\x44\xdb\xa9\x54\xa6\x44\x8c\xb5\xdb\xc0\xdc\x13\x4e\x78\x29\x99\x95\xe0\x9e\xb0\x42\x29\x9b\xd9\x0a\x00\xbb\xf2\xd5\x8b\x5d\xca\x66\x3a\x63\x41\xf1\xaa\xec\x88\xf7\xe7\x14\x5b\xc6\x67\x5b\x6a\xb1\x08\xb7\xf2\x99\x51\x07\x2d\x31\xec\x66\x1f\xb0\x78\x29\xa1\xad\x78\xce\x36\xa5\x74\xd6\x04\xf6\xa6\x3d\x29\x9f\x9d\xd0\x53\x19\x9f\x35\x05\x01\x33\x9a\x11\x5a\x57\x14\xd4\x9b\x5c\x26\x5e\x5d\xc5\x05\x0a\xde\x94\xd3\xb2\x3f\x27\xc1\xa6\x9c\xa6\x9b\x8a\x0b\xbe\x6d\x53\x52\x6b\x7a\x31\x68\x09\x25\xb5\xad\x4b\x60\x53\xa6\xac\xa6\x34\x18\x44\x21\x4d\x69\xcd\x66\x04\x54\xbc\x29\xaf\xa5\xe2\xe8\x2b\x60\x8a\x12\x9b\x2d\x42\xbb\x52\x62\x5b\x36\x01\xdb\xd1\x05\xdf\x8a\x9a\xb4\x66\xc4\xa6\x52\x3c\x12\xb6\xd5\x95\xda\xf4\x58\xc1\xa7\xea\xca\x6c\x25\x9b\x1a\x34\xc1\xba\x2e\xdd\x12\x9c\xa9\xae\xac\x53\xfd\xb9\x64\x74\x6b\x8a\x54\x8b\x1b\x91\x67\x57\x4e\x18\xca\x09\xb4\xa9\x4d\x84\x0e\xb0\x87\xde\xd7\x13\xc9\x18\x7a\xf9\xb4\x01\x4e\x6a\x37\x87\xd9\x1b\xd9\x13\x1d\xa8\x61\x7a\xaf\x3d\x08\x69\xc6\xd0\x53\xae\xc4\x96\xc4\x21\x63\x91\x2c\x91\xcd\xd1\x03\x68\xc4\x46\x9a\x31\x55\x2b\xe7\x3e\xd0\xca\x69\xd2\xa7\x9f\x6a\xe7\x86\x29\x9d\x7a\x54\x18\x3f\x0c\xf7\x7c\xae\x1a\xd3\x64\xa6\x3f\x17\x8d\x65\xa2\x31\x9f\xd7\x7b\x59\x53\x9b\xf3\x40\x33\x96\x12\xfa\x89\x66\x2c\xa3\x67\xfd\xb4\x28\x1a\x8b\xba\x48\xfa\x52\xcb\x7a\x42\xbd\x4c\xa8\x1a\x6b\x9b\x87\x51\xb9\x24\x3b\xbf\x83\x18\x9d\x64\x63\x93\xa1\xa4\x96\x7e\x5b\x43\xd5\x0c\x08\x4d\x19\x2f\xba\x80\xad\xd4\xd6\xf7\x81\x6a\x6c\xe5\x36\x25\xc3\x60\x73\xaf\xec\xd6\x0f\x5a\xc9\xc7\x81\x6b\xaf\xe9\x46\x5c\xc6\x56\x07\xc2\xf1\x70\xd9\x26\x35\x16\x82\xd0\x2a\x16\x23\x0c\xad\x08\x75\x62\x61\xa6\x5d\x05\x8a\x9a\x9e\x30\x87\x7c\xa2\x1e\x61\x7e\x54\x6f\x15\xc9\x47\x98\xfb\x3b\xb1\x1c\xe1\x56\x6b\xcb\x94\x44\xec\x66\x6c\x4e\x5c\x47\x98\x89\x48\x8b\x8e\x88\x53\xa2\x28\x81\x20\x11\x09\x6b\x8f\x55\xa8\x16\xa1\xb7\x4e\x34\xda\xc1\xcd\x0a\xeb\xf9\xcc\x49\x10\x10\xeb\xae\x74\x8d\x85\xd9\x48\x57\xad\xd2\x33\x4c\xec\x18\xdd\x22\x1e\x15\x69\x4e\x91\x5e\x54\x92\x18\xa4\x0b\xa8\x24\x31\x8c\xe6\x97\xe5\xaf\x74\xec\x87\xd2\xb6\xf2\xca\xee\x18\x56\x19\x09\xab\xc8\xed\x01\xbd\x7b\x4c\x23\x55\xd3\x45\x9a\xb1\x51\x1c\x08\xff\xb2\x84\xe8\x44\x4f\x1e\x2d\x99\xc2\x37\x79\xa0\xe3\xb2\x94\xf3\xf4\x4c\x16\x7e\xb1\xad\x94\x57\x47\x21\xe2\x56\xca\x53\x51\x19\xf8\xc5\xb6\x71\xde\x3a\x91\x95\xbc\x3c\x85\xd3\xae\x8b\xb6\xf6\x60\xf6\xf5\x9a\xb0\xa4\x45\xcd\x27\xc2\x92\x96\x88\x9a\xb0\xc4\x77\xf2\x41\x3d\x63\xd7\x3e\xc1\x62\x39\x53\x76\x0b\xbf\x30\x56\xd4\xdb\x5f\xaa\x60\x8d\x1a\xe8\x34\xdb\x9e\x2a\xf9\x0d\xe3\x4b\xf7\x96\x1e\xbc\x50\x96\x67\xce\xec\xc4\xa1\xe4\xa0\x58\x35\xb0\x9c\xe6\x38\x8c\x00\x3a\xf1\x4c\x5a\x77\x1f\xaa\x17\x41\x22\x9e\xd6\x4d\x9f\x48\x4c\x5a\xdf\xaa\x0c\x50\x17\xe1\xb7\x0e\xf1\x44\x62\xea\xf2\x60\xc8\x5e\x8a\x28\xaf\xba\x26\xc4\x36\x64\x7f\x2b\xdc\x41\x3d\xd7\x98\xf2\xcc\xde\x9f\x8c\x30\xd2\x54\xc9\x68\xe6\x37\xa9\x94\x65\x39\xa8\x87\x56\x38\x45\xc9\x5c\xa7\x24\x91\x79\x59\xa6\xd4\x6c\x19\xac\x4c\xb7\x4c\x57\x2b\x93\x18\xea\x76\x7a\x4b\x29\x22\x9a\x1a\x4e\x80\xfa\xd1\xb0\x38\xd3\xda\x38\xeb\x94\x29\xfb\xaf\xe9\x24\xa8\xed\x22\xee\x6d\x59\x44\x71\xa4\x35\xb5\x2d\x6c\x1f\x8a\x07\x53\xe4\x6d\x24\x38\xf5\x3d\x93\xda\xa6\x76\x29\x09\xba\x3b\xc6\x64\xf1\x52\x12\x3c\x11\x9b\x66\x2f\x76\x9a\x6f\x47\x51\xec\x67\x4f\x4b\x76\x09\x70\x15\x7b\x02\xf1\x55\x68\x8a\x45\xf5\x6d\xaa\x6e\xd0\x6b\x65\xb3\x44\x59\x9d\x4f\xd0\x83\x65\xb3\xd8\x33\xcb\x98\x83\x3e\xc0\xa7\x04\xd0\x4c\x19\x25\x9f\x9e\xb7\x95\xa5\x0a\xb4\x8e\xa5\x5b\x16\x2b\x50\x77\xde\x26\x75\xe7\x41\x67\xb3\x59\x6c\x33\x4f\xae\x40\xb3\x84\xc4\x36\x86\xd1\xbd\x65\x11\x2a\xb5\x8d\x5e\xff\x9a\xf9\x7e\xbb\x00\x18\x39\x76\x33\xe5\x96\xf6\x50\xe6\xd1\xcd\x61\x9f\x34\x01\xdd\xec\xb2\xde\x32\x4c\x67\xdd\xfc\x86\x3d\x23\xd1\xaf\x37\x8a\x7e\x71\x10\x98\x4c\xbb\x2f\x55\x15\x40\x7a\xee\xe6\x33\xbd\xd3\x20\xfc\xe6\x1a\x4f\xec\x6f\x37\x0b\xe8\x8b\x50\x25\xcd\xd1\x69\xa7\x11\x14\x31\x74\x73\x67\x3e\x85\x8a\x6f\x4e\x2b\xc6\xc9\xe9\xef\x66\x9c\xfc\xa5\x84\x6a\x69\x2e\xa8\x34\xae\x0d\x4a\xa4\xfa\xee\xd0\x04\x05\x7d\xe6\x61\xfe\x44\x57\xa1\xa7\xc4\x07\x2b\x5d\xb0\x08\xbd\x6f\x0e\xfb\x15\x49\x8e\xfa\xcc\x64\xef\xbf\xfe\x1b\x00\x00\xff\xff\x45\xfc\xb5\x80\xea\x24\x00\x00")

func configsInterleaverJsonBytes() ([]byte, error) {
	return bindataRead(
		_configsInterleaverJson,
		"configs/interleaver.json",
	)
}

func configsInterleaverJson() (*asset, error) {
	bytes, err := configsInterleaverJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/interleaver.json", size: 9450, mode: os.FileMode(436), modTime: time.Unix(1436091165, 0)}
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
	"configs/interleaver.json": configsInterleaverJson,
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
	"configs": &bintree{nil, map[string]*bintree{
		"interleaver.json": &bintree{configsInterleaverJson, map[string]*bintree{}},
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