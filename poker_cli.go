package pkr

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/litencatt/pkr/service"
)

const Name string = "pkr"

var Version = "dev"

type PokerCLI struct {
	DebugMode bool
	service   service.PokerService
}

func NewPokerCLI() *PokerCLI {
	return &PokerCLI{
		service: service.NewPokerService(service.PokerServiceConfig{
			DebugMode: true,
		}),
	}
}

func (cli *PokerCLI) Run() error {
	sleepSec := 1
	fmt.Println("*********************")
	fmt.Println("* Welcome to Poker! *")
	fmt.Println("*********************")
	fmt.Println()
	time.Sleep(time.Duration(sleepSec) * time.Second)

	for {
		if cli.service.IsStartRound() {
			rounds := cli.service.GetRounds()
			fmt.Printf("Round %d start\n", rounds)
			time.Sleep(time.Duration(sleepSec) * time.Second)

			ante := cli.service.GetCurrentAnteAmount()
			blind := cli.service.GetCurrentBlindMulti()
			fmt.Printf("Ante:%d, Blind:%v\n", ante, blind)
			time.Sleep(time.Duration(sleepSec) * time.Second)
			cli.service.StartRound()
		}

		roundStats := cli.service.GetRoundStats()
		fmt.Printf("Score at least: %d\n", roundStats.ScoreAtLeast)
		fmt.Printf("Round score: %d\n", roundStats.TotalScore)
		fmt.Printf("Hands: %d, Discards: %d\n", roundStats.Hands, roundStats.Discards)
		fmt.Println()

		// Draw cards
		drawNum := cli.service.GetNextDrawNum()
		cards, err := cli.service.DrawCard(drawNum)
		if err != nil {
			return err
		}
		fmt.Println("[Draw", drawNum, "cards]")
		if cli.DebugMode {
			for _, card := range cards {
				fmt.Println(card.String())
			}
			fmt.Println()
		}

		// Select cards
		var selectCards []string
		for {
			selectCards = nil
			promptMs := &survey.MultiSelect{
				Message: "Select cards",
				Options: cli.service.GetHandCardString(),
			}
			err := survey.AskOne(promptMs, &selectCards, survey.WithPageSize(8))
			if err == terminal.InterruptErr {
				fmt.Println("interrupted")
				os.Exit(0)
			}

			selectCardNum := len(selectCards)
			if selectCardNum <= 5 {
				break
			}
			fmt.Println("Please select less than 5 cards")
			fmt.Println()
		}
		fmt.Print("[Selected cards]\n")
		for _, card := range selectCards {
			fmt.Println(card)
		}
		fmt.Println()

		// Play or Discard or Cancel
		var selectAction string
		actions := cli.service.GetEnableActions()
		prompt := &survey.Select{
			Message: "Select action:",
			Options: actions,
		}
		if err := survey.AskOne(prompt, &selectAction); err == terminal.InterruptErr {
			fmt.Println("interrupted")
			os.Exit(0)
		}

		cli.service.SelectCards(selectCards)
		cli.service.SetSelectAction(selectAction)
		if selectAction == "Discard" {
			if err := cli.service.DiscardHand(); err != nil {
				return err
			}
			continue
		}
		if selectAction == "Cancel" {
			cli.service.CancelHand()
			continue
		}
		if selectAction == "Play" {
			r, err := cli.service.PlayHand()
			if err != nil {
				return err
			}

			fmt.Printf("\nHand: %s", r.HandType)
			fmt.Printf("\nChip: %d, Mult: %d\n", r.Chip, r.Mult)
			fmt.Printf("\nScore: %d\n\n", r.Score)

			time.Sleep(time.Duration(sleepSec) * time.Second)
		}

		// show remain cards
		if cli.DebugMode {
			fmt.Print("[Remain cards]\n")
			remainCards := cli.service.GetRemainCardString()
			for _, card := range remainCards {
				fmt.Println(card)
			}
			fmt.Println()
		}

		roundResultStats := cli.service.GetRoundStats()
		if cli.service.IsRoundWin() {
			fmt.Printf("Score at least: %d, Round score: %d\n", roundResultStats.ScoreAtLeast, roundResultStats.TotalScore)
			fmt.Println("")

			prompt := &survey.Select{
				Message: "You win this round!",
				Options: []string{"Next"},
			}
			if err := survey.AskOne(prompt, &selectAction); err == terminal.InterruptErr {
				fmt.Println("interrupted")
				os.Exit(0)
			}
			cli.service.NextRound()
			continue
		}

		if roundResultStats.Hands == 0 {
			fmt.Printf("Score at least: %d, Round score: %d\n", roundResultStats.ScoreAtLeast, roundResultStats.TotalScore)
			fmt.Println("You lose!")
			break
		}

	}

	return nil
}
