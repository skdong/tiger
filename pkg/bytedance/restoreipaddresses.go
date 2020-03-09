package bytedance

import "fmt"

func restoreIpAddresses(s string) []string {
	Ips := []string{}
	ipb := []byte(s)
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				last := len(ipb) - a - b - c
				if last < 1 || last > 3 {
					continue
				}
				ip := [][]byte{ipb[:a], ipb[a : a+b], ipb[a+b : a+b+c], ipb[a+b+c:]}
				i := 0
				for ; i < 4; i++ {
					if len(ip[i]) > 1 && ip[i][0] == '0' {
						break
					}
					if len(ip[i]) == 3 {
						by := 0
						for _, byi := range ip[i] {
							by = by*10 + int(byi-'0')
						}
						if by > 255 {
							break
						}

					}
				}
				if i == 4 {
					Ips = append(Ips, fmt.Sprintf("%s.%s.%s.%s", string(ip[0]), string(ip[1]), string(ip[2]), string(ip[3])))
				}
			}
		}
	}
	return Ips
}
