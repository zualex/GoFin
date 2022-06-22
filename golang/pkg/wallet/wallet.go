package wallet

type Wallet struct {
	Name     string `db:"name" form:"name" binding:"required"`
	Currency string `db:"currency" form:"currency" binding:"required"`
}
