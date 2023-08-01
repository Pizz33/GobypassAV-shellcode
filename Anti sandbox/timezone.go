func shiqu() {
    // 加载北京时区
    loc, err := time.LoadLocation("Asia/Shanghai")
    if err != nil {
        return
    }

    // 获取当前时间
    now := time.Now()

    // 判断当前时区是否为北京时区
    if now.Location() != loc {
        return
    }
}
