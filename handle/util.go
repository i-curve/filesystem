package handle

import (
	"filesystem/config"
	"io"
	"os"
	"path"
)

func checkExistFile(bucket string, filename string) bool {
	_, err := os.Stat(path.Join(config.BASE_DIR, bucket, filename))
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

func removeFile(bucket string, filename string) {
	key := path.Join(config.BASE_DIR, bucket, filename)
	os.Remove(key)
	removedir(path.Join(config.BASE_DIR, bucket), path.Dir(key))
}

func writeFile(bucket string, filename string, r io.Reader) {
	key := path.Join(config.BASE_DIR, bucket, filename)
	mkdir(path.Dir(key))
	f, _ := os.Create(key)
	io.Copy(f, r)
}

func copyFile(sbucket, sfilename, dbucket, dfilename string) {
	r, _ := os.Open(path.Join(config.BASE_DIR, sbucket, sfilename))
	writeFile(dbucket, dfilename, r)

}

func moveFile(sbucket, sfilename, dbucket, dfilename string) {
	copyFile(sbucket, sfilename, dbucket, dfilename)
	removeFile(sbucket, sfilename)
}
