package main

import (
	"errors"
	"flag"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/cloudwatch"
	mp "github.com/mackerelio/go-mackerel-plugin"
	"log"
	"os"
	"time"
)

const (
	Namespace          = "AWS/CloudFront"
	Region             = "us-east-1"
	MetricsTypeAverage = "Average"
	MetricsTypeSum     = "Sum"
)

var graphdef map[string](mp.Graphs) = map[string](mp.Graphs){
	"cloudfront.Requests": mp.Graphs{
		Label: "CloudFront Requests",
		Unit:  "integer",
		Metrics: [](mp.Metrics){
			mp.Metrics{Name: "Requests", Label: "Requests"},
		},
	},
	"cloudfront.Transfer": mp.Graphs{
		Label: "CloudFront Transfer",
		Unit:  "bytes",
		Metrics: [](mp.Metrics){
			mp.Metrics{Name: "BytesDownloaded", Label: "Download", Stacked: true},
			mp.Metrics{Name: "BytesUploaded", Label: "Upload", Stacked: true},
		},
	},
	"cloudfront.ErrorRate": mp.Graphs{
		Label: "CloudFront ErrorRate",
		Unit:  "percentage",
		Metrics: [](mp.Metrics){
			mp.Metrics{Name: "4xxErrorRate", Label: "4xx", Stacked: true},
			mp.Metrics{Name: "5xxErrorRate", Label: "5xx", Stacked: true},
		},
	},
}

type Metrics struct {
	Name string
	Type string
}

type CloudFrontPlugin struct {
	AccessKeyId     string
	SecretAccessKey string
	CloudWatch      *cloudwatch.CloudWatch
	Name            string
}

func (p *CloudFrontPlugin) Prepare() error {
	auth, err := aws.GetAuth(p.AccessKeyId, p.SecretAccessKey, "", time.Now())
	if err != nil {
		return err
	}

	p.CloudWatch, err = cloudwatch.NewCloudWatch(auth, aws.Regions[Region].CloudWatchServicepoint)
	if err != nil {
		return err
	}

	return nil
}

func (p CloudFrontPlugin) GetLastPoint(metric Metrics) (float64, error) {
	now := time.Now()

	dimensions := []cloudwatch.Dimension{
		cloudwatch.Dimension{
			Name:  "DistributionId",
			Value: p.Name,
		},
		cloudwatch.Dimension{
			Name:  "Region",
			Value: "Global",
		},
	}

	response, err := p.CloudWatch.GetMetricStatistics(&cloudwatch.GetMetricStatisticsRequest{
		Dimensions: dimensions,
		StartTime:  now.Add(time.Duration(180) * time.Second * -1), // 3 min (to fetch at least 1 data-point)
		EndTime:    now,
		MetricName: metric.Name,
		Period:     60,
		Statistics: []string{metric.Type},
		Namespace:  Namespace,
	})
	if err != nil {
		return 0, err
	}

	datapoints := response.GetMetricStatisticsResult.Datapoints
	if len(datapoints) == 0 {
		return 0, errors.New("fetched no datapoints")
	}

	// get a least recently datapoint
	// because a most recently datapoint is not stable.
	least := now
	var latestVal float64
	for _, dp := range datapoints {
		if dp.Timestamp.Before(least) {
			least = dp.Timestamp
			if metric.Type == MetricsTypeAverage {
				latestVal = dp.Average
			} else if metric.Type == MetricsTypeSum {
				latestVal = dp.Sum
			}
		}
	}

	return latestVal, nil
}

func (p CloudFrontPlugin) FetchMetrics() (map[string]float64, error) {
	stat := make(map[string]float64)

	for _, met := range [...]Metrics{
		{Name: "Requests", Type: MetricsTypeSum},
		{Name: "BytesDownloaded", Type: MetricsTypeSum},
		{Name: "BytesUploaded", Type: MetricsTypeSum},
		{Name: "4xxErrorRate", Type: MetricsTypeAverage},
		{Name: "5xxErrorRate", Type: MetricsTypeAverage},
	} {
		v, err := p.GetLastPoint(met)
		if err == nil {
			stat[met.Name] = v
		} else {
			log.Printf("%s: %s", met, err)
		}
	}

	return stat, nil
}

func (p CloudFrontPlugin) GraphDefinition() map[string](mp.Graphs) {
	return graphdef
}

func main() {
	optAccessKeyId := flag.String("access-key-id", "", "AWS Access Key ID")
	optSecretAccessKey := flag.String("secret-access-key", "", "AWS Secret Access Key")
	optIdentifier := flag.String("identifier", "", "Distribution ID")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	var plugin CloudFrontPlugin

	plugin.AccessKeyId = *optAccessKeyId
	plugin.SecretAccessKey = *optSecretAccessKey
	plugin.Name = *optIdentifier

	err := plugin.Prepare()
	if err != nil {
		log.Fatalln(err)
	}

	helper := mp.NewMackerelPlugin(plugin)
	if *optTempfile != "" {
		helper.Tempfile = *optTempfile
	} else {
		helper.Tempfile = "/tmp/mackerel-plugin-cloudfront"
	}

	if os.Getenv("MACKEREL_AGENT_PLUGIN_META") != "" {
		helper.OutputDefinitions()
	} else {
		helper.OutputValues()
	}
}
