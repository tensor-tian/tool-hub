package cmd

import (
	"regexp"
	"strings"
)

// escapeShellValue 转义单个值用于 shell 单引号字符串
// 将值转换为 'value' 格式，处理内部的单引号
func escapeShellValue(value string) string {
	// 替换单引号: ' -> '\''
	escaped := strings.ReplaceAll(value, "'", "'\\''")
	return "'" + escaped + "'"
}

// escapeShellScript 转义整个脚本用于 bash -c
func escapeShellScript(script string) string {
	escaped := strings.ReplaceAll(script, "'", "'\\''")
	return "'" + escaped + "'"
}

// replaceVariables 替换模板中的变量
// 支持 $varName 和 ${varName} 两种格式
func replaceVariables(template string, args map[string]string) string {
	script := template

	// 正则匹配 $varName 或 ${varName}
	// 使用两个独立的正则更清晰

	// 1. 替换 ${varName} 格式
	re1 := regexp.MustCompile(`\$\{([a-zA-Z_][a-zA-Z0-9_]*)\}`)
	script = re1.ReplaceAllStringFunc(script, func(match string) string {
		// 提取变量名
		varName := re1.FindStringSubmatch(match)[1]
		if value, ok := args[varName]; ok {
			return escapeShellValue(value)
		}
		// 如果找不到变量，保持原样（或者返回错误）
		return match
	})

	// 2. 替换 $varName 格式（但要避免 $1, $2 等位置参数）
	// 这个正则确保不会匹配数字开头的变量
	re2 := regexp.MustCompile(`\$([a-zA-Z_][a-zA-Z0-9_]*)`)
	script = re2.ReplaceAllStringFunc(script, func(match string) string {
		varName := re2.FindStringSubmatch(match)[1]
		if value, ok := args[varName]; ok {
			return escapeShellValue(value)
		}
		return match
	})

	return script
}

// BuildBashCommand constructs a bash command with the given template and arguments, safely escaping variables.
func BuildBashCommand(sh string, template string, args map[string]string) []string {
	scriptWithVars := replaceVariables(template, args)
	return []string{sh, "-c", scriptWithVars}
}
