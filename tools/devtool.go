// Package tools 开发工具包
package tools

import (
    "encoding/json"
    "github.com/lemonit-eric-mao/commons/logger"
    "log"
    "net"
    "sigs.k8s.io/yaml"
    "time"
    "unsafe"
)

// YamlToMap 将 YAML 字符串解析为 map[string]interface{}
func YamlToMap(yamlStr string) (map[string]interface{}, error) {

    if yamlStr == "" {
        return nil, nil
    }

    var valuesMap map[string]interface{}
    if err := yaml.Unmarshal([]byte(yamlStr), &valuesMap); err != nil {
        return nil, err
    }

    return valuesMap, nil
}

// BytesToString 字节转字符串
func BytesToString(data []byte) string {
    return *(*string)(unsafe.Pointer(&data))
}

// StringToBytes 字符串转字节数组
func StringToBytes(data string) []byte {
    return *(*[]byte)(unsafe.Pointer(&data))
}

// StructToString 结构体转字符串
func StructToString(a any) string {
    buf, err := json.Marshal(a)
    if err != nil {
        log.Panic(err)
    }
    return string(buf)
}

// SetInterval 自定义，定时器工具
/**
func main() {
    chanStop := tools.SetInterval(3e9, func() {})

    time.Sleep(10e9)
    // 关闭定时器
    chanStop <- true
}
*/
func SetInterval(ms time.Duration, f func()) chan bool {

    ticker := time.NewTicker(ms)

    stop := make(chan bool)

    go func(tk *time.Ticker) {
        defer tk.Stop()
        for {
            select {
            case <-ticker.C:
                logger.Infof("定时器%v运行\n", &ticker)
                f()
            case <-stop:
                logger.Infof("定时器%v停止\n", &ticker)
                return
            }
        }

    }(ticker)

    return stop
}


// GetLocalIP 获取本地 IPv4 地址
func GetLocalIP() string {

    addrs, _ := net.InterfaceAddrs()

    for _, address := range addrs {
        // 检查 IP 地址是否为 IPv4，并且不是回环地址
        if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsPrivate() {
            logger.Debug(ipnet.IP.String())
            return ipnet.IP.String()
        }
    }

    return ""
}
