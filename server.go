package sonycrapi

import "encoding/json"

// GetAvailableAPIList obtains a list of available API features from the camera
func (c *Camera) GetAvailableAPIList() (available []string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getAvailableApiList").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &available)
	}

	return
}

// GetApplicationInfo obtains the name and version of the application server
// on the camera.
func (c *Camera) GetApplicationInfo() (server string, version string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getApplicationInfo").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 1 {
		err = json.Unmarshal(resp.Result[0], &server)
		if err != nil {
			return
		}
		err = json.Unmarshal(resp.Result[1], &version)
	}

	return
}

// GetVersions obtains the supported versions of the API server at the given endpoint
func (c *Camera) GetVersions(ep endpoint) (versions []string, err error) {
	resp, err := c.newRequest(ep, "getVersions").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &versions)
	}

	return
}

// GetMethodTypes obtains the types for each of the given API functions for a given endpoint.
// An ampty string for version will return types for all versions
func (c *Camera) GetMethodTypes(ep endpoint, version string) (types [][]string, err error) {
	resp, err := c.newRequest(ep, "getMethodTypes", version).Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &types)
	}

	return
}
