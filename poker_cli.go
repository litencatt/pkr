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
	ante := cli.service.GetCurrentAnte()
	blind := cli.service.GetCurrentBlind()
	ScoreAtLeast := int(float64(ante) * blind)

	fmt.Println("Welcome to Poker!")
	fmt.Println()
	fmt.Printf("Round start\n")
	fmt.Printf("Ante:%d, Blind:%v\n", ante, blind)
	fmt.Printf("Score at least: %d\n", ScoreAtLeast)
	cli.service.StartRound(ScoreAtLeast)

	var selectCardNum int
	var nextDrawNum int
	var selectActioin string

	for {
		roundStats := cli.service.GetRoundStats()
		fmt.Printf("Score at least: %d\n", roundStats.ScoreAtLeast)
		fmt.Printf("Round score: %d\n", roundStats.TotalScore)
		fmt.Printf("Hands: %d, Discards: %d\n", roundStats.Hands, roundStats.Discards)
		fmt.Println()

		// Draw cards
		if selectActioin == "Cancel" {
			nextDrawNum = 0
		} else if selectCardNum == 0 {
			nextDrawNum = cli.service.GetNextDrawNum()
		} else {
			// Draw cards num is same as the last selected cards num
			nextDrawNum = selectCardNum
		}

		if err := cli.service.DrawCard(nextDrawNum); err != nil {
			return err
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

			selectCardNum = len(selectCards)
			if selectCardNum <= 5 {
				break
			}
			fmt.Println("Please select less than 5 cards")
			fmt.Println()
		}

		cli.service.SelectCards(selectCards)
		fmt.Print("[Selected cards]\n")
		for _, card := range selectCards {
			fmt.Println(card)
		}
		fmt.Println()

		// Play or Discard or Cancel
		actions := cli.service.GetEnableActions()
		prompt := &survey.Select{
			Message: "Select action:",
			Options: actions,
		}
		if err := survey.AskOne(prompt, &selectActioin); err == terminal.InterruptErr {
			fmt.Println("interrupted")
			os.Exit(0)
		}

		if selectActioin == "Discard" {
			cli.service.DiscardHand()
			continue
		}
		if selectActioin == "Cancel" {
			cli.service.CancelHand()
			continue
		}
		if selectActioin == "Play" {
			r, err := cli.service.PlayHand()
			if err != nil {
				return err
			}

			fmt.Printf("\nHand: %s", r.HandType)
			fmt.Printf("\nChip: %d, Mult: %d\n", r.Chip, r.Mult)
			fmt.Printf("\nScore: %d\n\n", r.Score)

			time.Sleep(1 * time.Second)
		} else {
			if err := cli.service.DiscardHand(); err != nil {
				return err
			}
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
		if roundResultStats.TotalScore >= roundResultStats.ScoreAtLeast {
			fmt.Printf("Score at least: %d, Round score: %d\n", roundResultStats.ScoreAtLeast, roundResultStats.TotalScore)
			fmt.Println("You win!")
			break
		} else if roundResultStats.Hands <= 0 {
			fmt.Printf("Score at least: %d, Round score: %d\n", roundResultStats.ScoreAtLeast, roundResultStats.TotalScore)
			fmt.Println("You lose!")
		} else {
			if roundResultStats.Hands > 0 && roundResultStats.TotalScore < roundResultStats.ScoreAtLeast {
				continue
			}
		}
	}

	return nil
}
