// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// V2ListSupportedOpenshiftVersionsURL generates an URL for the v2 list supported openshift versions operation
type V2ListSupportedOpenshiftVersionsURL struct {
	OnlyLatest *bool
	Version    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2ListSupportedOpenshiftVersionsURL) WithBasePath(bp string) *V2ListSupportedOpenshiftVersionsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2ListSupportedOpenshiftVersionsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *V2ListSupportedOpenshiftVersionsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/openshift-versions"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/assisted-install"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var onlyLatestQ string
	if o.OnlyLatest != nil {
		onlyLatestQ = swag.FormatBool(*o.OnlyLatest)
	}
	if onlyLatestQ != "" {
		qs.Set("only_latest", onlyLatestQ)
	}

	var versionQ string
	if o.Version != nil {
		versionQ = *o.Version
	}
	if versionQ != "" {
		qs.Set("version", versionQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *V2ListSupportedOpenshiftVersionsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *V2ListSupportedOpenshiftVersionsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *V2ListSupportedOpenshiftVersionsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on V2ListSupportedOpenshiftVersionsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on V2ListSupportedOpenshiftVersionsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *V2ListSupportedOpenshiftVersionsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
