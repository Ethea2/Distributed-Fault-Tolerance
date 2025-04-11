package proxy

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"path"
)

type ServiceProxy struct {
	baseURL *url.URL
	client  *http.Client
}

func NewServiceProxy(serviceUrl string) *ServiceProxy {
	parsedURL, err := url.Parse(serviceUrl)

	if err != nil {
		panic(err)
	}

	return &ServiceProxy{
		baseURL: parsedURL,
		client:  &http.Client{},
	}
}

func (p *ServiceProxy) ForwardRequest(r *http.Request, endpoint string) (*http.Response, error) {
	targetURL := *p.baseURL
	targetURL.Path = path.Join(targetURL.Path, endpoint)

	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
		r.Body.Close()

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	proxyReq, err := http.NewRequestWithContext(
		r.Context(),
		r.Method,
		targetURL.String(),
		bytes.NewBuffer(bodyBytes),
	)

	if err != nil {
		return nil, err
	}

	for key, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	proxyReq.URL.RawQuery = r.URL.RawQuery

	return p.client.Do(proxyReq)
}

func (p *ServiceProxy) ForwardRequestAndCopyResponse(w http.ResponseWriter, r *http.Request, endpoint string) error {
	resp, err := p.ForwardRequest(r, endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	return err
}
