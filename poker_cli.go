package pkr

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

var clear map[string]func()

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearTerminal() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func (cli *PokerCLI) Run() error {
	sleepSec := 1
	ClearTerminal()

	fmt.Println("*********************")
	fmt.Println("* Welcome to Poker! *")
	fmt.Println("*********************")
	fmt.Println()
	time.Sleep(time.Duration(sleepSec) * time.Second)

	for {
		ClearTerminal()
		if cli.service.IsStartRound() {
			rounds := cli.service.GetRounds()
			fmt.Println("")
			fmt.Printf("Round %d start\n\n", rounds)
			time.Sleep(time.Duration(sleepSec) * time.Second)

			ante := cli.service.GetCurrentAnteAmount()
			blind := cli.service.GetCurrentBlindMulti()
			fmt.Printf("Ante:%d, Blind:%v\n\n", ante, blind)
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
		cli.service.SetAction(selectAction)
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

		stats := cli.service.GetRoundStats()
		if cli.service.IsRoundWin() {
			fmt.Printf("Score at least: %d, Round score: %d\n", stats.ScoreAtLeast, stats.TotalScore)
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

		if stats.Hands == 0 {
			fmt.Printf("Score at least: %d, Round score: %d\n", stats.ScoreAtLeast, stats.TotalScore)
			fmt.Println("You lose!")
			break
		}

	}

	return nil
}
