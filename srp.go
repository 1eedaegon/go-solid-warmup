package main

// Case 1 Single responsibility
type FinancialReport struct {
	report string
}

func (f *FinancialReport) SendReport(email string) {}

type MarketingReport struct {
	report string
}

func (m *MarketingReport) SendReport(email string) {}
// 각 report는 report만 담당해야 하는데, report를 보내는 책임도 같이 포함되어있다.

// Report의 책임을 정의
type Report interface {
	Report() string
}

type FinancialReport struct {
	report string
}

func (r *FinancialReport) Report() {
	return r.report
}

type ReportSender struct {
	r Report
	target string
}

func (s *ReportSender) SendReport(report Report) {
	s.r = report
	// ...
}
