package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"gitlab.tdnet.it/cochise/golang/template-project-ms/pkg/config"
	"golang.org/x/net/context"

	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

var restClient *RestClient

//RestClient : struttura per le chiamate http client con zipkin
type RestClient struct {
	BaseURL    *url.URL
	httpClient *zipkinhttp.Client
	config     *config.Config
}

// RestClientFactory factory rest client
func RestClientFactory(url *url.URL, httpClient *zipkinhttp.Client, config *config.Config) *RestClient {
	return &RestClient{
		BaseURL:    url,
		httpClient: httpClient,
		config:     config,
	}
}

//NewRequest : crea una nuova richiesta http
func (rs *RestClient) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, _ := url.Parse(path)
	u := rs.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Do esegue una chiamata rest
func (rs *RestClient) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	log.Printf("------ REST CLIENT START %#+v %#+v\n ---- ", req.Method, req.URL.String())
	//log.Printf("request:  %#+v %#+v\n", req.Method, req.URL.String())
	//log.Printf("header: %#+v\n", req.Header)
	span := zipkin.SpanFromContext(ctx)
	newCtx := zipkin.NewContext(req.Context(), span)
	newRequest := req.WithContext(newCtx)
	span.Annotate(time.Now(), "START REQUEST "+req.URL.String())
	resp, err := rs.httpClient.DoWithAppSpan(newRequest, req.URL.String())
	if err != nil {
		log.Printf("Errore in client do: %#+v\n", err)
		span.Annotate(time.Now(), "END REQUEST "+req.URL.String())
		return nil, err
	}
	span.Annotate(time.Now(), "END REQUEST "+req.URL.String())
	log.Printf("Response Code : %#+v\n", resp.Status)
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		bodyB, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(bytes.Replace(bodyB, []byte("\r"), []byte("\r\n"), -1))
		log.Printf("Response err  : %#+v\n", bodyStr)
		log.Println("------ REST CLIENT END ---- ", req.URL.String())
		return resp, errors.New(bodyStr)
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		log.Printf("json decode err  : %#+v\n", err)
	} else {
		//log.Printf("Response Body : %#+v\n", v)
	}
	log.Println("------ REST CLIENT END ---- ", req.URL.String())
	return resp, err
}

//GetRestClient : metodo per il recupero della struttura RestClient
func GetRestClient(uri string) *RestClient {
	zipkin := GetZipkinInstance()
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	zipkinClient, err := zipkinhttp.NewClient(zipkin.Tracer, zipkinhttp.WithClient(netClient))
	if err != nil {
		log.Printf("unable to create client: %+v\n", err)
	}
	url, err := url.Parse(uri)
	return RestClientFactory(url, zipkinClient, config.GetInstance())
}
