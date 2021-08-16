package contracts

type CreditAssigner interface {
	Asssign(investment int32) (int32, int32, int32, error)
}
