package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 52, Capacity: 93, Latency: 16, Risk: 6, Weight: 13}
	if got := Score(signal); got != 156 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 86, Capacity: 87, Latency: 10, Risk: 14, Weight: 12}
	if got := Score(signal); got != 181 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 107, Capacity: 76, Latency: 13, Risk: 24, Weight: 5}
	if got := Score(signal); got != 95 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
