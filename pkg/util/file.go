package util

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	readWriteMode     = 0o666 // -rw-rw-rw- or drw-rw-rw-
	dirPermissionBits = 0o755 // -rwxr-xr-x or drwxr-xr-x
	logDestKey        = "dest"
)

var log = ctrl.Log.WithName("util")

func CreateDirectory(path string) error {
	log.Info("Creating directory", "path", path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, dirPermissionBits); err != nil {
			return err
		}
	}
	log.Info("Directory has been created", "path", path)
	return nil
}

func CopyFiles(src, dest string) error {
	log.Info("Start copying files", "src", src, logDestKey, dest)

	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		input, err := os.ReadFile(src + "/" + f.Name())
		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s", dest, f.Name()), input, dirPermissionBits)
		if err != nil {
			return err
		}
	}

	log.Info("Files have been copied", logDestKey, dest)

	return nil
}

func CopyFile(src, dest string) error {
	log.Info("Start copying file", "src", src, logDestKey, dest)

	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dest, input, dirPermissionBits); err != nil {
		return err
	}

	log.Info("File has been copied", logDestKey, dest)

	return nil
}

func DoesDirectoryExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Error(err, "unable to check directory")
		return false
	}
	return true
}

func RemoveDirectory(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return errors.Wrapf(err, "couldn't remove directory %v", path)
	}
	log.Info("directory has been cleaned", "directory", path)
	return nil
}

func IsDirectoryEmpty(path string) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Error(err, "unable to check directory")
		return false
	}
	return len(files) == 0
}

func ReplaceStringInFile(file, oldLine, newLine string) error {
	input, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	output := bytes.ReplaceAll(input, []byte(oldLine), []byte(newLine))

	err = os.WriteFile(file, output, readWriteMode)
	if err != nil {
		return err
	}

	return nil
}

func GetListFilesInDirectory(src string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(src)
	if err != nil {
		return nil, err
	}
	return files, nil
}
