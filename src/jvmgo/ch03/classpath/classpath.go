package classpath

import(
    "os"
    "path/filepath"
)

type Classpath struct {
    bootClasspath Entry
    extClasspath Entry
    userClasspath Entry
}

//解析classpath
func Parse(jreOption, cpOption string) *Classpath  {
    cp := &Classpath{}
    //jre classpath
    cp.parseBootAndExtClasspath(jreOption)
    //用户classpath
    cp.parseUserClasspath(cpOption)
    return cp
}

//解析用户classpath
func (self *Classpath) parseBootAndExtClasspath(jreOption string) { 
    jreDir := getJreDir(jreOption)
    // jre/lib/* 类路径
    jreLibPath := filepath.Join(jreDir, "lib", "*")
    self.bootClasspath = newWildcardEntry(jreLibPath)
    // jre/lib/ext/* 扩展类路径
    jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
    self.extClasspath = newWildcardEntry(jreExtPath) 
}

//解析用户classpath
func (self *Classpath) parseUserClasspath(cpOption string) {
    if cpOption == "" {
        cpOption = "."
    }
    self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error){
    //这里传入的className不包括后缀，需要加上后缀
    className = className + ".class"
    //从类路径读取
    if data, entry, err := self.bootClasspath.readClass(className); err == nil {
        return data, entry, err
    }
    //从扩展类路径读取
    if data, entry, err := self.extClasspath.readClass(className); err == nil {
        return data, entry, err
    }
    //从用户类路径读取
    return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
    return self.userClasspath.String()
}

func getJreDir(jreOption string) string {
    if jreOption != "" && exists(jreOption) {
        return jreOption
    }
    if exists("./jre") {
        return "./jre"
    }
    if jh := os.Getenv("JAVA_HOME"); jh != "" {
        // fmt.Print(jh)
        return filepath.Join(jh, "jre")
    }
    panic("Can not find jre folder!")
}

func exists(path string) bool {
    if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return false
        } 
    }
    return true 
}