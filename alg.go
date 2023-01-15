package solver

func invertMove(move string) string {
	switch move {
	case "U":
		return "U'"
	case "U'":
		return "U"
	case "U2":
		return "U2"
	case "F":
		return "F'"
	case "F'":
		return "F"
	case "F2":
		return "F2"
	case "D":
		return "D'"
	case "D'":
		return "D"
	case "D2":
		return "D2"
	case "B":
		return "B'"
	case "B'":
		return "B"
	case "B2":
		return "B2"
	case "L":
		return "L'"
	case "L'":
		return "L"
	case "L2":
		return "L2"
	case "R":
		return "R'"
	case "R'":
		return "R"
	case "R2":
		return "R2"
	}
	return ""
}

func invertAlg(alg []string) []string {
	inverse := make([]string, len(alg))
	for i, move := range alg {
		inverse[len(alg)-i-1] = invertMove(move)
	}
	return inverse
}
