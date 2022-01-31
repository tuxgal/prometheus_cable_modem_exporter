package main

import "github.com/prometheus/client_golang/prometheus"

const (
	cmInstanceLabel = "cable_modem_instance"
)

var (
	invalid = prometheus.NewDesc(
		"cable_modem_error",
		"Error collecting metrics from cable modem",
		nil,
		nil,
	)

	// Up metric to indicate whether the cable modem is down or up.
	up        = makeDesc("up", "Cable Modem Up")
	descModel = makeDesc(
		"cable_modem_info_model",
		"Cable Modem Model",
		"model",
	)
	descSerialNumber = makeDesc(
		"cable_modem_info_serial_number",
		"Cable Modem Serial Number",
		"serial_number",
	)
	descMACAddress = makeDesc(
		"cable_modem_info_mac_address",
		"Cable Modem MAC Address",
		"mac_address",
	)
	descFrontPanelLightsOn = makeDesc(
		"cable_modem_settings_front_panel_lights_on",
		"Cable Modem Settings Front Panel Lights On",
	)
	descEnergyEffEthOn = makeDesc(
		"cable_modem_settings_energy_efficient_ethernet_on",
		"Cable Modem Settings Energy Efficient Ethernet On",
	)
	descAskMeLaterOn = makeDesc(
		"cable_modem_settings_ask_me_later_on",
		"Cable Modem Settings Ask Me Later On",
	)
	descNeverAskOn = makeDesc(
		"cable_modem_settings_never_ask_on",
		"Cable Modem Settings Never Ask On",
	)
	descCertInstalled = makeDesc(
		"cable_modem_software_certificate_installed",
		"Cable Modem Software Certificate Installed",
	)
	descFwVer = makeDesc(
		"cable_modem_software_firmware_version",
		"Cable Modem Software Firmware Version",
		"firmware_version",
	)
	descCustomerVer = makeDesc(
		"cable_modem_software_customer_version",
		"Cable Modem Software Customer Version",
		"customer_version",
	)
	descHDVerVer = makeDesc(
		"cable_modem_software_hd_version",
		"Cable Modem Software HD Version",
		"hd_version",
	)
	descDOCSISVer = makeDesc(
		"cable_modem_software_docsis_version",
		"Cable Modem Software DOCSIS Version",
		"docsis_version",
	)
	descDsPower = makeDesc(
		"cable_modem_connection_downstream_signal_power_dbmv",
		"Cable Modem Downstream Signal Power in dB mV",
	)
	descDsSNR = makeDesc(
		"cable_modem_connection_downstream_signal_snr_db",
		"Cable Modem Downstream Signal SNR in dB",
	)
	allMetrics = []*prometheus.Desc{
		up,
		descModel,
		descSerialNumber,
		descMACAddress,
		descFrontPanelLightsOn,
		descEnergyEffEthOn,
		descAskMeLaterOn,
		descNeverAskOn,
		descCertInstalled,
		descFwVer,
		descCustomerVer,
		descHDVerVer,
		descDOCSISVer,
		descDsPower,
		descDsSNR,
	}
)

func makeDesc(metric string, desc string, labels ...string) *prometheus.Desc {
	labels = append([]string{cmInstanceLabel}, labels...)
	return prometheus.NewDesc(metric, desc, labels, nil)
}

type metricsHelper struct {
	host string
	ch   chan<- prometheus.Metric
}

func newMetricsHelper(host string, ch chan<- prometheus.Metric) *metricsHelper {
	return &metricsHelper{
		host: host,
		ch:   ch,
	}
}

func (m *metricsHelper) raiseError(err error) {
	m.ch <- prometheus.NewInvalidMetric(invalid, err)
}

func (m *metricsHelper) setStr(desc *prometheus.Desc, labelValue ...string) {
	m.setGauge(desc, 1, labelValue...)
}

func (m *metricsHelper) setInt32(desc *prometheus.Desc, value int32) {
	m.setGauge(desc, float64(value))
}

func (m *metricsHelper) setBool(desc *prometheus.Desc, state bool) {
	var value float64
	if state {
		value = 1
	}
	m.setGauge(desc, value)
}

func (m *metricsHelper) setGauge(desc *prometheus.Desc, value float64, labelValues ...string) {
	labelValues = append([]string{m.host}, labelValues...)
	m.ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		value,
		labelValues...,
	)
}
