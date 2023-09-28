package interactions_test

import (
	"testing"

	"github.com/crazyvan/go-specs-greet/domain/interactions"
	"github.com/crazyvan/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}
