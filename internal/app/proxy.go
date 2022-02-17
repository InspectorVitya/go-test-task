package app

import (
	"github.com/inspectorvitya/go-test-task/internal/cache"
	"io"
	"net/http"
)

type Proxy struct {
	cache      cache.Cache
	backendURL string
}

func NewAppProxy(cache cache.Cache, backend string) *Proxy {
	return &Proxy{
		cache:      cache,
		backendURL: backend,
	}
}
func (p *Proxy) Get(url string) (interface{}, error) {
	data, exist := p.cache.Get(cache.Key(url))
	if exist {
		return data, nil
	}
	newData, err := p.requestOnBackend(url)
	if err != nil {
		return nil, err
	}
	p.cache.Set(cache.Key(url), newData)
	return newData, nil
}

func (p *Proxy) requestOnBackend(url string) (interface{}, error) {
	res, err := http.Get(p.backendURL + url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
