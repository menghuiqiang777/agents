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
	agent := &agents.Agent{
		Name:         "MyAgent",
		Instructions: "Provide helpful responses",
		ModelName:    "doubao-1-5-pro-32k-250115",
		Provider:     "ARK",
	}
	agent_qwen := &agents.Agent{
		Name:         "ArkAgent",
		Instructions: "Provide helpful responses",
		ModelName:    "qwen-plus-2025-04-28",
		Provider:     "QWEN",
	}

	input := "你是谁？"
	messages := agents.CreateMessages(agent, input)
	agent.RunStream(ctx, messages)
	agent_qwen.RunStream(ctx, messages)

}


```
