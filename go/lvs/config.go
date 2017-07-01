package lvs

import (
	"fmt"
	"strings"
)

//ConfigBlock 配置块结构
type ConfigBlock map[string]interface{}

//Unmarshal 解析配置
func Unmarshal(config string) (*ConfigBlock, error) {

	block := ConfigBlock{}

	r := strings.NewReader(config)

	contextStack := []*ConfigBlock{&block}

	previous := ""
	current := ""

	serverID := -1
	for r.Len() > 0 {
		char, _, _ := r.ReadRune()
		m := string(char)
		switch m {
		case "\t":
			fallthrough
		case " ": //当前字符为空
			if current == "" {
				continue
			}
			//得到了一个key，或者value
			previous = current
			current = ""
		case "{": //遇到{进行切换到下部context
			newConfigBlock := &ConfigBlock{}
			if previous == "real_server" {
				serverID++
				previous = fmt.Sprintf("real_server_%d", serverID)
			}
			(*contextStack[len(contextStack)-1])[previous] = newConfigBlock
			contextStack = append(contextStack, newConfigBlock)
			previous = ""
			current = ""
		case "}": //遇到}进行切换回上层context
			contextStack = contextStack[:len(contextStack)-1]
			previous = ""
			current = ""
		case "\n": //遇到换行了
			if previous != "" && current != "" {
				(*contextStack[len(contextStack)-1])[previous] = strings.Trim(current, `'"`)
			}
			previous = ""
			current = ""
		default:
			current += m
		}

	}

	//转换real_server 为数组
	if serverID > -1 {
		serverBlocks := []interface{}{}
		vertualServer := block["virtual_server"].(*ConfigBlock)
		for i := 0; i <= serverID; i++ {
			key := fmt.Sprintf("real_server_%d", i)
			serverBlocks = append(serverBlocks, (*vertualServer)[key])
			delete(*vertualServer, key)
		}
		(*vertualServer)["real_server"] = serverBlocks
	}

	return &block, nil
}
