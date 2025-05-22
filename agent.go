package agents

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino-ext/components/model/qwen"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// Agent 结构体定义，增加 modelName 和 provider 字段
type Agent struct {
	name         string
	instructions string
	modelName    string
	provider     string
}

// newModel 根据 provider 创建不同的模型
func (a *Agent) newModel(ctx context.Context) (interface{}, error) {
	switch a.provider {
	case "ARK":
		return ark.NewChatModel(ctx, &ark.ChatModelConfig{
			APIKey: os.Getenv("ARK_API_KEY"),
			Model:  a.modelName,
		})
	case "QWEN":
		return qwen.NewChatModel(ctx, &qwen.ChatModelConfig{
			// 假设 QWEN 模型需要类似的配置，具体根据实际情况调整
			BaseURL: "https://dashscope.aliyuncs.com/compatible-mode/v1",
			APIKey:  os.Getenv("QWEN_API_KEY"),
			Model:   a.modelName,
		})
	default:
		return nil, fmt.Errorf("unsupported provider: %s", a.provider)
	}
}

// run 方法，调用 generate.go 中的 generate 方法
func (a *Agent) run(ctx context.Context, in []*schema.Message) *schema.Message {
	model, err := a.newModel(ctx)
	if err != nil {
		log.Fatalf("create chat model failed, err=%v", err)
	}

	var result *schema.Message
	switch m := model.(type) {
	case *ark.ChatModel:
		result, err = m.Generate(ctx, in)
	case *qwen.ChatModel:
		result, err = m.Generate(ctx, in)
	default:
		log.Fatalf("unsupported model type")
	}

	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}

// run_stream 方法，调用 stream.go 中的 reportStream 方法
func (a *Agent) run_stream(ctx context.Context, in []*schema.Message) {
	model, err := a.newModel(ctx)
	if err != nil {
		log.Fatalf("create chat model failed, err=%v", err)
	}

	var sr *schema.StreamReader[*schema.Message]
	var errStream error
	switch m := model.(type) {
	case *ark.ChatModel:
		sr, errStream = m.Stream(ctx, in)
	case *qwen.ChatModel:
		sr, errStream = m.Stream(ctx, in)
	default:
		log.Fatalf("unsupported model type")
	}

	if errStream != nil {
		log.Fatalf("llm stream generate failed: %v", errStream)
	}

	reportStream(sr)
}

func reportStream(sr *schema.StreamReader[*schema.Message]) {
	defer sr.Close()

	for {
		message, err := sr.Recv()
		if err == io.EOF {
			fmt.Println()
			return
		}
		if err != nil {
			log.Fatalf("recv failed: %v", err)
		}
		fmt.Print(message.Content)
	}
}

func createMessages(agent *Agent, input string) []*schema.Message {
	// 创建模板，使用 FString 格式
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板，使用 agent 的 instructions 字段
		schema.SystemMessage(agent.instructions),
		// 用户消息模板
		schema.UserMessage("{input}"),
	)

	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"input": input,
	})
	if err != nil {
		log.Fatalf("format template failed: %v\n", err)
	}
	return messages
}

// 原 createTemplate 和 createMessagesFromTemplate 函数可删除
