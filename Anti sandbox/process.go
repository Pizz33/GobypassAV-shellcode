func process() {
  executablePath, err := os.Executable()
  if err != nil {
    // 处理错误
    return
  }

  sourceFilename := filepath.Base(executablePath) // 源文件名称
  processName := filepath.Base(os.Args[0])        // 当前运行进程名称

  // 比较源文件名称和进程名称是否相同
  if strings.EqualFold(sourceFilename, processName) {
    // 进程名称和源文件名称相同，正常运行
    // 在这里执行正常的操作
  } else {
    // 进程名称和源文件名称不同，可能在沙箱环境中运行
    // 在这里执行相应的防沙箱操作
    os.Exit(0) // 退出程序
  }
}
