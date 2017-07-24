package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// LoadBalancerPool represents a load balancer pool's properties.
type LoadBalancerPool struct {
	ID                string               `json:"id,omitempty"`
	CreatedOn         *time.Time           `json:"created_on,omitempty"`
	ModifiedOn        *time.Time           `json:"modified_on,omitempty"`
	Description       string               `json:"description"`
	Name              string               `json:"name"`
	Enabled           bool                 `json:"enabled"`
	Monitor           string               `json:"monitor,omitempty"`
	Origins           []LoadBalancerOrigin `json:"origins"`
	NotificationEmail string               `json:"notification_email,omitempty"`
}

type LoadBalancerOrigin struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Enabled bool   `json:"enabled"`
}

// loadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type loadBalancerPoolResponse struct {
	Response
	Result LoadBalancerPool `json:"result"`
}

// loadBalancerPoolListResponse represents the response from the List Pools endpoint.
type loadBalancerPoolListResponse struct {
	Response
	Result     []LoadBalancerPool `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// CreateLoadBalancerPool creates a new load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-create-a-pool
func (api *API) CreateLoadBalancerPool(pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := "/user/load_balancers/pools"
	res, err := api.makeRequest("POST", uri, pool)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-list-pools
func (api *API) ListLoadBalancerPools() ([]LoadBalancerPool, error) {
	uri := "/user/load_balancers/pools"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerPoolDetails returns the details for a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-details
func (api *API) LoadBalancerPoolDetails(poolID string) (LoadBalancerPool, error) {
	uri := "/user/load_balancers/pools/" + poolID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerPool disables and deletes a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-delete-a-pool
func (api *API) DeleteLoadBalancerPool(poolID string) error {
	uri := "/user/load_balancers/pools/" + poolID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// ModifyLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-modify-a-pool
func (api *API) ModifyLoadBalancerPool(pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := "/user/load_balancers/pools/" + pool.ID
	res, err := api.makeRequest("PUT", uri, pool)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
