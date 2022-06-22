package wallet

type Wallet struct {
	Id       string `db:"id" form:"id"`
	Name     string `db:"name" form:"name" binding:"required"`
	Currency string `db:"currency" form:"currency" binding:"required"`
}
