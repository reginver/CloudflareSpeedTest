package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	version = "2.2.5"
	defaultRoutines  = 200
	defaultPingTimes = 4
	defaultMaxDelay  = 300
	defaultMinSpeed  = 0
	defaultOutputNum = 20
)

var (
	// Test configuration flags
	routines   int
	pingTimes  int
	maxDelay   int
	minDelay   int
	minSpeed   float64
	outputNum  int
	noOutput   bool
	ipFile     string
	outputFile string
	testAll    bool
	ipv6Mode   bool
	printVersion bool
)

func init() {
	flag.IntVar(&routines, "n", defaultRoutines, "Latency test threads; the number of concurrent goroutines")
	flag.IntVar(&pingTimes, "t", defaultPingTimes, "Latency test times; number of pings per IP")
	flag.IntVar(&maxDelay, "tl", defaultMaxDelay, "Upper latency limit (ms); filter IPs above this threshold")
	flag.IntVar(&minDelay, "tll", 0, "Lower latency limit (ms); filter IPs below this threshold")
	flag.Float64Var(&minSpeed, "sl", defaultMinSpeed, "Lower download speed limit (MB/s); filter IPs below this speed")
	flag.IntVar(&outputNum, "p", defaultOutputNum, "Number of results to display after testing")
	flag.BoolVar(&noOutput, "dd", false, "Disable download speed test")
	flag.StringVar(&ipFile, "f", "ip.txt", "IP range file path (supports CIDR notation)")
	flag.StringVar(&outputFile, "o", "result.csv", "Output file path for results (CSV format)")
	flag.BoolVar(&testAll, "allip", false, "Test all IPs in CIDR range (default tests one IP per /24 block)")
	flag.BoolVar(&ipv6Mode, "ipv6", false, "Enable IPv6 mode")
	flag.BoolVar(&printVersion, "v", false, "Print version information")

	// Override default usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "CloudflareSpeedTest v%s\n", version)
		fmt.Fprintf(os.Stderr, "Usage: CloudflareST [options]\n\nOptions:\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if printVersion {
		fmt.Printf("CloudflareSpeedTest v%s\n", version)
		return
	}

	fmt.Printf("# CloudflareSpeedTest v%s\n", version)

	// Validate arguments
	if routines <= 0 {
		fmt.Println("[Error] Latency test threads (-n) must be greater than 0")
		os.Exit(1)
	}
	if pingTimes <= 0 {
		fmt.Println("[Error] Latency test times (-t) must be greater than 0")
		os.Exit(1)
	}
	if outputNum <= 0 {
		fmt.Println("[Error] Output number (-p) must be greater than 0")
		os.Exit(1)
	}

	startTime := time.Now()

	fmt.Println("\nStarting latency test...")
	// TODO: Run latency test against Cloudflare IPs
	// results := task.NewLatencyTest(ipFile, routines, pingTimes, maxDelay, minDelay, ipv6Mode, testAll).Run()

	if !noOutput {
		fmt.Println("\nStarting download speed test...")
		// TODO: Run download speed test on top latency results
		// results = task.NewSpeedTest(results).Run()
	}

	// TODO: Filter results and write to output
	// filtered := utils.FilterResults(results, maxDelay, minDelay, minSpeed)
	// utils.PrintResults(filtered, outputNum)
	// if outputFile != "" {
	//     utils.WriteCSV(filtered, outputFile)
	// }

	fmt.Printf("\nTest completed in %v\n", time.Since(startTime))
}
