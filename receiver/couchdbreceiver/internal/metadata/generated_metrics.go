// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for couchdbreceiver metrics.
type MetricsSettings struct {
	CouchdbAverageRequestTime MetricSettings `mapstructure:"couchdb.average_request_time"`
	CouchdbDatabaseOpen       MetricSettings `mapstructure:"couchdb.database.open"`
	CouchdbDatabaseOperations MetricSettings `mapstructure:"couchdb.database.operations"`
	CouchdbFileDescriptorOpen MetricSettings `mapstructure:"couchdb.file_descriptor.open"`
	CouchdbHttpdBulkRequests  MetricSettings `mapstructure:"couchdb.httpd.bulk_requests"`
	CouchdbHttpdRequests      MetricSettings `mapstructure:"couchdb.httpd.requests"`
	CouchdbHttpdResponses     MetricSettings `mapstructure:"couchdb.httpd.responses"`
	CouchdbHttpdViews         MetricSettings `mapstructure:"couchdb.httpd.views"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		CouchdbAverageRequestTime: MetricSettings{
			Enabled: true,
		},
		CouchdbDatabaseOpen: MetricSettings{
			Enabled: true,
		},
		CouchdbDatabaseOperations: MetricSettings{
			Enabled: true,
		},
		CouchdbFileDescriptorOpen: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdBulkRequests: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdRequests: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdResponses: MetricSettings{
			Enabled: true,
		},
		CouchdbHttpdViews: MetricSettings{
			Enabled: true,
		},
	}
}

// ResourceAttributeSettings provides common settings for a particular resource attribute.
type ResourceAttributeSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesSettings provides settings for couchdbreceiver resource attributes.
type ResourceAttributesSettings struct {
	CouchdbNodeName ResourceAttributeSettings `mapstructure:"couchdb.node.name"`
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{
		CouchdbNodeName: ResourceAttributeSettings{
			Enabled: true,
		},
	}
}

// AttributeHTTPMethod specifies the a value http.method attribute.
type AttributeHTTPMethod int

const (
	_ AttributeHTTPMethod = iota
	AttributeHTTPMethodCOPY
	AttributeHTTPMethodDELETE
	AttributeHTTPMethodGET
	AttributeHTTPMethodHEAD
	AttributeHTTPMethodOPTIONS
	AttributeHTTPMethodPOST
	AttributeHTTPMethodPUT
)

// String returns the string representation of the AttributeHTTPMethod.
func (av AttributeHTTPMethod) String() string {
	switch av {
	case AttributeHTTPMethodCOPY:
		return "COPY"
	case AttributeHTTPMethodDELETE:
		return "DELETE"
	case AttributeHTTPMethodGET:
		return "GET"
	case AttributeHTTPMethodHEAD:
		return "HEAD"
	case AttributeHTTPMethodOPTIONS:
		return "OPTIONS"
	case AttributeHTTPMethodPOST:
		return "POST"
	case AttributeHTTPMethodPUT:
		return "PUT"
	}
	return ""
}

// MapAttributeHTTPMethod is a helper map of string to AttributeHTTPMethod attribute value.
var MapAttributeHTTPMethod = map[string]AttributeHTTPMethod{
	"COPY":    AttributeHTTPMethodCOPY,
	"DELETE":  AttributeHTTPMethodDELETE,
	"GET":     AttributeHTTPMethodGET,
	"HEAD":    AttributeHTTPMethodHEAD,
	"OPTIONS": AttributeHTTPMethodOPTIONS,
	"POST":    AttributeHTTPMethodPOST,
	"PUT":     AttributeHTTPMethodPUT,
}

// AttributeOperation specifies the a value operation attribute.
type AttributeOperation int

const (
	_ AttributeOperation = iota
	AttributeOperationWrites
	AttributeOperationReads
)

// String returns the string representation of the AttributeOperation.
func (av AttributeOperation) String() string {
	switch av {
	case AttributeOperationWrites:
		return "writes"
	case AttributeOperationReads:
		return "reads"
	}
	return ""
}

// MapAttributeOperation is a helper map of string to AttributeOperation attribute value.
var MapAttributeOperation = map[string]AttributeOperation{
	"writes": AttributeOperationWrites,
	"reads":  AttributeOperationReads,
}

// AttributeView specifies the a value view attribute.
type AttributeView int

const (
	_ AttributeView = iota
	AttributeViewTemporaryViewReads
	AttributeViewViewReads
)

// String returns the string representation of the AttributeView.
func (av AttributeView) String() string {
	switch av {
	case AttributeViewTemporaryViewReads:
		return "temporary_view_reads"
	case AttributeViewViewReads:
		return "view_reads"
	}
	return ""
}

// MapAttributeView is a helper map of string to AttributeView attribute value.
var MapAttributeView = map[string]AttributeView{
	"temporary_view_reads": AttributeViewTemporaryViewReads,
	"view_reads":           AttributeViewViewReads,
}

type metricCouchdbAverageRequestTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.average_request_time metric with initial data.
func (m *metricCouchdbAverageRequestTime) init() {
	m.data.SetName("couchdb.average_request_time")
	m.data.SetDescription("The average duration of a served request.")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
}

func (m *metricCouchdbAverageRequestTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbAverageRequestTime) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbAverageRequestTime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbAverageRequestTime(settings MetricSettings) metricCouchdbAverageRequestTime {
	m := metricCouchdbAverageRequestTime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbDatabaseOpen struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.database.open metric with initial data.
func (m *metricCouchdbDatabaseOpen) init() {
	m.data.SetName("couchdb.database.open")
	m.data.SetDescription("The number of open databases.")
	m.data.SetUnit("{databases}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricCouchdbDatabaseOpen) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbDatabaseOpen) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbDatabaseOpen) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbDatabaseOpen(settings MetricSettings) metricCouchdbDatabaseOpen {
	m := metricCouchdbDatabaseOpen{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbDatabaseOperations struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.database.operations metric with initial data.
func (m *metricCouchdbDatabaseOperations) init() {
	m.data.SetName("couchdb.database.operations")
	m.data.SetDescription("The number of database operations.")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbDatabaseOperations) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("operation", operationAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbDatabaseOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbDatabaseOperations) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbDatabaseOperations(settings MetricSettings) metricCouchdbDatabaseOperations {
	m := metricCouchdbDatabaseOperations{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbFileDescriptorOpen struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.file_descriptor.open metric with initial data.
func (m *metricCouchdbFileDescriptorOpen) init() {
	m.data.SetName("couchdb.file_descriptor.open")
	m.data.SetDescription("The number of open file descriptors.")
	m.data.SetUnit("{files}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricCouchdbFileDescriptorOpen) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbFileDescriptorOpen) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbFileDescriptorOpen) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbFileDescriptorOpen(settings MetricSettings) metricCouchdbFileDescriptorOpen {
	m := metricCouchdbFileDescriptorOpen{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdBulkRequests struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.bulk_requests metric with initial data.
func (m *metricCouchdbHttpdBulkRequests) init() {
	m.data.SetName("couchdb.httpd.bulk_requests")
	m.data.SetDescription("The number of bulk requests.")
	m.data.SetUnit("{requests}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricCouchdbHttpdBulkRequests) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdBulkRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdBulkRequests) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdBulkRequests(settings MetricSettings) metricCouchdbHttpdBulkRequests {
	m := metricCouchdbHttpdBulkRequests{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdRequests struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.requests metric with initial data.
func (m *metricCouchdbHttpdRequests) init() {
	m.data.SetName("couchdb.httpd.requests")
	m.data.SetDescription("The number of HTTP requests by method.")
	m.data.SetUnit("{requests}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbHttpdRequests) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, httpMethodAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("http.method", httpMethodAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdRequests) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdRequests(settings MetricSettings) metricCouchdbHttpdRequests {
	m := metricCouchdbHttpdRequests{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdResponses struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.responses metric with initial data.
func (m *metricCouchdbHttpdResponses) init() {
	m.data.SetName("couchdb.httpd.responses")
	m.data.SetDescription("The number of each HTTP status code.")
	m.data.SetUnit("{responses}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbHttpdResponses) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, httpStatusCodeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("http.status_code", httpStatusCodeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdResponses) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdResponses) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdResponses(settings MetricSettings) metricCouchdbHttpdResponses {
	m := metricCouchdbHttpdResponses{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricCouchdbHttpdViews struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills couchdb.httpd.views metric with initial data.
func (m *metricCouchdbHttpdViews) init() {
	m.data.SetName("couchdb.httpd.views")
	m.data.SetDescription("The number of views read.")
	m.data.SetUnit("{views}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricCouchdbHttpdViews) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, viewAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("view", viewAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricCouchdbHttpdViews) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricCouchdbHttpdViews) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricCouchdbHttpdViews(settings MetricSettings) metricCouchdbHttpdViews {
	m := metricCouchdbHttpdViews{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilderConfig is a structural subset of an otherwise 1-1 copy of metadata.yaml
type MetricsBuilderConfig struct {
	Metrics            MetricsSettings            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesSettings `mapstructure:"resource_attributes"`
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                       pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                 int                 // maximum observed number of metrics per resource.
	resourceCapacity                int                 // maximum observed number of resource attributes.
	metricsBuffer                   pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                       component.BuildInfo // contains version information
	resourceAttributesSettings      ResourceAttributesSettings
	metricCouchdbAverageRequestTime metricCouchdbAverageRequestTime
	metricCouchdbDatabaseOpen       metricCouchdbDatabaseOpen
	metricCouchdbDatabaseOperations metricCouchdbDatabaseOperations
	metricCouchdbFileDescriptorOpen metricCouchdbFileDescriptorOpen
	metricCouchdbHttpdBulkRequests  metricCouchdbHttpdBulkRequests
	metricCouchdbHttpdRequests      metricCouchdbHttpdRequests
	metricCouchdbHttpdResponses     metricCouchdbHttpdResponses
	metricCouchdbHttpdViews         metricCouchdbHttpdViews
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsSettings(),
		ResourceAttributes: DefaultResourceAttributesSettings(),
	}
}

func NewMetricsBuilderConfig(ms MetricsSettings, ras ResourceAttributesSettings) MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            ms,
		ResourceAttributes: ras,
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                       pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                   pmetric.NewMetrics(),
		buildInfo:                       settings.BuildInfo,
		resourceAttributesSettings:      mbc.ResourceAttributes,
		metricCouchdbAverageRequestTime: newMetricCouchdbAverageRequestTime(mbc.Metrics.CouchdbAverageRequestTime),
		metricCouchdbDatabaseOpen:       newMetricCouchdbDatabaseOpen(mbc.Metrics.CouchdbDatabaseOpen),
		metricCouchdbDatabaseOperations: newMetricCouchdbDatabaseOperations(mbc.Metrics.CouchdbDatabaseOperations),
		metricCouchdbFileDescriptorOpen: newMetricCouchdbFileDescriptorOpen(mbc.Metrics.CouchdbFileDescriptorOpen),
		metricCouchdbHttpdBulkRequests:  newMetricCouchdbHttpdBulkRequests(mbc.Metrics.CouchdbHttpdBulkRequests),
		metricCouchdbHttpdRequests:      newMetricCouchdbHttpdRequests(mbc.Metrics.CouchdbHttpdRequests),
		metricCouchdbHttpdResponses:     newMetricCouchdbHttpdResponses(mbc.Metrics.CouchdbHttpdResponses),
		metricCouchdbHttpdViews:         newMetricCouchdbHttpdViews(mbc.Metrics.CouchdbHttpdViews),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(ResourceAttributesSettings, pmetric.ResourceMetrics)

// WithCouchdbNodeName sets provided value as "couchdb.node.name" attribute for current resource.
func WithCouchdbNodeName(val string) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		if ras.CouchdbNodeName.Enabled {
			rm.Resource().Attributes().PutStr("couchdb.node.name", val)
		}
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/couchdbreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricCouchdbAverageRequestTime.emit(ils.Metrics())
	mb.metricCouchdbDatabaseOpen.emit(ils.Metrics())
	mb.metricCouchdbDatabaseOperations.emit(ils.Metrics())
	mb.metricCouchdbFileDescriptorOpen.emit(ils.Metrics())
	mb.metricCouchdbHttpdBulkRequests.emit(ils.Metrics())
	mb.metricCouchdbHttpdRequests.emit(ils.Metrics())
	mb.metricCouchdbHttpdResponses.emit(ils.Metrics())
	mb.metricCouchdbHttpdViews.emit(ils.Metrics())

	for _, op := range rmo {
		op(mb.resourceAttributesSettings, rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordCouchdbAverageRequestTimeDataPoint adds a data point to couchdb.average_request_time metric.
func (mb *MetricsBuilder) RecordCouchdbAverageRequestTimeDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricCouchdbAverageRequestTime.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbDatabaseOpenDataPoint adds a data point to couchdb.database.open metric.
func (mb *MetricsBuilder) RecordCouchdbDatabaseOpenDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricCouchdbDatabaseOpen.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbDatabaseOperationsDataPoint adds a data point to couchdb.database.operations metric.
func (mb *MetricsBuilder) RecordCouchdbDatabaseOperationsDataPoint(ts pcommon.Timestamp, val int64, operationAttributeValue AttributeOperation) {
	mb.metricCouchdbDatabaseOperations.recordDataPoint(mb.startTime, ts, val, operationAttributeValue.String())
}

// RecordCouchdbFileDescriptorOpenDataPoint adds a data point to couchdb.file_descriptor.open metric.
func (mb *MetricsBuilder) RecordCouchdbFileDescriptorOpenDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricCouchdbFileDescriptorOpen.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbHttpdBulkRequestsDataPoint adds a data point to couchdb.httpd.bulk_requests metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdBulkRequestsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricCouchdbHttpdBulkRequests.recordDataPoint(mb.startTime, ts, val)
}

// RecordCouchdbHttpdRequestsDataPoint adds a data point to couchdb.httpd.requests metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdRequestsDataPoint(ts pcommon.Timestamp, val int64, httpMethodAttributeValue AttributeHTTPMethod) {
	mb.metricCouchdbHttpdRequests.recordDataPoint(mb.startTime, ts, val, httpMethodAttributeValue.String())
}

// RecordCouchdbHttpdResponsesDataPoint adds a data point to couchdb.httpd.responses metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdResponsesDataPoint(ts pcommon.Timestamp, val int64, httpStatusCodeAttributeValue string) {
	mb.metricCouchdbHttpdResponses.recordDataPoint(mb.startTime, ts, val, httpStatusCodeAttributeValue)
}

// RecordCouchdbHttpdViewsDataPoint adds a data point to couchdb.httpd.views metric.
func (mb *MetricsBuilder) RecordCouchdbHttpdViewsDataPoint(ts pcommon.Timestamp, val int64, viewAttributeValue AttributeView) {
	mb.metricCouchdbHttpdViews.recordDataPoint(mb.startTime, ts, val, viewAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
