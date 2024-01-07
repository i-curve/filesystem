package handle

import (
	"filesystem/config"
	"io"
	"os"
	"path"
)

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
