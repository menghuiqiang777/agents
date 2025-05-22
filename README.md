#智能体引用

/etc/environment中定义火山引擎及通义模型的apiKey

```bash

QWEN_API_KEY="sk-xxxxxx"

ARK_API_KEY="xxxxx"

```

引用样例

```go
package main

import (
	"context"

	"github.com/menghuiqiang777/agents"
)

func main() {
	ctx := context.Background()

	// 创建 Agent 实例，使用新的工厂函数
	agent := agents.NewAgent("ArkAgent", "Provide helpful responses", "doubao-1-5-pro-32k-250115")

	agent_qwen := agents.NewAgent("QwenAgent", "Provide helpful responses", "qwen-plus-2025-04-28", "QWEN")

	input := "你是谁？"
	messages := agents.CreateMessages(agent, input)
	agent.RunStream(ctx, messages)
	agent_qwen.RunStream(ctx, messages)

}


```
