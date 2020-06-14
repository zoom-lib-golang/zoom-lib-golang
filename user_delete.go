package zoom

import "fmt"

// DeleteUserPath - v2 path for deleting a user
const DeleteUserPath = "/users/%s"

// DeleteUserOptions are the options to delete a user with
type DeleteUserOptions struct {
	EmailOrID         string           `url:"-"`
	Action            DeleteUserAction `url:"action,omitempty"`
	TransferEmail     string           `url:"transfer_email,omitempty"`
	TransferMeeting   bool             `url:"transfer_meeting,omitempty"`
	TransferWebinar   bool             `url:"transfer_webinar,omitempty"`
	TransferRecording bool             `url:"transfer_recording,omitempty"`
}

// DeleteUser calls DELETE /users/{userID}
func DeleteUser(opts DeleteUserOptions) error {
	return defaultClient.DeleteUser(opts)
}

// DeleteUser calls DELETE /users/{userID}
// https://marketplace.zoom.us/docs/api-reference/zoom-api/users/userdelete
func (c *Client) DeleteUser(opts DeleteUserOptions) error {
	return c.requestV2(requestV2Opts{
		Method:        Delete,
		Path:          fmt.Sprintf(DeleteUserPath, opts.EmailOrID),
		URLParameters: &opts,
		HeadResponse:  true,
	})
}
