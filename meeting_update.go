package zoom

import "fmt"

// UpdateMeetingOptions are the options to update a meeting
type UpdateMeetingOptions struct {
	MeetingID int `url:"-"`
	Body      interface{}
}

// UpdateStatusMeetingPath UpdateMeetingPath - v2 retrieve the details of a meeting
const UpdateStatusMeetingPath = "/meetings/%d/status"

// UpdateStatus calls PUT /meetings/{ID}/status
func (c *Client) UpdateStatus(opts UpdateMeetingOptions) (Meeting, error) {
	var ret = Meeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateStatusMeetingPath, opts.MeetingID),
		URLParameters:  &opts,
		DataParameters: &opts.Body,
		Ret:            &ret,
	})
}
