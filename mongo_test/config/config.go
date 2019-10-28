package config

//https://github.com/widuu/goini
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var instance *Config

func init() {
	//instance = SetConfig("/Users/liyujun/programs/repos/nono/video-content-server/etc/config.ini")
	mode := os.Getenv("mode")
	var _fp string
	var fname string
	if mode == "test" {
		fname = "config_test.ini"
	} else if mode == "prod" {
		fname = "config_prod.ini"
	} else {
		fname = "config.ini"
	}
	_t, _ := filepath.Abs(".")
	fmt.Printf("root dir is %v \n", _t)
	_fp = "../etc/" + fname
	if _, err := os.Stat(_fp); err != nil && os.IsNotExist(err) {
		fmt.Printf("can't not find config file : %v, try load next dir. \n", _fp)
		_fp = "./src/etc/" + fname
		if _, err := os.Stat(_fp); err != nil && os.IsNotExist(err) {
			fmt.Printf("can't not find config file : %v, some thing wrong.please check. \n", _fp)
			os.Exit(-1)
		}
	}
	_absfp, _ := filepath.Abs(_fp)
	fmt.Printf("ready to load config file %v. \n", _absfp)
	instance = SetConfig(_fp)
	fmt.Printf("init %v file ok. \n", _absfp)
}

func GetInstance() *Config {
	return instance
}

type Config struct {
	filepath string                       //your ini file path directory+file
	conflist map[string]map[string]string //configuration information slice
}

//Create an empty configuration file
func SetConfig(filepath string) *Config {
	c := new(Config)
	c.filepath = filepath
	c.conflist = c.ReadList()

	return c
}

//To obtain corresponding value of the key values
func (c *Config) GetStringValue(section, name string, defaultValue string) string {
	if m, se := c.conflist[section]; !se {
		return defaultValue
	} else {
		if v, ne := m[name]; !ne {
			return defaultValue
		} else {
			return v
		}
	}
}

//To obtain corresponding value of the key values
func (c *Config) GetFloatValue(section, name string, defaultValue float32) float32 {
	if m, se := c.conflist[section]; !se {
		return defaultValue
	} else {
		if v, ne := m[name]; !ne {
			return defaultValue
		} else {
			f, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return defaultValue
			}
			return float32(f)
		}
	}
}

func (c *Config) GetIntValue(section, name string, defaultValue int) int {
	if m, se := c.conflist[section]; !se {
		return defaultValue
	} else {
		if v, ne := m[name]; !ne {
			return defaultValue
		} else {
			f, err := strconv.Atoi(v)
			if err != nil {
				return defaultValue
			}
			return f
		}
	}
}

func GetStringValue(section, name string, defaultValue string) string {
	return instance.GetStringValue(section, name, defaultValue)
}
func GetFloatValue(section, name string, defaultValue float32) float32 {
	return instance.GetFloatValue(section, name, defaultValue)
}
func GetIntValue(section, name string, defaultValue int) int {
	return instance.GetIntValue(section, name, defaultValue)
}

func GetGlobalStringValue(name string, defaultValue string) string {
	return instance.GetStringValue(GLOBAL_SECTION, name, defaultValue)
}
func GetGlobalFloatValue(name string, defaultValue float32) float32 {
	return instance.GetFloatValue(GLOBAL_SECTION, name, defaultValue)
}
func GetGlobalIntValue(name string, defaultValue int) int {
	return instance.GetIntValue(GLOBAL_SECTION, name, defaultValue)
}

//To obtain corresponding value of the key values
func (c *Config) GetGlobalStringValue(name string, defaultValue string) string {
	return c.GetStringValue(GLOBAL_SECTION, name, defaultValue)
}

//To obtain corresponding value of the key values
func (c *Config) GetGlobalFloatValue(name string, defaultValue float32) float32 {
	return c.GetFloatValue(GLOBAL_SECTION, name, defaultValue)
}
func (c *Config) GetGlobalIntValue(section, name string, defaultValue int) int {
	return c.GetIntValue(GLOBAL_SECTION, name, defaultValue)
}

//Set the corresponding value of the key value, if not add, if there is a key change
//func (c *Config) SetValue(section, key, value string) bool {
//	c.ReadList()
//	data := c.conflist
//	var ok bool
//	var index = make(map[int]bool)
//	var conf = make(map[string]map[string]string)
//	for i, v := range data {
//		_, ok = v[section]
//		index[i] = ok
//	}

//	i, ok := func(m map[int]bool) (i int, v bool) {
//		for i, v := range m {
//			if v == true {
//				return i, true
//			}
//		}
//		return 0, false
//	}(index)

//	if ok {
//		c.conflist[i][section][key] = value
//		return true
//	} else {
//		conf[section] = make(map[string]string)
//		conf[section][key] = value
//		c.conflist = append(c.conflist, conf)
//		return true
//	}

//	return false
//}

//Delete the corresponding key values
//func (c *Config) DeleteValue(section, name string) bool {
//	c.ReadList()
//	data := c.conflist
//	for i, v := range data {
//		for key, _ := range v {
//			if key == section {
//				delete(c.conflist[i][key], name)
//				return true
//			}
//		}
//	}
//	return false
//}

//List all the configuration file
func (c *Config) ReadList() map[string]map[string]string {
	file, err := os.Open(c.filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	var data map[string]map[string]string
	data = make(map[string]map[string]string)
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#": //增加配置文件备注
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			if _, exists := data[section]; !exists {
				data[section] = make(map[string]string)
			}
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			data[section][strings.TrimSpace(line[0:i])] = value
		}

	}
	return data

}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}

//Ban repeated appended to the slice method
func (c *Config) uniquappend(conf string) bool {
	for _, v := range c.conflist {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}
