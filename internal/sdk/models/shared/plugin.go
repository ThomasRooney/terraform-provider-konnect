// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Protocols string

const (
	ProtocolsGrpc           Protocols = "grpc"
	ProtocolsGrpcs          Protocols = "grpcs"
	ProtocolsHTTP           Protocols = "http"
	ProtocolsHTTPS          Protocols = "https"
	ProtocolsTCP            Protocols = "tcp"
	ProtocolsTLS            Protocols = "tls"
	ProtocolsTLSPassthrough Protocols = "tls_passthrough"
	ProtocolsUDP            Protocols = "udp"
	ProtocolsWs             Protocols = "ws"
	ProtocolsWss            Protocols = "wss"
)

func (e Protocols) ToPointer() *Protocols {
	return &e
}
func (e *Protocols) UnmarshalJSON(data []byte) error {
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
		*e = Protocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Protocols: %v", v)
	}
}

// PluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type PluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *PluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *PluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type PluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *PluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type PluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *PluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Plugin struct {
	// The configuration properties for the Plugin which can be found on the plugins documentation page in the [Kong Hub](https://docs.konghq.com/hub/).
	Config any `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	// The name of the Plugin that's going to be added. Currently, the Plugin must be installed in every Kong instance separately.
	Name *string `json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []Protocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *PluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *PluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *PluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *PluginService `json:"service,omitempty"`
}

func (o *Plugin) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Plugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Plugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Plugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Plugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *Plugin) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Plugin) GetProtocols() []Protocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *Plugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Plugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Plugin) GetConsumer() *PluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *Plugin) GetConsumerGroup() *PluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *Plugin) GetRoute() *PluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *Plugin) GetService() *PluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

type PluginInput struct {
	// The configuration properties for the Plugin which can be found on the plugins documentation page in the [Kong Hub](https://docs.konghq.com/hub/).
	Config any `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	// The name of the Plugin that's going to be added. Currently, the Plugin must be installed in every Kong instance separately.
	Name *string `json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []Protocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *PluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *PluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *PluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *PluginService `json:"service,omitempty"`
}

func (o *PluginInput) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *PluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *PluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *PluginInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PluginInput) GetProtocols() []Protocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *PluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PluginInput) GetConsumer() *PluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *PluginInput) GetConsumerGroup() *PluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *PluginInput) GetRoute() *PluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *PluginInput) GetService() *PluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
