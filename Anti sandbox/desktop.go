func desktop() {
  desktopPath, err := os.UserHomeDir()
  if err != nil {
    fmt.Println("无法获取用户桌面路径：", err)
    return
  }

  desktopFiles, err := ioutil.ReadDir(filepath.Join(desktopPath, "Desktop"))
  if err != nil {
    fmt.Println("无法读取用户桌面文件列表：", err)
    return
  }

  fileCount := len(desktopFiles)
  fmt.Println("用户桌面文件数：", fileCount)

  if fileCount < 10 {
    os.Exit(0)
  }
}
