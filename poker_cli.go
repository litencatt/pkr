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
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
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

func printBox(title, content string) {
	fmt.Println("┌─────────────────────────────────────────┐")
	fmt.Printf("│ %-39s │\n", title)
	fmt.Println("├─────────────────────────────────────────┤")
	fmt.Printf("│ %-39s │\n", content)
	fmt.Println("└─────────────────────────────────────────┘")
}

func printProgressBar(current, target int) {
	barWidth := 30
	var progress float64
	if target > 0 {
		progress = float64(current) / float64(target)
		if progress > 1.0 {
			progress = 1.0
		}
	}

	filled := int(progress * float64(barWidth))
	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	percentage := int(progress * 100)
	fmt.Printf("📊 Score Progress: [%s] %d%% (%d/%d)\n", bar, percentage, current, target)
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
			ante := cli.service.GetCurrentAnteAmount()
			blind := cli.service.GetCurrentBlindMulti()

			printBox(
				fmt.Sprintf("🃏 ROUND %d START", rounds),
				fmt.Sprintf("Ante: %d  |  Blind: %.1f", ante, blind),
			)
			fmt.Println()
			time.Sleep(time.Duration(sleepSec) * time.Second)

			if err := cli.service.StartRound(); err != nil {
				return err
			}
		}

		roundStats := cli.service.GetRoundStats()
		printProgressBar(roundStats.TotalScore, roundStats.ScoreAtLeast)
		fmt.Printf("🃏 Hands: %d  |  🗑️  Discards: %d\n", roundStats.Hands, roundStats.Discards)
		fmt.Println()

		// Draw cards
		drawNum := cli.service.GetNextDrawNum()
		cards, err := cli.service.DrawCard(drawNum)
		if err != nil {
			return err
		}
		fmt.Printf("🎲 Draw %d cards\n", drawNum)
		if cli.DebugMode {
			fmt.Println("────────── Drawn Cards ──────────")
			for _, card := range cards {
				fmt.Printf("  • %s\n", card.String())
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
		fmt.Println("✅ Selected Cards:")
		if len(selectCards) > 0 {
			for _, card := range selectCards {
				fmt.Printf("  🃏 %s\n", card)
			}
		} else {
			fmt.Println("  (No cards selected)")
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

		if err := cli.service.SelectCards(selectCards); err != nil {
			return err
		}
		cli.service.SetAction(selectAction)
		if selectAction == "Discard" {
			if err := cli.service.DiscardHand(); err != nil {
				return err
			}
			continue
		}
		if selectAction == "Cancel" {
			if err := cli.service.CancelHand(); err != nil {
				return err
			}
			continue
		}
		if selectAction == "Play" {
			r, err := cli.service.PlayHand()
			if err != nil {
				return err
			}

			fmt.Println("┌─────────────────────────────────────────┐")
			fmt.Printf("│ 🎯 HAND RESULT: %-22s │\n", r.HandType)
			fmt.Println("├─────────────────────────────────────────┤")
			fmt.Printf("│ 💰 Chip: %-6d  |  ✨ Mult: %-6d │\n", r.Chip, r.Mult)
			fmt.Printf("│ 🏆 Score: %-29d │\n", r.Score)
			fmt.Println("└─────────────────────────────────────────┘")
			fmt.Println()

			time.Sleep(time.Duration(sleepSec) * time.Second)
		}

		// show remain cards
		if cli.DebugMode {
			fmt.Println("────────── Remaining Cards ──────────")
			remainCards := cli.service.GetRemainCardString()
			if len(remainCards) > 0 {
				for _, card := range remainCards {
					fmt.Printf("  • %s\n", card)
				}
			} else {
				fmt.Println("  (No remaining cards)")
			}
			fmt.Println()
		}

		stats := cli.service.GetRoundStats()
		if cli.service.IsRoundWin() {
			fmt.Println("🎉 ROUND CLEAR! 🎉")
			printProgressBar(stats.TotalScore, stats.ScoreAtLeast)
			fmt.Println()

			prompt := &survey.Select{
				Message: "🏆 You win this round! Ready for next?",
				Options: []string{"Next Round →"},
			}
			if err := survey.AskOne(prompt, &selectAction); err == terminal.InterruptErr {
				fmt.Println("interrupted")
				os.Exit(0)
			}
			if err := cli.service.NextRound(); err != nil {
				return err
			}
			continue
		}

		if stats.Hands == 0 {
			fmt.Println("💀 GAME OVER 💀")
			printProgressBar(stats.TotalScore, stats.ScoreAtLeast)
			fmt.Println("😢 Better luck next time!")
			break
		}

	}

	return nil
}
