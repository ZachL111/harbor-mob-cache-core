package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 52, Capacity: 93, Latency: 16, Risk: 6, Weight: 13}, wantScore: 156, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 86, Capacity: 87, Latency: 10, Risk: 14, Weight: 12}, wantScore: 181, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 107, Capacity: 76, Latency: 13, Risk: 24, Weight: 5}, wantScore: 95, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
