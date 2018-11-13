package main

import (
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/os/gfile"
    "gitee.com/johng/gf/g"
)

func main() {
    path := "/tmp/glog-cat"
    glog.SetPath(path)
    glog.StdPrint(false).Cat("cat1").Cat("cat2").Println("test")
    list, err := gfile.ScanDir(path, "*", true)
    g.Dump(err)
    g.Dump(list)
}


