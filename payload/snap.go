package payload

type (
	AcquireTokenSnapRequest struct {
		TransactionDetailsSnap TransactionDetailsSnap `json:"transaction_details"`
		CustomerDetailsSnap    CustomerDetailsSnap    `json:"customer_details"`
	}

	AcquireTokenSnapResponse struct {
		Token       string `json:"token"`
		RedirectUrl string `json:"redirect_url"`
	}

	TransactionDetailsSnap struct {
		OrderId     string `json:"orderId"`
		GrossAmount int    `json:"gross_amount"`
	}

	CustomerDetailsSnap struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"las_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
)
