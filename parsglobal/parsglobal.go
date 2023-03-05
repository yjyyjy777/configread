package parsglobal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	JdbcUrl      string `mapstructure:"jdbc.url"`
	DatabaseHost string `mapstructure:"database.host"`
	DatabasePort string `mapstructure:"database.port"`
	DatabaseName string `mapstructure:"database.name"`
	DBUserName   string `json:"db_username"`
	DBPasswd     string `json:"db_passwd"`
}

func Parsjdbcurl(url string) (string, string, string) {
	url = strings.ReplaceAll(url, `\`, "")

	re := regexp.MustCompile(`\/\/([^:]+):(\d+)\/(.+)\?`)
	match := re.FindStringSubmatch(url)
	if len(match) < 4 {
		fmt.Println("Unable to parse URL:", url)

	}

	// 提取IP地址、端口号和数据库名称
	ip := match[1]
	port := match[2]
	dbName := match[3]
	return ip, port, dbName
}

func Globalconfig() Config {
	file, err := os.Open("global.properties")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	Config := Config{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 && line[0] != '#' {
			parts := strings.Split(line, "=")

			if len(parts) >= 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				//fmt.Println(key, "///", value)
				if key == "jdbc.url" {
					Config.JdbcUrl = value
				} else if key == "database.host" {
					Config.DatabaseHost = value
				} else if key == "database.port" {
					Config.DatabasePort = value
				} else if key == "database.name" {
					Config.DatabaseName = value
				} else if key == "jdbc.username" {
					Config.DBUserName = value
				} else if key == "jdbc.password" {
					Config.DBPasswd = value
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	Config.DatabaseHost, Config.DatabasePort, Config.DatabaseName = Parsjdbcurl(Config.JdbcUrl)

	return Config

}
