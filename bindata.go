package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _include_basht_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\xd1\x6f\x9b\x3e\x10\x7e\xc6\x7f\xc5\xfd\x90\x55\x88\x2a\x7e\x4a\xb6\xb7\x54\x68\xe9\xba\x54\xab\xd4\x25\x53\x68\x9f\xa2\x08\x21\x62\x1a\x34\x0a\x91\x6d\xb6\x49\x19\xff\xfb\xee\x6c\x87\x42\xb6\xe5\x01\xc7\xe7\xef\xbe\xef\xee\x3b\x9b\xa5\x69\xde\xd4\x3a\x2b\x6b\x15\x4e\xe0\xc4\xbc\xaa\xc9\xb3\x0a\xc4\x0d\x14\x8d\x04\x01\x65\x0d\x3e\x3f\x2d\xe6\xef\x3a\xff\x06\xf6\x0d\x6c\xb7\xb8\x17\x3e\xc4\x31\xae\x33\x1f\x76\x3b\xb8\xba\x02\x29\x74\x2b\x6b\x98\x12\xa4\xc6\x5c\xb7\x9f\xb1\x8e\xa1\x40\x91\x95\xd5\x97\x2c\x97\xcd\x50\xa1\x2a\x6b\x11\x13\xc5\x0d\xa8\x43\x59\x68\xe6\xa5\xa9\x16\x4a\xa7\x4a\x67\xba\x55\xf1\xac\x0f\xbc\x0a\xa5\xb2\x17\x02\x53\xce\x1c\xf8\xc2\x47\xde\x27\x43\x1b\x07\xe2\x3b\x92\x0d\x34\x80\x7f\xbc\x4d\x3e\xa7\xc9\xfa\x79\x73\xb7\x9c\xf3\xc7\x87\xd5\x72\xb5\x0e\x18\x7b\xc5\x16\x87\xfa\xb2\xad\x81\x92\xc4\x1e\x50\x51\x6a\xfc\x36\x47\xd8\xb7\x32\xd3\x65\x53\x33\x6f\x2f\xf2\x2a\x93\x02\xa2\x0c\x8e\xb2\xc9\xb1\x06\xb1\x67\x1e\x66\xc5\x53\xe6\xd9\x44\xf3\x0f\x5d\x2a\x70\x43\x46\xf1\x05\xf5\xcf\x3c\x4f\x35\xad\xcc\x05\x1a\x44\x27\x3e\x06\x08\xa5\x0d\x24\xec\x79\xef\xe1\x17\xbc\x48\x71\x84\xa0\x0f\x15\xf0\x94\x06\x18\xce\x7e\x7c\x83\xe0\x74\x94\x65\xad\x81\xbf\xef\x82\x89\xe3\xf5\xca\x02\xfe\x83\xb7\x89\xa1\x82\xf6\x69\x3e\x7d\x85\xdb\xc5\x8e\x06\xa5\x0f\xa2\x26\xbc\xd7\xd6\x4a\x68\x18\x39\x0b\x63\x5b\x0d\x4c\xe4\x87\x06\xfc\x18\x87\xba\x79\x5e\x01\x92\x9a\xa8\xf1\x05\x6d\x4f\x96\x77\xeb\xd5\xa7\xc4\x06\xb9\x36\xcb\x78\x58\xfc\x34\xda\xcf\x23\xfe\xa1\x73\x14\xcd\xf1\x92\xe1\xec\x71\xcc\xc3\x90\x13\x20\xe2\x46\x69\x32\x31\xc7\x7d\x2f\xd7\x71\x48\xfd\xd9\x28\x19\x4f\x78\x5c\xaf\x67\x0e\x89\x6e\x98\xeb\x38\xd2\x36\x57\x73\x8a\xf7\x72\xe0\x82\xeb\x2f\x8a\x22\xf8\x7a\x9b\x24\xd8\x20\x84\xfc\x74\xae\xa3\x53\x13\x5b\x98\xa8\x94\xb5\xe3\x3c\x5f\x12\xb4\x7f\x7b\xcd\x01\xd5\xfd\xed\xc3\xe3\xbf\xa8\x2e\x6b\x73\x66\xfb\x17\x75\x39\x36\xc0\xdf\x25\x72\x00\x70\x35\x95\xec\x6d\x31\x5f\x7a\x6b\xcc\x7d\x2d\xee\xac\x6a\x8b\xb6\xaf\x74\x3a\x12\xb5\x82\x9b\x0c\xaf\x22\xbd\x00\x92\x54\xff\xfb\xac\x17\xb2\xe7\xe4\x12\x06\x9d\x21\x7f\xcd\x01\x27\xe2\x1e\xd1\x9f\x1c\x64\x8f\x09\xfe\x2c\xf5\x19\xcc\xa8\xf0\x8e\xfd\x0e\x00\x00\xff\xff\xb6\xec\xc7\x36\x75\x04\x00\x00")

func include_basht_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_basht_bash,
		"include/basht.bash",
	)
}

func include_basht_bash() (*asset, error) {
	bytes, err := include_basht_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/basht.bash", size: 1141, mode: os.FileMode(420), modTime: time.Unix(1426344369, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"include/basht.bash": include_basht_bash,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"include": &_bintree_t{nil, map[string]*_bintree_t{
		"basht.bash": &_bintree_t{include_basht_bash, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

