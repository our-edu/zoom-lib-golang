package zoom

import (
	"fmt"
)

// UpdateMeetingStatusOptions are the options to update a meeting
type UpdateMeetingStatusOptions struct {
	MeetingID int    `url:"-"`
	Action    string `json:"action,omitempty"`
}

// UpdateMeetingStatusPath UpdateMeetingPath - v2 retrieve the details of a meeting
const UpdateMeetingStatusPath = "/meetings/%d/status"

// UpdateStatus calls PUT /meetings/{ID}/status
func UpdateStatus(opts UpdateMeetingStatusOptions) error {
	return defaultClient.UpdateStatus(opts)
}

// UpdateStatus calls PUT /meetings/{ID}/status
func (c *Client) UpdateStatus(opts UpdateMeetingStatusOptions) error {
	return c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateMeetingStatusPath, opts.MeetingID),
		DataParameters: &opts,
		HeadResponse:   true,
	})
}
