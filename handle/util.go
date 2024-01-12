package handle

import (
	"filesystem/config"
	"io"
	"os"
	"path"
	"sort"
)

type Dirs struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
	IsDir    bool   `json:"is_dir"`
}

func listDir(dirs string) ([]*Dirs, error) {
	var data []*Dirs
	entries, err := os.ReadDir(path.Join(config.BASE_DIR, dirs))
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		data = append(data, &Dirs{
			Path:     dirs,
			Filename: entry.Name(),
			IsDir:    entry.IsDir(),
		})
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i].IsDir == data[j].IsDir {
			return data[i].Filename < data[j].Filename
		}
		return data[i].IsDir
	})
	return data, nil
}

func checkExistFile(bucket string, key string) bool {
	_, err := os.Stat(path.Join(config.BASE_DIR, bucket, key))
	return err == nil
}

func mkdir(dirkey string) {
	os.MkdirAll(dirkey, os.ModePerm)
}

func removedir(root, dirkey string) {
	dirs, err := os.ReadDir(dirkey)
	if err == nil && len(dirs) == 0 && root != dirkey {
		os.Remove(dirkey)
		removedir(root, path.Dir(dirkey))
	}
}

func removeFile(bucket string, key string) {
	filename := path.Join(config.BASE_DIR, bucket, key)
	os.Remove(filename)
	removedir(path.Join(config.BASE_DIR, bucket), path.Dir(filename))
}

func writeFile(bucket string, key string, r io.Reader) {
	filename := path.Join(config.BASE_DIR, bucket, key)
	mkdir(path.Dir(filename))
	f, _ := os.Create(filename)
	io.Copy(f, r)
}

func copyFile(sbucket, skey, dbucket, dkey string) {
	r, _ := os.Open(path.Join(config.BASE_DIR, sbucket, skey))
	writeFile(dbucket, dkey, r)

}

func moveFile(sbucket, skey, dbucket, dkey string) {
	copyFile(sbucket, skey, dbucket, dkey)
	removeFile(sbucket, skey)
}
