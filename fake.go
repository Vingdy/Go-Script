package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println(opendir("./"))
	//fmt.Println(ReadLine("./18.197.180.159-2019-11-21-00.golang.gz"))
	//opendir("./")
}

func opendir(path string) error {//遍历文件夹
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			opendir(path + "/" + file.Name())
		} else {
			err := ReadLine(file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*
18.197.180.159-2019-11-21-00.golang.gz
2019-11-29 15:12:40	/event?campuuid=&clickid=&type=appsflyer&device=samsung&pl=android&ip=114.4.221.44&os=9&devid=4af219fb13029568&gaid=&imei=&country=ID&event_name=GRABFOOD_RESTAURANT%3AGRABFOOD_BASKET&event_value=%7B%7D&event_time=2019-11-29%2007%3A06%3A28.928&idfa=6201527e-4043-4432-8fa0-da68ef23f544&sha1_idfa=78a8c9ef490387a4f6a6ca3b3421b27d022048c0&md5_idfa=7a1a7617159c636566bb5438d6d070f9&sha1_gaid=78a8c9ef490387a4f6a6ca3b3421b27d022048c0&md5_gaid=7a1a7617159c636566bb5438d6d070f9&sha1_devid=90482afd849c2137d8d701d0a21de7a7f820e237&md5_devid=6b466167a89c7e3d9d13e3e8c5e45c30&sha1_imei=&md5_imei=&sha1_mac=da39a3ee5e6b4b0d3255bfef95601890afd80709&md5_mac=d41d8cd98f00b204e9800998ecf8427e&app_version=5.71.0&language=Indonesia&match_type=&is_retargeting=&appsflyer_id=1572660212085-1618710342610571481&mac=&app_version=5.71.0&sdk_version=v4.11.0&app_id=com.grabtaxi.passenger&is_wifi=false&city=Kartasura&brand=samsung&carrier=TELKOMSEL&language=Indonesia&install_time=1572658737&click_time=&gp_referrer_click_time=&gp_referrer_install_time=&is_tracking_disable=&model=SM-A605G&attribu_type=0&ua=Dalvik%2F2.1.0%20%28Linux%3B%20U%3B%20Android%209%3B%20SM-A605G%20Build%2FPPR1.180610.011%29	18.197.158.180	/event?campuuid=&clickid=&type=appsflyer&device=samsung&pl=android&ip=114.4.221.44&os=9&devid=4af219fb13029568&gaid=&imei=&country=ID&event_name=GRABFOOD_RESTAURANT%3AGRABFOOD_BASKET&event_value=%7B%7D&event_time=2019-11-29%2007%3A06%3A28.928&idfa=6201527e-4043-4432-8fa0-da68ef23f544&sha1_idfa=78a8c9ef490387a4f6a6ca3b3421b27d022048c0&md5_idfa=7a1a7617159c636566bb5438d6d070f9&sha1_gaid=78a8c9ef490387a4f6a6ca3b3421b27d022048c0&md5_gaid=7a1a7617159c636566bb5438d6d070f9&sha1_devid=90482afd849c2137d8d701d0a21de7a7f820e237&md5_devid=6b466167a89c7e3d9d13e3e8c5e45c30&sha1_imei=&md5_imei=&sha1_mac=da39a3ee5e6b4b0d3255bfef95601890afd80709&md5_mac=d41d8cd98f00b204e9800998ecf8427e&app_version=5.71.0&language=Indonesia&match_type=&is_retargeting=&appsflyer_id=1572660212085-1618710342610571481&mac=&app_version=5.71.0&sdk_version=v4.11.0&app_id=com.grabtaxi.passenger&is_wifi=false&city=Kartasura&brand=samsung&carrier=TELKOMSEL&language=Indonesia&install_time=1572658737&click_time=&gp_referrer_click_time=&gp_referrer_install_time=&is_tracking_disable=&model=SM-A605G&attribu_type=0&ua=Dalvik%2F2.1.0%20%28Linux%3B%20U%3B%20Android%209%3B%20SM-A605G%20Build%2FPPR1.180610.011%29
 */
func ReadLine(fileName string) error {//按行读取
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	tarReader, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(tarReader)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		//fmt.Println(line)
		Print(line)
	}
	return nil
}

func Print(line string) error {//获取参数
	params := strings.Fields(line)
	err := send(params[2])
	return err
}

func send(params string) error {//发送get请求
	fmt.Println("http://127.0.0.1"+params)
	res, err := http.Get("http://127.0.0.1" + params)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}
