package classpath

import(
    "os"
    "path/filepath"
    "strings"
)

//通配符classpath

func newWildcardEntry(path string) CompositeEntry {
    //去掉通配符"*"
    baseDir := path[:len(path)-1]
    compositeEntry := []Entry{}
    walkFn := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        //不处理目录
        if info.IsDir() && path != baseDir {
            return filepath.SkipDir
        }
        //创建ZipEntry 并放到
        if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
            jarEntry := newZipEntry(path)
            compositeEntry = append(compositeEntry, jarEntry)
        }
        return nil
    }
    filepath.Walk(baseDir, walkFn)
    return compositeEntry
}