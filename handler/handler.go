package handler

func HandleMetric(metric string) string {
	switch metric {
	case "cpu.load":
		return "Aktuelle CPU Auslastung"
	default:
		return "Invalid Metric!"

	}
}