package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/admpub/nging/v3/application/library/ip2region"
)

// NetInterface 本机网络
type NetInterface struct {
	Name    string
	Address []string
}

var ipv6Unicast *net.IPNet
var ipv6Regexp = regexp.MustCompile(ip2region.IPv6Rule)
var ipv4Regexp = regexp.MustCompile(ip2region.IPv4Rule)
var client = http.Client{Timeout: 10 * time.Second}

func init() {
	var err error
	// https://en.wikipedia.org/wiki/IPv6_address#General_allocation
	_, ipv6Unicast, err = net.ParseCIDR("2000::/3")
	if err != nil {
		panic(err)
	}
}

// GetNetInterface 获得网卡地址 (返回ipv4, ipv6地址)
func GetNetInterface() (ipv4NetInterfaces []NetInterface, ipv6NetInterfaces []NetInterface, err error) {
	allNetInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return ipv4NetInterfaces, ipv6NetInterfaces, err
	}

	for _, netInter := range allNetInterfaces {
		if (netInter.Flags & net.FlagUp) == 0 {
			continue
		}
		addrs, _ := netInter.Addrs()
		var ipv4 []string
		var ipv6 []string

		for _, address := range addrs {
			ipnet, ok := address.(*net.IPNet)
			if !ok || !ipnet.IP.IsGlobalUnicast() {
				continue
			}
			// 需匹配全局单播地址
			_, maskSize := ipnet.Mask.Size()
			switch maskSize {
			case 128:
				if ipv6Unicast.Contains(ipnet.IP) {
					ipv6 = append(ipv6, ipnet.IP.String())
				}
			case 32:
				ipv4 = append(ipv4, ipnet.IP.String())
			}
		}

		if len(ipv4) > 0 {
			ipv4NetInterfaces = append(
				ipv4NetInterfaces,
				NetInterface{
					Name:    netInter.Name,
					Address: ipv4,
				},
			)
		}

		if len(ipv6) > 0 {
			ipv6NetInterfaces = append(
				ipv6NetInterfaces,
				NetInterface{
					Name:    netInter.Name,
					Address: ipv6,
				},
			)
		}
	}

	return ipv4NetInterfaces, ipv6NetInterfaces, nil
}

// GetIPv4Addr 获得IPv4地址
func GetIPv4Addr(interfaceType string, interfaceName string, wanIPApiUrl string) (result string) {
	// 判断从哪里获取IP
	if interfaceType == "netInterface" {
		// 从网卡获取IP
		ipv4, _, err := GetNetInterface()
		if err != nil {
			log.Println("从网卡获得IPv4失败!")
			return
		}

		for _, netInterface := range ipv4 {
			if netInterface.Name == interfaceName && len(netInterface.Address) > 0 {
				return netInterface.Address[0]
			}
		}

		log.Println("从网卡中获得IPv4失败! 网卡名: ", interfaceName)
		return
	}
	resp, err := client.Get(wanIPApiUrl)
	if err != nil {
		log.Println(fmt.Sprintf("未能获得IPv4地址! <a target='blank' href='%s'>点击查看接口能否返回IPv4地址</a>,", wanIPApiUrl))
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取IPv4结果失败! 查询URL: ", wanIPApiUrl)
		return
	}
	matches := ipv4Regexp.FindAllStringSubmatch(string(body), 1)
	if len(matches) > 0 && len(matches[0]) > 1 {
		result = matches[0][1]
	}
	return
}

// GetIPv6Addr 获得IPv6地址
func GetIPv6Addr(interfaceType string, interfaceName string, wanIPApiUrl string) (result string) {
	// 判断从哪里获取IP
	if interfaceType == "netInterface" {
		// 从网卡获取IP
		_, ipv6, err := GetNetInterface()
		if err != nil {
			log.Println("从网卡获得IPv6失败!")
			return
		}

		for _, netInterface := range ipv6 {
			if netInterface.Name == interfaceName && len(netInterface.Address) > 0 {
				return netInterface.Address[0]
			}
		}

		log.Println("从网卡中获得IPv6失败! 网卡名: ", interfaceName)
		return
	}

	resp, err := client.Get(wanIPApiUrl)
	if err != nil {
		log.Println(fmt.Sprintf("未能获得IPv6地址! <a target='blank' href='%s'>点击查看接口能否返回IPv6地址</a>", wanIPApiUrl))
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取IPv6结果失败! 查询URL: ", wanIPApiUrl)
		return
	}
	matches := ipv6Regexp.FindAllStringSubmatch(string(body), 1)
	if len(matches) > 0 && len(matches[0]) > 1 {
		result = matches[0][1]
	}
	return
}