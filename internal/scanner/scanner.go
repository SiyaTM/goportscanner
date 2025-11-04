package scanner

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Scanner checks if a TCP port is open/closed/filtered
func ScanPort(host string, port int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	address := host + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			results <- fmt.Sprintf("Port %d: filtered (timeout)", port)
		} else {
			results <- fmt.Sprintf("Port %d: closed", port)
		}
		return
	}

	results <- fmt.Sprintf("Port %d: open", port)
	conn.Close()

}

// ParsePortRange converts "20-1000" into start/end integers.
func ParsePortRange(r string) (int, int, error) {
	parts := strings.Split(r, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range format: %s", r)
	}
	start, end := parts[0], parts[1]
	s, err1 := strconv.Atoi(start)
	e, err2 := strconv.Atoi(end)
	if err1 != nil || err2 != nil || s < 1 || e > 65535 || s > e {
		return 0, 0, fmt.Errorf("invalid port range: %s", r)
	}
	return s, e, nil
}
