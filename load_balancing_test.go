package cloudflare

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{
              "description": "Primary data center - Provider XYZ",
              "name": "primary-dc-1",
              "enabled": true,
              "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
              "origins": [
                {
                  "name": "app-server-1",
                  "address": "0.0.0.0",
                  "enabled": true
                }
              ],
              "notification_email": "someone@example.com"
						}`, string(b))
		}
		fmt.Fprint(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": {
              "id": "17b5962d775c646f3f9725cbc7a53df4",
              "created_on": "2014-01-01T05:20:00.12345Z",
              "modified_on": "2014-02-01T05:20:00.12345Z",
              "description": "Primary data center - Provider XYZ",
              "name": "primary-dc-1",
              "enabled": true,
              "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
              "origins": [
                {
                  "name": "app-server-1",
                  "address": "0.0.0.0",
                  "enabled": true
                }
              ],
              "notification_email": "someone@example.com"
            }
        }`)
	}

	mux.HandleFunc("/user/load_balancers/pools", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := LoadBalancerPool{
		ID:          "17b5962d775c646f3f9725cbc7a53df4",
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     true,
		Monitor:     "f1aba936b94213e5b8dca0c0dbf1f9cc",
		Origins: []LoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
			},
		},
		NotificationEmail: "someone@example.com",
	}
	request := LoadBalancerPool{
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     true,
		Monitor:     "f1aba936b94213e5b8dca0c0dbf1f9cc",
		Origins: []LoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
			},
		},
		NotificationEmail: "someone@example.com",
	}

	actual, err := client.CreateLoadBalancerPool(request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListLoadBalancerPools(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": [
                {
                    "id": "17b5962d775c646f3f9725cbc7a53df4",
                    "created_on": "2014-01-01T05:20:00.12345Z",
                    "modified_on": "2014-02-01T05:20:00.12345Z",
                    "description": "Primary data center - Provider XYZ",
                    "name": "primary-dc-1",
                    "enabled": true,
                    "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
                    "origins": [
                      {
                        "name": "app-server-1",
                        "address": "0.0.0.0",
                        "enabled": true
                      }
                    ],
                    "notification_email": "someone@example.com"
                }
            ],
            "result_info": {
                "page": 1,
                "per_page": 20,
                "count": 1,
                "total_count": 2000
            }
        }`)
	}

	mux.HandleFunc("/user/load_balancers/pools", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := []LoadBalancerPool{
		{
			ID:          "17b5962d775c646f3f9725cbc7a53df4",
			CreatedOn:   &createdOn,
			ModifiedOn:  &modifiedOn,
			Description: "Primary data center - Provider XYZ",
			Name:        "primary-dc-1",
			Enabled:     true,
			Monitor:     "f1aba936b94213e5b8dca0c0dbf1f9cc",
			Origins: []LoadBalancerOrigin{
				{
					Name:    "app-server-1",
					Address: "0.0.0.0",
					Enabled: true,
				},
			},
			NotificationEmail: "someone@example.com",
		},
	}

	actual, err := client.ListLoadBalancerPools()
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestLoadBalancerPoolDetails(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": {
              "id": "17b5962d775c646f3f9725cbc7a53df4",
              "created_on": "2014-01-01T05:20:00.12345Z",
              "modified_on": "2014-02-01T05:20:00.12345Z",
              "description": "Primary data center - Provider XYZ",
              "name": "primary-dc-1",
              "enabled": true,
              "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
              "origins": [
                {
                  "name": "app-server-1",
                  "address": "0.0.0.0",
                  "enabled": true
                }
              ],
              "notification_email": "someone@example.com"
            }
        }`)
	}

	mux.HandleFunc("/user/load_balancers/pools/17b5962d775c646f3f9725cbc7a53df4", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := LoadBalancerPool{
		ID:          "17b5962d775c646f3f9725cbc7a53df4",
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     true,
		Monitor:     "f1aba936b94213e5b8dca0c0dbf1f9cc",
		Origins: []LoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
			},
		},
		NotificationEmail: "someone@example.com",
	}

	actual, err := client.LoadBalancerPoolDetails("17b5962d775c646f3f9725cbc7a53df4")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	_, err = client.LoadBalancerPoolDetails("bar")
	assert.Error(t, err)
}

func TestDeleteLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": {
              "id": "17b5962d775c646f3f9725cbc7a53df4"
            }
        }`)
	}

	mux.HandleFunc("/user/load_balancers/pools/17b5962d775c646f3f9725cbc7a53df4", handler)
	assert.NoError(t, client.DeleteLoadBalancerPool("17b5962d775c646f3f9725cbc7a53df4"))
	assert.Error(t, client.DeleteLoadBalancerPool("bar"))
}

func TestModifyLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "PUT", "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{
              "id": "17b5962d775c646f3f9725cbc7a53df4",
              "description": "Primary data center - Provider XYZZY",
              "name": "primary-dc-2",
              "enabled": false,
              "origins": [
                {
                  "name": "app-server-2",
                  "address": "0.0.0.1",
                  "enabled": false
                }
              ],
              "notification_email": "nobody@example.com"
						}`, string(b))
		}
		fmt.Fprint(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": {
              "id": "17b5962d775c646f3f9725cbc7a53df4",
              "created_on": "2014-01-01T05:20:00.12345Z",
              "modified_on": "2017-02-01T05:20:00.12345Z",
              "description": "Primary data center - Provider XYZZY",
              "name": "primary-dc-2",
              "enabled": false,
              "origins": [
                {
                  "name": "app-server-2",
                  "address": "0.0.0.1",
                  "enabled": false
                }
              ],
              "notification_email": "nobody@example.com"
            }
        }`)
	}

	mux.HandleFunc("/user/load_balancers/pools/17b5962d775c646f3f9725cbc7a53df4", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2017-02-01T05:20:00.12345Z")
	want := LoadBalancerPool{
		ID:          "17b5962d775c646f3f9725cbc7a53df4",
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Description: "Primary data center - Provider XYZZY",
		Name:        "primary-dc-2",
		Enabled:     false,
		Origins: []LoadBalancerOrigin{
			{
				Name:    "app-server-2",
				Address: "0.0.0.1",
				Enabled: false,
			},
		},
		NotificationEmail: "nobody@example.com",
	}
	request := LoadBalancerPool{
		ID:          "17b5962d775c646f3f9725cbc7a53df4",
		Description: "Primary data center - Provider XYZZY",
		Name:        "primary-dc-2",
		Enabled:     false,
		Origins: []LoadBalancerOrigin{
			{
				Name:    "app-server-2",
				Address: "0.0.0.1",
				Enabled: false,
			},
		},
		NotificationEmail: "nobody@example.com",
	}

	actual, err := client.ModifyLoadBalancerPool(request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
