package aiseg2

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/morikuni/failure/v2"
	"golang.org/x/net/html"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector/aiseg2/digest"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type Client struct {
	config     *Config
	httpclient *http.Client

	evElectricityFlowName string
}

var _ collector.Collector = new(Client)

type ClientOption func(*Client)

func WithHTTPClient(httpclient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpclient = httpclient
	}
}

func WitEVElectricityFlowName(evElectricityFlowName string) ClientOption {
	return func(c *Client) {
		c.evElectricityFlowName = evElectricityFlowName
	}
}

func NewClient(opts ...ClientOption) (*Client, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}
	c := &Client{
		config:                cfg,
		httpclient:            http.DefaultClient,
		evElectricityFlowName: "電気自動車１",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) getPage(ctx context.Context, path string, query url.Values) (*html.Node, error) {
	r := digest.New(c.config.User, c.config.Password, "MD5")
	resp, err := r.Do(ctx, http.MethodGet, c.config.url(path, query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return htmlquery.Parse(resp.Body)
}

func (c *Client) GetSurplusPower(ctx context.Context) (*model.PowerMetric, error) {
	const path = "/page/electricflow/111"
	doc, err := c.getPage(ctx, path, nil)
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
		Name:      "surplus",
		Watt:      int((generation - usage) * 1000),
		Timestamp: time.Now(),
	}, nil
}

func (c *Client) GetEVUsagePower(ctx context.Context) (*model.PowerMetric, error) {
	metrics, err := c.listDetailUsagePower(ctx)
	if err != nil {
		return nil, err
	}
	for _, metric := range metrics {
		if metric.Name == c.evElectricityFlowName {
			return metric, nil
		}
	}
	return nil, failure.New(model.ErrCodeMetricNotFound, failure.Message("metric not found"))
}

func (c *Client) listDetailUsagePower(ctx context.Context) (model.PowerMetricList, error) {
	const path = "/page/electricflow/1113"
	const maxPage = 20
	var lastPageNames string
	metrics := make(model.PowerMetricList, 0, 30)

	for page := 1; page <= maxPage; page++ {
		doc, err := c.getPage(ctx, path, url.Values{"id": {fmt.Sprint(page)}})
		if err != nil {
			return nil, err
		}

		// check duplication
		// if the names are same as the last page, it means end of the list
		names := make([]string, 0, 10)
		for i := 1; i <= 10; i++ {
			nameNode := htmlquery.FindOne(doc, fmt.Sprintf("//div[@id='stage_%d']/div[@class='c_device']", i))
			if nameNode == nil {
				break
			}
			name := htmlquery.InnerText(nameNode)
			if len(name) == 0 {
				break
			}
			names = append(names, name)
		}
		if strings.Join(names, ",") == lastPageNames {
			break
		}

		for i := 1; i <= 10; i++ {
			nameNode := htmlquery.FindOne(doc, fmt.Sprintf("//div[@id='stage_%d']/div[@class='c_device']", i))
			if nameNode == nil {
				break
			}
			name := htmlquery.InnerText(nameNode)
			if len(name) == 0 {
				break
			}
			valueNode := htmlquery.FindOne(doc, fmt.Sprintf("//div[@id='stage_%d']/div[@class='c_value']", i))
			if valueNode == nil {
				break
			}
			var value int
			_, _ = fmt.Sscanf(htmlquery.InnerText(valueNode), "%dW", &value)
			metrics = append(metrics, &model.PowerMetric{
				Name:      name,
				Watt:      value,
				Timestamp: time.Now(),
			})
		}

		lastPageNames = strings.Join(names, ",")
	}
	return metrics, nil
}
