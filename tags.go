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

// TagService handles communication with the tag related methods
// of the KEPServerEX API.
type TagService struct {
	client *Client
}

// Tag represents a tag.
type Tag struct {
	Name           string         `json:"common.ALLTYPES_NAME"`
	Description    string         `json:"common.ALLTYPES_DESCRIPTION"`
	ProjectID      int64          `json:"PROJECT_ID"`
	Address        string         `json:"servermain.TAG_ADDRESS"`
	DataType       DataType       `json:"servermain.TAG_DATA_TYPE"`
	ClientAccess   ClientAccess   `json:"servermain.TAG_READ_WRITE_ACCESS"`
	ScanRate       int            `json:"servermain.TAG_SCAN_RATE_MILLISECONDS"`
	Autogenerated  bool           `json:"servermain.TAG_AUTOGENERATED"`
	ScalingType    ScalingType    `json:"servermain.TAG_SCALING_TYPE"`
	RawLow         int            `json:"servermain.TAG_SCALING_RAW_LOW"`
	RawHigh        int            `json:"servermain.TAG_SCALING_RAW_HIGH"`
	ScaledDataType ScaledDataType `json:"servermain.TAG_SCALING_SCALED_DATA_TYPE"`
	ScaledLow      int            `json:"servermain.TAG_SCALING_SCALED_LOW"`
	ScaledHigh     int            `json:"servermain.TAG_SCALING_SCALED_HIGH"`
	ClampLow       bool           `json:"servermain.TAG_SCALING_CLAMP_LOW"`
	ClampHigh      bool           `json:"servermain.TAG_SCALING_CLAMP_HIGH"`
	NegateValue    bool           `json:"servermain.TAG_SCALING_NEGATE_VALUE"`
	Units          string         `json:"servermain.TAG_SCALING_UNITS"`
}

// TagOptions represents all tag options.
type TagOptions struct {
	Name           *string         `json:"common.ALLTYPES_NAME,omitempty"`
	Description    *string         `json:"common.ALLTYPES_DESCRIPTION,omitempty"`
	Address        *string         `json:"servermain.TAG_ADDRESS,omitempty"`
	DataType       *DataType       `json:"servermain.TAG_DATA_TYPE,omitempty"`
	ClientAccess   *ClientAccess   `json:"servermain.TAG_READ_WRITE_ACCESS,omitempty"`
	ScanRate       *int            `json:"servermain.TAG_SCAN_RATE_MILLISECONDS,omitempty"`
	Autogenerated  *bool           `json:"servermain.TAG_AUTOGENERATED,omitempty"`
	ScalingType    *ScalingType    `json:"servermain.TAG_SCALING_TYPE,omitempty"`
	RawLow         *int            `json:"servermain.TAG_SCALING_RAW_LOW,omitempty"`
	RawHigh        *int            `json:"servermain.TAG_SCALING_RAW_HIGH,omitempty"`
	ScaledDataType *ScaledDataType `json:"servermain.TAG_SCALING_SCALED_DATA_TYPE,omitempty"`
	ScaledLow      *int            `json:"servermain.TAG_SCALING_SCALED_LOW,omitempty"`
	ScaledHigh     *int            `json:"servermain.TAG_SCALING_SCALED_HIGH,omitempty"`
	ClampLow       *bool           `json:"servermain.TAG_SCALING_CLAMP_LOW,omitempty"`
	ClampHigh      *bool           `json:"servermain.TAG_SCALING_CLAMP_HIGH,omitempty"`
	NegateValue    *bool           `json:"servermain.TAG_SCALING_NEGATE_VALUE,omitempty"`
	Units          *string         `json:"servermain.TAG_SCALING_UNITS,omitempty"`
}

// ListTags gets a list of tags.
func (s *TagService) ListTags(channel, device, group string) ([]*Tag, error) {
	u := fmt.Sprintf("channels/%s/devices/%s/tag_groups/%s/tags",
		url.PathEscape(channel),
		url.PathEscape(device),
		url.PathEscape(group),
	)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var tags []*Tag
	if err = s.client.Do(req, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

// CreateTag creates a new tag.
func (s *TagService) CreateTag(channel, device, group string, options *TagOptions) error {
	u := fmt.Sprintf("channels/%s/devices/%s/tag_groups/%s/tags",
		url.PathEscape(channel),
		url.PathEscape(device),
		url.PathEscape(group),
	)
	req, err := s.client.NewRequest("POST", u, options)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}

// GetTag gets a specific tag.
func (s *TagService) GetTag(channel, device, group, name string) (*Tag, error) {
	u := fmt.Sprintf("channels/%s/devices/%s/tag_groups/%s/tags/%s",
		url.PathEscape(channel),
		url.PathEscape(device),
		url.PathEscape(group),
		url.PathEscape(name),
	)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var tag *Tag
	if err = s.client.Do(req, tag); err != nil {
		return nil, err
	}

	return tag, nil
}

// UpdateTag updates an existing tag.
func (s *TagService) UpdateTag(channel, device, group, name string, options *TagOptions) error {
	u := fmt.Sprintf("channels/%s/devices/%s/tag_groups/%s/tags/%s",
		url.PathEscape(channel),
		url.PathEscape(device),
		url.PathEscape(group),
		url.PathEscape(name),
	)
	req, err := s.client.NewRequest("PUT", u, options)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}

// DeleteTag deletes a tag.
func (s *TagService) DeleteTag(channel, device, group, name string) error {
	u := fmt.Sprintf("channels/%s/devices/%s/tag_groups/%s/tags/%s",
		url.PathEscape(channel),
		url.PathEscape(device),
		url.PathEscape(group),
		url.PathEscape(name),
	)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}
	return s.client.Do(req, nil)
}