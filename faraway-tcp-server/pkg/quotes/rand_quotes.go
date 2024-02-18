package quotes

import "math/rand"

// generate random quote for response
func GetQuote() string {
	quotes := []string{
		"The only way to do great work is to love what you do. - Steve Jobs",
		"Strive not to be a success, but rather to be of value. - Albert Einstein",
		"Believe you can and you're halfway there. - Theodore Roosevelt",
		"It does not matter how slowly you go as long as you do not stop. - Confucius",
		"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
		"“Stay away from those people who try to disparage your ambitions. Small minds will always do that, " +
			"but great minds will give you a feeling that you can become great too.” — Mark Twain",
		"“When you give joy to other people, you get more joy in return. You should give a good thought to happiness that you can give out.”— Eleanor Roosevelt",
		"Success is not final; failure is not fatal: It is the courage to continue that counts. — Winston S. Churchill",
		"The pessimist sees difficulty in every opportunity. The optimist sees opportunity in every difficulty. — Winston Churchill",
	}
	quoteIndex := rand.Intn(len(quotes))
	return quotes[quoteIndex]
}
