package apparel

import (
	"context"
	"fmt"
	"os"

	"github.com/imroc/req/v3"
	"go.uber.org/ratelimit"
	"golang.org/x/sync/semaphore"
)


type Client struct {
	*req.Client
	ratelimiter ratelimit.Limiter
	gate *semaphore.Weighted

	Persons *PersonsService
	RewardAccount *RewardAccountService
	Stores *StoresService
}

type service struct {
	client *Client
}

type CustomHttp struct {
	*req.Client
	rateLimiter ratelimit.Limiter
}

type ClientConfig struct {
	User string
	Password string
	Name string
	Host string
	Port string
	CountryCode string
}

func NewConfigFromEnv() (*ClientConfig, error) {
	user, exists := os.LookupEnv("AP21_APIUSER")
	if !exists {
		return nil, fmt.Errorf("AP21_APIUSER Not Set")
	}

	password, exists := os.LookupEnv("AP21_APIPASS")
	if !exists {
		return nil, fmt.Errorf("AP21_APIPASS Not Set")
	}

	host, exists := os.LookupEnv("AP21_APIHOST")
	if !exists {
		return nil, fmt.Errorf("AP21_APIHOST Not Set")
	}

	name, exists := os.LookupEnv("AP21_APINAME")
	if !exists {
		return nil, fmt.Errorf("AP21_APINAME Not Set")
	}

	port, exists := os.LookupEnv("AP21_APIPORT")
	if !exists {
		return nil, fmt.Errorf("AP21_APIPORT Not Set")
	}

	countryCode, exists := os.LookupEnv("AP21_APICOUNTRYCODE")
	if !exists {
		return nil, fmt.Errorf("AP21_APICOUNTRYCODE Not Set")
	}

	return &ClientConfig{
		User: user,
		Password: password,
		Host: host,
		Name: name,
		Port: port,
		CountryCode: countryCode,
	}, nil
}

func withThrottler(c *Client) *Client{
    c.ratelimiter = ratelimit.New(100)
    c.WrapRoundTripFunc(func (rt req.RoundTripper) req.RoundTripFunc {
        return func(req *req.Request) (resp *req.Response, err error) {
            c.ratelimiter.Take()
		    return rt.RoundTrip(req)
	    }
    })
    return c
}

func withMaxConcurrent(c *Client) *Client {
    c.gate = semaphore.NewWeighted(20)
    c.WrapRoundTripFunc(func (rt req.RoundTripper) req.RoundTripFunc {
        return func(req *req.Request) (resp *req.Response, err error) {
            ctx := context.Background()
            c.gate.Acquire(ctx, 1)
            defer c.gate.Release(1)
		    return rt.RoundTrip(req)
	    }
    })
    return c
}

func WithCustomData(c *Client) *Client {
	c.SetCommonQueryParam("CustomData", "true")
	return c
}

func NewClient(cC *ClientConfig) *Client {
	client := withMaxConcurrent(withThrottler(&Client{ req.C().
		SetCommonBasicAuth(cC.User, cC.Password).
		SetBaseURL(fmt.Sprintf("https://%s:%s/%s", cC.Host , cC.Port, cC.Name)).
		SetCommonHeader("Content-Type", "text/xml").
		SetCommonHeader("Accept", "version_4.0").
		SetCommonQueryParam("CountryCode", cC.CountryCode),
		nil,
		nil,
		nil,
		nil, 
		nil,
	}))

	client.Persons = &PersonsService{client}
	client.RewardAccount = &RewardAccountService{client}
	client.Stores = &StoresService{client}

	return client
}



