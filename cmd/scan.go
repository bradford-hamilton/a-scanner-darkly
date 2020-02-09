package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// scanme.nmap.org
const portValidationErr = "Please ensure to request port numbers in the format `n-n`. For example `--ports 1-1024`. Available ports: 1 - 65535"

type scanJob struct {
	port    int
	message string
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "asd scan https://google.com [--ports=1-1024]",
	Long: `
asd (A Scanner Darkly) is a simple and fast port scanner. 
`,
	Args: cobra.MinimumNArgs(1),
	Run:  runScanCmd,
}

func runScanCmd(cmd *cobra.Command, args []string) {
	if Ports == "" {
		fmt.Println(portValidationErr)
		os.Exit(1)
	}

	startPort, endPort, err := fetchAndValidatePorts()
	if err != nil {
		fmt.Println(portValidationErr)
		os.Exit(1)
	}

	target := args[0]
	numWorkers := 100 // Make this adjustable
	numPorts := endPort - startPort
	jobChan := make(chan scanJob, numWorkers)
	resultsChan := make(chan scanJob, numPorts)

	createWorkerPool(target, numWorkers, jobChan, resultsChan)
	distributePortsToJobs(jobChan, startPort, endPort)

	var openPorts []string
	for i := 0; i < numPorts; i++ {
		if scanJob := <-resultsChan; scanJob.port != 0 {
			openPorts = append(openPorts, scanJob.message)
		}
	}
	close(resultsChan)

	for _, portMsg := range openPorts {
		fmt.Printf(portMsg)
	}
}

func createWorkerPool(target string, numWorkers int, jobChan, resultsChan chan scanJob) {
	for i := 0; i < numWorkers; i++ {
		go worker(jobChan, resultsChan, target)
	}
}

func distributePortsToJobs(jobChan chan scanJob, startPort, endPort int) {
	go func() {
		for i := startPort; i < endPort; i++ {
			jobChan <- scanJob{port: i}
		}
		close(jobChan)
	}()
}

func fetchAndValidatePorts() (int, int, error) {
	// Break apart the ports to scan from the ports flag.
	// Ex. --ports 1-1024 we need 1 and 1024 for startPort/endPort as ints
	ports := strings.Split(Ports, "-")
	if len(ports) != 2 {
		fmt.Println(portValidationErr)
		os.Exit(1)
	}

	// Get the starting and ending ports to scan
	startPort, err := strconv.Atoi(ports[0])
	if err != nil {
		return 0, 0, err
	}
	endPort, err := strconv.Atoi(ports[1])
	if err != nil {
		return 0, 0, err
	}

	// Ensure ports are between 1 - 65535
	if startPort < 1 || endPort > 65535 {
		return 0, 0, errors.New("Ports must be between 1 - 65535")
	}

	return startPort, endPort, nil
}

func worker(jobChan, resultsChan chan scanJob, target string) {
	for j := range jobChan {
		// Small throttle. Update this to be adjustable
		time.Sleep(250 * time.Millisecond)

		// Try to connecting to the address
		addr := fmt.Sprintf("%s:%d", target, j.port)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			resultsChan <- scanJob{port: 0}
			continue
		}
		defer conn.Close()

		// Message in format 10.57.0.110:49546 open
		j.message = fmt.Sprintf("%s open\n", conn.LocalAddr().String())

		// Place open port number on the results channel
		resultsChan <- j
	}
}
