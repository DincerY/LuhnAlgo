package luhn

import "fmt"

func RunAlgo(a string) string {
	var sum uint8
	for i := 0; i < len(a); i += 2 {
		b := a[i] - 48
		b = b * 2
		if b >= 10 {
			temp := b % 10
			b = temp + (b / 10)
		}
		if i != 0 {
			sum += a[i-1] - 48
		}
		sum += b
	}
	d := a[len(a)-1] - 48
	if (sum+d)%10 == 0 {
		return "valid card number"
	} else {
		return fmt.Sprintf("invalid card number last number must be %d", (10-(sum%10))%10)
	}
}
