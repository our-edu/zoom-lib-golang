package zoom

import (
	"fmt"
)

// GetUserPath - v2 path for getting a specific user
const GetUserPath = "/users/%s"

// GetUserOpts contains options for GetUser
type GetUserOpts struct {
	EmailOrID string         `url:"-"`
	LoginType *UserLoginType `url:"login_type,omitempty"` // use pointer so it can be null
}

// GetUser calls /users/{userId}, searching for a user by ID or email, using the default client
func GetUser(opts GetUserOpts) (User, error) {
	return defaultClient.GetUser(opts)
}

// GetUser calls /users/{userId}, searching for a user by ID or email, using a specific client
func (c *Client) GetUser(opts GetUserOpts) (User, error) {
	var ret = User{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetUserPath, opts.EmailOrID),
		URLParameters: opts,
		Ret:           &ret,
	})
}

// GetUserTokenOpts contains options for GetUserToken
type GetUserTokenOpts struct {
	Id   string `url:"-"`
	Type string `url:"type,omitempty"`
	TTL  int    `url:"ttl,omitempty"`
}

// TokenResponse represents response of GetUserToken()
type TokenResponse struct {
	token string
}

// GetUserToken calls /users/{userId}/token
func GetUserToken(opt GetUserTokenOpts) (string, error) {
	var ret map[string]interface{}
	err := defaultClient.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf("/users/%s/token", opt.Id),
		URLParameters: opt,
		Ret:           &ret,
	})

	if err != nil {
		fmt.Printf("=========== ZAK err %+v\n", err)
	}

	return ret["token"].(string), err
}
