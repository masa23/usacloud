package funcs

import "errors"

func getIPv6AddrFromArgs(args []string) (string, error) {

	if len(args) == 0 {
		return "", errors.New("args[IPAddress] is required")
	}

	ip := args[0]

	if errs := validateIPv6AddressArgs(ip); len(errs) > 0 {
		return "", errs[0]
	}

	return ip, nil
}
