//
// Copyright 2019, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package kepserverex

import (
	"fmt"
	"net/url"
)

// ChannelService handles communication with the channel related methods
// of the KEPServerEX API.
type ChannelService struct {
	client *Client
}

// Channel represents a KEPServerEX channel.
type Channel struct {
	Name                string              `json:"common.ALLTYPES_NAME"`
	Description         string              `json:"common.ALLTYPES_DESCRIPTION"`
	UniqueID            int64               `json:"servermain.CHANNEL_UNIQUE_ID"`
	ProjectID           int64               `json:"PROJECT_ID"`
	Driver              string              `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER"`
	DiagnosticsCapture  bool                `json:"servermain.CHANNEL_DIAGNOSTICS_CAPTURE"`
	NetworkAdapter      string              `json:"servermain.CHANNEL_ETHERNET_COMMUNICATIONS_NETWORK_ADAPTER_STRING"`
	OptimizationMethod  OptimizationMethod  `json:"servermain.CHANNEL_WRITE_OPTIMIZATIONS_METHOD"`
	DutyCycle           int                 `json:"servermain.CHANNEL_WRITE_OPTIMIZATIONS_DUTY_CYCLE"`
	FloatingPointValues FloatingPointValues `json:"servermain.CHANNEL_NON_NORMALIZED_FLOATING_POINT_HANDLING"`
}

// ChannelOptions represents all channel options.
type ChannelOptions struct {
	Name                string              `json:"common.ALLTYPES_NAME,omitempty"`
	Description         string              `json:"common.ALLTYPES_DESCRIPTION,omitempty"`
	Driver              string              `json:"servermain.MULTIPLE_TYPES_DEVICE_DRIVER,omitempty"`
	DiagnosticsCapture  bool                `json:"servermain.CHANNEL_DIAGNOSTICS_CAPTURE,omitempty"`
	NetworkAdapter      string              `json:"servermain.CHANNEL_SERIAL_COMMUNICATIONS_NETWORK_ADAPTER_STRING,omitempty"`
	OptimizationMethod  OptimizationMethod  `json:"servermain.CHANNEL_WRITE_OPTIMIZATIONS_METHOD,omitempty"`
	DutyCycle           int                 `json:"servermain.CHANNEL_WRITE_OPTIMIZATIONS_DUTY_CYCLE,omitempty"`
	FloatingPointValues FloatingPointValues `json:"servermain.CHANNEL_NON_NORMALIZED_FLOATING_POINT_HANDLING,omitempty"`
}

// ListChannels gets a list of channels.
func (s *ChannelService) ListChannels() ([]*Channel, error) {
	req, err := s.client.NewRequest("GET", "channels", nil)
	if err != nil {
		return nil, err
	}

	var cannels []*Channel
	if err = s.client.Do(req, &cannels); err != nil {
		return nil, err
	}

	return cannels, nil
}

// CreateChannel creates a new channel.
func (s *ChannelService) CreateChannel(options *ChannelOptions) error {
	req, err := s.client.NewRequest("POST", "channels", options)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}

// GetChannel gets a channel.
func (s *ChannelService) GetChannel(name string) (*Channel, error) {
	u := fmt.Sprintf("channels/%s", url.PathEscape(name))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c *Channel
	if err = s.client.Do(req, c); err != nil {
		return nil, err
	}

	return c, nil
}

// UpdateChannel updates an existing channel.
func (s *ChannelService) UpdateChannel(name string, options *ChannelOptions) error {
	u := fmt.Sprintf("channels/%s", url.PathEscape(name))
	req, err := s.client.NewRequest("PUT", u, options)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}

// DeleteChannel deletes a channel.
func (s *ChannelService) DeleteChannel(name string) error {
	u := fmt.Sprintf("channels/%s", url.PathEscape(name))
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}
