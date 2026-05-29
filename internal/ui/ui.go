package ui

import (
	"fmt"
	"time"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
	Reset  = "\033[0m"
)

func PrintBanner() {
	banner := `
█     █░ ██░ ██  ▒█████      ██▓  ██████    ▓██   ██▓ ▒█████   █    ██  ██▀███     ▓█████▄  ▄▄▄      ▓█████▄ ▓█████▄▓██   ██▓
▓█░ █ ░█░▓██░ ██▒▒██▒  ██▒   ▓██▒▒██    ▒     ▒██  ██▒▒██▒  ██▒ ██  ▓██▒▓██ ▒ ██▒   ▒██▀ ██▌▒████▄    ▒██▀ ██▌▒██▀ ██▌▒██  ██▒
▒█░ █ ░█ ▒██▀▀██░▒██░  ██▒   ▒██▒░ ▓██▄        ▒██ ██░▒██░  ██▒▓██  ▒██░▓██ ░▄█ ▒   ░██   █▌▒██  ▀█▄  ░██   █▌░██   █▌ ▒██ ██░
░█░ █ ░█ ░▓█ ░██ ▒██   ██░   ░██░  ▒   ██▒     ░ ▐██▓░▒██   ██░▓▓█  ░██░▒██▀▀█▄     ░▓█▄   ▌░██▄▄▄▄██ ░▓█▄   ▌░▓█▄   ▌ ░ ▐██▓░
░░██▒██▓ ░▓█▒░██▓░ ████▓▒░   ░██░▒██████▒▒     ░ ██▒▓░░ ████▓▒░▒▒█████▓ ░██▓ ▒██▒   ░▒████▓  ▓█   ▓██▒░▒████▓ ░▒████▓  ░ ██▒▓░
░ ▓░▒ ▒   ▒ ░░▒░▒░ ▒░▒░▒░    ░▓  ▒ ▒▓▒ ▒ ░      ██▒▒▒ ░ ▒░▒░▒░ ░▒▓▒ ▒ ▒ ░ ▒▓ ░▒▓░    ▒▒▓  ▒  ▒▒   ▓▒█░ ▒▒▓  ▒  ▒▒▓  ▒   ██▒▒▒ 
  ▒ ░ ░   ▒ ░▒░ ░  ░ ▒ ▒░     ▒ ░░ ░▒  ░ ░    ▓██ ░▒░   ░ ▒ ▒░ ░░▒░ ░ ░   ░▒ ░ ▒░    ░ ▒  ▒   ▒   ▒▒ ░ ░ ▒  ▒  ░ ▒  ▒ ▓██ ░▒░ 
  ░   ░   ░  ░░ ░░ ░ ░ ▒      ▒ ░░  ░  ░      ▒ ▒ ░░  ░ ░ ░ ▒   ░░░ ░ ░   ░░   ░     ░ ░  ░   ░   ▒    ░ ░  ░  ░ ░  ░ ▒ ▒ ░░  
    ░     ░  ░  ░    ░ ░      ░        ░      ░ ░         ░ ░     ░        ░           ░          ░  ░   ░       ░    ░ ░     
                                              ░ ░                                    ░                 ░       ░      ░ ░     `

	fmt.Println(Red + banner + Reset)
	fmt.Printf("%s  v0.1.0%s — attack surface mapper\n", Bold, Reset)
	fmt.Printf("%s  by OnlyManuel%s\n\n", Yellow, Reset)
}

func PrintHelp() {
	PrintBanner()
	fmt.Printf("%sUSAGE%s\n", Bold, Reset)
	fmt.Printf("  wiyd %s-target%s <domain>\n\n", Cyan, Reset)
	fmt.Printf("  wiyd %s-target%s <domain> %s-wordlist%s <wordlist>\n\n", Cyan, Reset, Cyan, Reset)
	fmt.Printf("  wiyd %s-target%s <domain> %s-wordlist%s <wordlist> %s-output%s report.txt \n\n", Cyan, Reset, Cyan, Reset, Cyan, Reset)

	fmt.Printf("%sOPTIONS%s\n", Bold, Reset)
	fmt.Printf("  %s-target%s    target domain to enumerate\n", Cyan, Reset)
	fmt.Printf("  %s-wordlist%s  external wordlist to use\n", Cyan, Reset)
	fmt.Printf("  %s-output%s    wether to save output\n", Cyan, Reset)
	fmt.Printf("  %s-h%s         show this help menu\n\n", Cyan, Reset)

	fmt.Printf("%sEXAMPLES%s\n", Bold, Reset)
	fmt.Printf("  wiyd -target example.com\n")
	fmt.Printf("  wiyd %s-target%s hackthebox.com %s-wordlist%s seclist/common_domains.txt \n\n", Cyan, Reset, Cyan, Reset)
	fmt.Printf("  wiyd %s-target%s hackthebox.com %s-wordlist%s seclist/common_domains.txt %s-output%s report\n\n", Cyan, Reset, Cyan, Reset)

	fmt.Printf("%sMODULES%s\n", Bold, Reset)
	fmt.Printf("  %scrt.sh%s      certificate transparency subdomain enumeration\n", Yellow, Reset)
	fmt.Printf("  %sdns-enum%s    dns bruteforce with builtin wordlist\n\n", Yellow, Reset)
}

func Spinner(label string, done chan bool) {
	frames := []string{".", "..", "..."}
	i := 0
	for {
		select {
		case <-done:
			fmt.Printf("\r%s[+] %s done%s          \n", Green, label, Reset)
			return
		default:
			fmt.Printf("\r%s[~] %s%s%s", Cyan, label, frames[i%3], Reset)
			i++
			time.Sleep(300 * time.Millisecond)
		}
	}
}
