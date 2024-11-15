package bank_test

import (
	"github.com/Azzerith/bank/bank"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bank Suite")
}

var _ = Describe("Bank Account", func() {
	var account bank.Account

	BeforeEach(func() {
		account = bank.Account{}
	})

	Describe("Deposit", func() {
		It("should increase the balance when the amount is valid", func() {
			err := account.Deposit(100)
			Expect(err).To(BeNil())
			Expect(account.GetBalance()).To(Equal(100.0))
		})

		It("should return an error for zero or negative amount", func() {
			err := account.Deposit(0)
			Expect(err).To(MatchError("deposit amount must be greater than zero"))

			err = account.Deposit(-10)
			Expect(err).To(MatchError("deposit amount must be greater than zero"))
		})
	})

	Describe("Withdraw", func() {
		BeforeEach(func() {
			account.Deposit(100)
		})

		It("should decrease the balance when the amount is valid", func() {
			err := account.Withdraw(50)
			Expect(err).To(BeNil())
			Expect(account.GetBalance()).To(Equal(50.0))
		})

		It("should return an error for zero or negative amount", func() {
			err := account.Withdraw(0)
			Expect(err).To(MatchError("withdrawal amount must be greater than zero"))

			err = account.Withdraw(-10)
			Expect(err).To(MatchError("withdrawal amount must be greater than zero"))
		})

		It("should return an error when the amount exceeds the balance", func() {
			err := account.Withdraw(200)
			Expect(err).To(MatchError("insufficient funds"))
		})
	})

	Describe("GetBalance", func() {
		It("should return the current balance", func() {
			account.Deposit(100)
			Expect(account.GetBalance()).To(Equal(100.0))
		})
	})
})
