func PathExists(path string) (bool, error) { //判断文件是否存在
  _, err := os.Stat(path)
  if err == nil {
    return true, nil
  }
  if os.IsNotExist(err) {
    return false, nil
  }
  return false, err
}
func fack(path string) { //判断虚拟机关键文件是否存在
  b, _ := PathExists(path)
  if b {
    os.Exit(1) //如果是虚拟机就退出当前进程
  }
}
func check() {
  fack("C:\\windows\\System32\\Drivers\\Vmmouse.sys")
  fack("C:\\windows\\System32\\Drivers\\vmtray.dll")
  fack("C:\\windows\\System32\\Drivers\\VMToolsHook.dll")
  fack("C:\\windows\\System32\\Drivers\\vmmousever.dll")
  fack("C:\\windows\\System32\\Drivers\\vmhgfs.dll")
  fack("C:\\windows\\System32\\Drivers\\vmGuestLib.dll")
  fack("C:\\windows\\System32\\Drivers\\VBoxMouse.sys")
  fack("C:\\windows\\System32\\Drivers\\VBoxGuest.sys")
  fack("C:\\windows\\System32\\Drivers\\VBoxSF.sys")
  fack("C:\\windows\\System32\\Drivers\\VBoxVideo.sys")
  fack("C:\\windows\\System32\\vboxdisp.dll")
  fack("C:\\windows\\System32\\vboxhook.dll")
  fack("C:\\windows\\System32\\vboxoglerrorspu.dll")
  fack("C:\\windows\\System32\\vboxoglpassthroughspu.dll")
  fack("C:\\windows\\System32\\vboxservice.exe")
  fack("C:\\windows\\System32\\vboxtray.exe")
  fack("C:\\windows\\System32\\VBoxControl.exe")
}
func main() {
  a, _ := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
    if a[0] != "zh-CN" {
      os.Exit(1)   
    } else {
                         check()   
         }
}
