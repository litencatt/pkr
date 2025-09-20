package service

import (
	"testing"

	"github.com/litencatt/pkr/entity"
)

func TestNewPokerService(t *testing.T) {
	config := PokerServiceConfig{
		DebugMode: true,
	}
	service := NewPokerService(config)

	if service == nil {
		t.Fatal("NewPokerService() returned nil")
	}

	// Type assertion to access private fields for testing
	ps, ok := service.(*pokerService)
	if !ok {
		t.Fatal("Service is not of type *pokerService")
	}

	if ps.config.DebugMode != true {
		t.Error("DebugMode should be true")
	}

	if ps.runInfo == nil {
		t.Fatal("runInfo should not be nil")
	}

	if ps.round == nil {
		t.Fatal("round should not be nil")
	}
}

func TestGetNextDrawNum(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Initial state: should return default deal
	drawNum := service.GetNextDrawNum()
	if drawNum != ps.runInfo.DefaultDeal {
		t.Errorf("GetNextDrawNum() = %d, want %d", drawNum, ps.runInfo.DefaultDeal)
	}

	// After Cancel action
	ps.round.BeforeSelectAction = "Cancel"
	drawNum = service.GetNextDrawNum()
	if drawNum != 0 {
		t.Errorf("After Cancel, GetNextDrawNum() = %d, want 0", drawNum)
	}

	// After other action with selected cards
	ps.round.BeforeSelectAction = "Play"
	ps.round.SelectedCards = []entity.Trump{
		{Suit: entity.Spades, Rank: entity.Ace},
		{Suit: entity.Hearts, Rank: entity.King},
		{Suit: entity.Diamonds, Rank: entity.Queen},
	}
	drawNum = service.GetNextDrawNum()
	if drawNum != 3 {
		t.Errorf("With 3 selected cards, GetNextDrawNum() = %d, want 3", drawNum)
	}
}

func TestIsStartRound(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Initially should be true
	if !service.IsStartRound() {
		t.Error("IsStartRound() should be true initially")
	}

	// After unsetting start next
	ps.runInfo.UnsetStartNext()
	if service.IsStartRound() {
		t.Error("IsStartRound() should be false after UnsetStartNext()")
	}
}

func TestStartRound(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Start round
	err := service.StartRound()
	if err != nil {
		t.Fatalf("StartRound() returned error: %v", err)
	}

	// Check that StartNext is unset
	if service.IsStartRound() {
		t.Error("IsStartRound() should be false after StartRound()")
	}

	// Check that round is initialized with correct score
	scoreAtLeast := int(float64(service.GetCurrentAnteAmount()) * service.GetCurrentBlindMulti())
	stats := ps.round.GetRoundStats()
	if stats.ScoreAtLeast != scoreAtLeast {
		t.Errorf("ScoreAtLeast = %d, want %d", stats.ScoreAtLeast, scoreAtLeast)
	}
}

func TestGetRounds(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	rounds := service.GetRounds()
	if rounds != ps.runInfo.Rounds {
		t.Errorf("GetRounds() = %d, want %d", rounds, ps.runInfo.Rounds)
	}
}

func TestIsRoundWin(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Initially should not win
	if service.IsRoundWin() {
		t.Error("IsRoundWin() should be false initially")
	}

	// Set score to win
	ps.round.Stats.TotalScore = ps.round.Stats.ScoreAtLeast + 100
	if !service.IsRoundWin() {
		t.Error("IsRoundWin() should be true when score is sufficient")
	}
}

func TestNextRound(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	initialRounds := ps.runInfo.Rounds
	initialBlindIndex := ps.runInfo.BlindIndex

	err := service.NextRound()
	if err != nil {
		t.Fatalf("NextRound() returned error: %v", err)
	}

	// Check that rounds increased
	if ps.runInfo.Rounds != initialRounds+1 {
		t.Errorf("Rounds = %d, want %d", ps.runInfo.Rounds, initialRounds+1)
	}

	// Check that blind index increased (ante index only increases when blind index wraps)
	if ps.runInfo.BlindIndex != initialBlindIndex+1 {
		t.Errorf("BlindIndex = %d, want %d", ps.runInfo.BlindIndex, initialBlindIndex+1)
	}

	// Check that StartNext is set
	if !service.IsStartRound() {
		t.Error("IsStartRound() should be true after NextRound()")
	}
}

func TestGetRoundStats(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Set some stats
	ps.round.Stats.Hands = 3
	ps.round.Stats.Discards = 2
	ps.round.Stats.TotalScore = 150
	ps.round.Stats.ScoreAtLeast = 300

	stats := service.GetRoundStats()
	if stats == nil {
		t.Fatal("GetRoundStats() returned nil")
	}

	if stats.Hands != 3 {
		t.Errorf("Stats.Hands = %d, want 3", stats.Hands)
	}

	if stats.Discards != 2 {
		t.Errorf("Stats.Discards = %d, want 2", stats.Discards)
	}

	if stats.TotalScore != 150 {
		t.Errorf("Stats.TotalScore = %d, want 150", stats.TotalScore)
	}

	if stats.ScoreAtLeast != 300 {
		t.Errorf("Stats.ScoreAtLeast = %d, want 300", stats.ScoreAtLeast)
	}
}

func TestSelectCards(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Set up hand cards
	ps.round.HandCards = []entity.Trump{
		{Suit: entity.Spades, Rank: entity.Ace},
		{Suit: entity.Hearts, Rank: entity.King},
		{Suit: entity.Diamonds, Rank: entity.Queen},
		{Suit: entity.Clubs, Rank: entity.Jack},
		{Suit: entity.Spades, Rank: entity.Ten},
	}

	// Select cards
	selectedStrings := []string{"A of Spades", "K of Hearts", "T of Spades"}
	err := service.SelectCards(selectedStrings)

	if err != nil {
		t.Fatalf("SelectCards() returned error: %v", err)
	}

	if len(ps.round.SelectedCards) != 3 {
		t.Errorf("SelectedCards has %d cards, want 3", len(ps.round.SelectedCards))
	}
}

func TestDrawCard(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Shuffle deck first
	ps.round.Deck.Shuffle()

	// Draw cards
	cards, err := service.DrawCard(5)
	if err != nil {
		t.Fatalf("DrawCard() returned error: %v", err)
	}

	if len(cards) != 5 {
		t.Errorf("DrawCard(5) returned %d cards, want 5", len(cards))
	}

	if len(ps.round.HandCards) != 5 {
		t.Errorf("HandCards has %d cards, want 5", len(ps.round.HandCards))
	}
}

func TestGetCurrentAnteAmount(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	ante := service.GetCurrentAnteAmount()
	expectedAnte := entity.AnteAmounts[ps.runInfo.AnteIndex]

	if ante != expectedAnte {
		t.Errorf("GetCurrentAnteAmount() = %d, want %d", ante, expectedAnte)
	}
}

func TestGetCurrentBlindMulti(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	blind := service.GetCurrentBlindMulti()
	expectedBlind := entity.BlindMultis[ps.runInfo.BlindIndex]

	if blind != expectedBlind {
		t.Errorf("GetCurrentBlindMulti() = %f, want %f", blind, expectedBlind)
	}
}

func TestGetHandCardString(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Set up hand cards
	ps.round.HandCards = []entity.Trump{
		{Suit: entity.Spades, Rank: entity.Ace},
		{Suit: entity.Hearts, Rank: entity.King},
		{Suit: entity.Diamonds, Rank: entity.Queen},
	}

	cardStrings := service.GetHandCardString()

	if len(cardStrings) != 3 {
		t.Errorf("GetHandCardString() returned %d strings, want 3", len(cardStrings))
	}

	expectedStrings := []string{"A of Spades", "K of Hearts", "Q of Diamonds"}
	for i, expected := range expectedStrings {
		if cardStrings[i] != expected {
			t.Errorf("GetHandCardString()[%d] = %s, want %s", i, cardStrings[i], expected)
		}
	}
}

func TestGetRemainCardString(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Set up remain cards
	ps.round.RemainCards = []entity.Trump{
		{Suit: entity.Clubs, Rank: entity.Jack},
		{Suit: entity.Spades, Rank: entity.Ten},
	}

	cardStrings := service.GetRemainCardString()

	if len(cardStrings) != 2 {
		t.Errorf("GetRemainCardString() returned %d strings, want 2", len(cardStrings))
	}

	expectedStrings := []string{"J of Clubs", "T of Spades"}
	for i, expected := range expectedStrings {
		if cardStrings[i] != expected {
			t.Errorf("GetRemainCardString()[%d] = %s, want %s", i, cardStrings[i], expected)
		}
	}
}

func TestGetEnableActions(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Initial state with hands and discards
	ps.round.Stats.Hands = 3
	ps.round.Stats.Discards = 2
	actions := service.GetEnableActions()

	// Should have Play, Discard, Cancel
	expectedActions := []string{"Play", "Discard", "Cancel"}
	if len(actions) != len(expectedActions) {
		t.Errorf("GetEnableActions() returned %d actions, want %d", len(actions), len(expectedActions))
	}

	for i, expected := range expectedActions {
		if actions[i] != expected {
			t.Errorf("GetEnableActions()[%d] = %s, want %s", i, actions[i], expected)
		}
	}

	// No discards left
	ps.round.Stats.Discards = 0
	actions = service.GetEnableActions()

	// Should have Play, Cancel
	expectedActions = []string{"Play", "Cancel"}
	if len(actions) != len(expectedActions) {
		t.Errorf("With no discards, GetEnableActions() returned %d actions, want %d", len(actions), len(expectedActions))
	}

	for i, expected := range expectedActions {
		if actions[i] != expected {
			t.Errorf("GetEnableActions()[%d] = %s, want %s", i, actions[i], expected)
		}
	}
}

func TestSetAction(t *testing.T) {
	service := NewPokerService(PokerServiceConfig{})
	ps := service.(*pokerService)

	// Test setting different actions
	actions := []string{"Play", "Discard", "Cancel"}

	for _, action := range actions {
		service.SetAction(action)
		if ps.round.BeforeSelectAction != action {
			t.Errorf("After SetAction(%s), BeforeSelectAction = %s", action, ps.round.BeforeSelectAction)
		}
	}
}
