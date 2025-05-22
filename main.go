/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package agents

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	// 使用模版创建messages
	log.Printf("===create messages===\n")

	// 创建 Agent 实例
	agent := &Agent{
		Name:         "ArkAgent",
		Instructions: "Provide helpful responses",
		ModelName:    "doubao-1-5-pro-32k-250115",
		Provider:     "ARK",
	}

	agent_qwen := &Agent{
		Name:         "ArkAgent",
		Instructions: "Provide helpful responses",
		ModelName:    "qwen-plus-2025-04-28",
		Provider:     "QWEN",
	}

	log.Printf("Created agent: %+v\n", agent)
	input := "你是谁？"
	messages := CreateMessages(agent, input)
	log.Printf("messages: %+v\n\n", messages)
	//fmt.Println("messages: ", messages)

	agent.RunStream(ctx, messages)

	agent_qwen.RunStream(ctx, messages)
}
