// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateAIPromptDecoratorPluginRole string

const (
	CreateAIPromptDecoratorPluginRoleSystem    CreateAIPromptDecoratorPluginRole = "system"
	CreateAIPromptDecoratorPluginRoleAssistant CreateAIPromptDecoratorPluginRole = "assistant"
	CreateAIPromptDecoratorPluginRoleUser      CreateAIPromptDecoratorPluginRole = "user"
)

func (e CreateAIPromptDecoratorPluginRole) ToPointer() *CreateAIPromptDecoratorPluginRole {
	return &e
}
func (e *CreateAIPromptDecoratorPluginRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system":
		fallthrough
	case "assistant":
		fallthrough
	case "user":
		*e = CreateAIPromptDecoratorPluginRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIPromptDecoratorPluginRole: %v", v)
	}
}

type CreateAIPromptDecoratorPluginAppend struct {
	Content string                             `json:"content"`
	Role    *CreateAIPromptDecoratorPluginRole `json:"role,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginAppend) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *CreateAIPromptDecoratorPluginAppend) GetRole() *CreateAIPromptDecoratorPluginRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type CreateAIPromptDecoratorPluginConfigRole string

const (
	CreateAIPromptDecoratorPluginConfigRoleSystem    CreateAIPromptDecoratorPluginConfigRole = "system"
	CreateAIPromptDecoratorPluginConfigRoleAssistant CreateAIPromptDecoratorPluginConfigRole = "assistant"
	CreateAIPromptDecoratorPluginConfigRoleUser      CreateAIPromptDecoratorPluginConfigRole = "user"
)

func (e CreateAIPromptDecoratorPluginConfigRole) ToPointer() *CreateAIPromptDecoratorPluginConfigRole {
	return &e
}
func (e *CreateAIPromptDecoratorPluginConfigRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system":
		fallthrough
	case "assistant":
		fallthrough
	case "user":
		*e = CreateAIPromptDecoratorPluginConfigRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIPromptDecoratorPluginConfigRole: %v", v)
	}
}

type CreateAIPromptDecoratorPluginPrepend struct {
	Content string                                   `json:"content"`
	Role    *CreateAIPromptDecoratorPluginConfigRole `json:"role,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginPrepend) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *CreateAIPromptDecoratorPluginPrepend) GetRole() *CreateAIPromptDecoratorPluginConfigRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type CreateAIPromptDecoratorPluginPrompts struct {
	// Insert chat messages at the end of the chat message array. This array preserves exact order when adding messages.
	Append []CreateAIPromptDecoratorPluginAppend `json:"append,omitempty"`
	// Insert chat messages at the beginning of the chat message array. This array preserves exact order when adding messages.
	Prepend []CreateAIPromptDecoratorPluginPrepend `json:"prepend,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginPrompts) GetAppend() []CreateAIPromptDecoratorPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *CreateAIPromptDecoratorPluginPrompts) GetPrepend() []CreateAIPromptDecoratorPluginPrepend {
	if o == nil {
		return nil
	}
	return o.Prepend
}

type CreateAIPromptDecoratorPluginConfig struct {
	Prompts *CreateAIPromptDecoratorPluginPrompts `json:"prompts,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginConfig) GetPrompts() *CreateAIPromptDecoratorPluginPrompts {
	if o == nil {
		return nil
	}
	return o.Prompts
}

type CreateAIPromptDecoratorPluginProtocols string

const (
	CreateAIPromptDecoratorPluginProtocolsGrpc           CreateAIPromptDecoratorPluginProtocols = "grpc"
	CreateAIPromptDecoratorPluginProtocolsGrpcs          CreateAIPromptDecoratorPluginProtocols = "grpcs"
	CreateAIPromptDecoratorPluginProtocolsHTTP           CreateAIPromptDecoratorPluginProtocols = "http"
	CreateAIPromptDecoratorPluginProtocolsHTTPS          CreateAIPromptDecoratorPluginProtocols = "https"
	CreateAIPromptDecoratorPluginProtocolsTCP            CreateAIPromptDecoratorPluginProtocols = "tcp"
	CreateAIPromptDecoratorPluginProtocolsTLS            CreateAIPromptDecoratorPluginProtocols = "tls"
	CreateAIPromptDecoratorPluginProtocolsTLSPassthrough CreateAIPromptDecoratorPluginProtocols = "tls_passthrough"
	CreateAIPromptDecoratorPluginProtocolsUDP            CreateAIPromptDecoratorPluginProtocols = "udp"
	CreateAIPromptDecoratorPluginProtocolsWs             CreateAIPromptDecoratorPluginProtocols = "ws"
	CreateAIPromptDecoratorPluginProtocolsWss            CreateAIPromptDecoratorPluginProtocols = "wss"
)

func (e CreateAIPromptDecoratorPluginProtocols) ToPointer() *CreateAIPromptDecoratorPluginProtocols {
	return &e
}
func (e *CreateAIPromptDecoratorPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = CreateAIPromptDecoratorPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIPromptDecoratorPluginProtocols: %v", v)
	}
}

// CreateAIPromptDecoratorPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateAIPromptDecoratorPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAIPromptDecoratorPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAIPromptDecoratorPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateAIPromptDecoratorPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAIPromptDecoratorPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateAIPromptDecoratorPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIPromptDecoratorPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAIPromptDecoratorPlugin struct {
	Config *CreateAIPromptDecoratorPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"ai-prompt-decorator" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateAIPromptDecoratorPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateAIPromptDecoratorPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateAIPromptDecoratorPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateAIPromptDecoratorPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateAIPromptDecoratorPluginService `json:"service,omitempty"`
}

func (c CreateAIPromptDecoratorPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAIPromptDecoratorPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAIPromptDecoratorPlugin) GetConfig() *CreateAIPromptDecoratorPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateAIPromptDecoratorPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateAIPromptDecoratorPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateAIPromptDecoratorPlugin) GetName() *string {
	return types.String("ai-prompt-decorator")
}

func (o *CreateAIPromptDecoratorPlugin) GetProtocols() []CreateAIPromptDecoratorPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateAIPromptDecoratorPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateAIPromptDecoratorPlugin) GetConsumer() *CreateAIPromptDecoratorPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateAIPromptDecoratorPlugin) GetConsumerGroup() *CreateAIPromptDecoratorPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateAIPromptDecoratorPlugin) GetRoute() *CreateAIPromptDecoratorPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateAIPromptDecoratorPlugin) GetService() *CreateAIPromptDecoratorPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
