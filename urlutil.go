/*
* Package urlutil supplies utility methods for url.URL
*
* WIP Warning
 */
package urlutil

import "net/url"

// Set sets the key/value pair for the URL Query values
func Set(u *url.URL, key, value string) {
	v := u.Query()
	v.Set(key, value)
	u.RawQuery = v.Encode()
}

// Add adds a key/value pair for the URL Query values
func Add(u *url.URL, key, value string) {
	v := u.Query()
	v.Add(key, value)
	u.RawQuery = v.Encode()
}

// Del deletes the values associated to the key for the URL Query values
func Del(u *url.URL, key string) {
	v := u.Query()
	v.Del(key)
	u.RawQuery = v.Encode()
}
