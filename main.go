package main

import (
	"fmt"
	"math/rand"
	"os"
)

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return fp.Sync()
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}

	if _, err = fp.Write(data); err != nil {
		fp.Close()
		os.Remove(tmp)
		return err
	}
	if err = fp.Sync(); err != nil {
		fp.Close()
		os.Remove(tmp)
		return err
	}

	fp.Close() // close BEFORE renaming (required on Windows)

	err = os.Rename(tmp, path)
	return err
}

func main() {
	// Test SaveData1
	err := SaveData1("test1.db", []byte("hello from SaveData1"))
	if err != nil {
		fmt.Println("SaveData1 error:", err)
	} else {
		fmt.Println("SaveData1 succeeded")
	}

	// Test SaveData2
	err = SaveData2("test2.db", []byte("hello from SaveData2"))
	if err != nil {
		fmt.Println("SaveData2 error:", err)
	} else {
		fmt.Println("SaveData2 succeeded")
	}
}