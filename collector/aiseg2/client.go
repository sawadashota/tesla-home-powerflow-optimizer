package aiseg2

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector"

	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector/aiseg2/digest"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type Client struct {
	config     *Config
	httpclient *http.Client
}

var _ collector.Collector = new(Client)

type ClientOption func(*Client)

func WithHTTPClient(httpclient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpclient = httpclient
	}
}

func NewClient(opts ...ClientOption) (*Client, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}
	c := &Client{
		config:     cfg,
		httpclient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) getPage(ctx context.Context, path string) (*html.Node, error) {
	r := digest.New(c.config.User, c.config.Password, "MD5")
	resp, err := r.Do(ctx, http.MethodGet, c.config.url(path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return htmlquery.Parse(resp.Body)
}

func (c *Client) GetSurplusPower(ctx context.Context) (*model.PowerMetric, error) {
	const path = "/page/electricflow/111"
	doc, err := c.getPage(ctx, path)
	if err != nil {
		return nil, err
	}

	generationNode := htmlquery.FindOne(doc, "//div[@id='g_capacity']")
	usageNode := htmlquery.FindOne(doc, "//div[@id='u_capacity']")
	var (
		generation float64
		usage      float64
	)
	_, _ = fmt.Sscanf(htmlquery.InnerText(generationNode), "%fKW", &generation)
	_, _ = fmt.Sscanf(htmlquery.InnerText(usageNode), "%fKW", &usage)

	return &model.PowerMetric{
		SurplusWatt: int((generation - usage) * 1000),
		Timestamp:   time.Now(),
	}, nil
}
