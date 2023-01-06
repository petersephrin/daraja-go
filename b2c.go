package darajago

//B2C API is an API used to make payments from a Business to Customers (Pay Outs).
//Also known as Bulk Disbursements. B2C API is used in several scenarios by
//businesses that require to either make Salary Payments, Cashback payments,
//Promotional Payments(e.g. betting winning payouts), winnings, financial
//institutions withdrawal of funds, loan disbursements etc.

//NOTE: For you to use this API on production you are required to apply
//for a Bulk Disbursement Account and get a Short code, you cannot do this payment from a
//Pay Bill or Buy Goods (Till Number). To apply for a Bulk disbursement account
//follow this link. https://www.safaricom.co.ke/business/sme/m-pesa-payment-solutions

// B2CPayload represents a request payload for the B2C API.
type B2CPayload struct {
	// InitiatorName is the initiator name.
	// This is the API operator's username as set on the portal when the user was created.
	// For Sandbox users, the username is already created and assigned to them and is available
	// on the test credentials page as Initiator Name (Shortcode 1).
	InitiatorName string `json:"InitiatorName"`

	// SecurityCredential is the security credential.
	SecurityCredential string `json:"SecurityCredential"`

	// CommandID eg.
	//
	//· SalaryPayment
	//
	//· BusinessPayment
	//
	//· PromotionPayment
	CommandID string `json:"CommandID"`

	// Amount is the amount to be transferred.
	Amount string `json:"Amount"`

	// PartyA is the party A (the organization making the payment).
	// This is the B2C organization shortcode from which the money is sent from.
	PartyA string `json:"PartyA"`

	// PartyB is the party B (the customer receiving the payment).
	// This is the customer mobile number to receive the amount.
	//-The number should have the country code (254) without the plus sign.
	PartyB string `json:"PartyB"`

	// Remarks are any remarks for the request.
	Remarks string `json:"Remarks"`

	// QueueTimeOutURL is the queue timeout URL.
	// This is the URL to be specified in your request that will be used by API Proxy
	// to send notification in case the payment request is timed out while awaiting processing in the queue.
	QueueTimeOutURL string `json:"QueueTimeOutURL"`

	// ResultURL is the result URL.
	ResultURL string `json:"ResultURL"`

	// Occasion is the occasion for the payment.
	// Any additional information to be associated with the transaction.
	Occasion string `json:"Occasion"`
}

// B2CResponse is the response from the B2C API
type B2CResponse struct {
	// OriginatorConversationID is the unique request ID for tracking a transaction.
	OriginatorConversationID string `json:"OriginatorConversationID"`

	// ConversationID is the unique request ID returned by Mpesa for each request made.
	ConversationID string `json:"ConversationID"`

	// ResponseDescription is the response description message.
	ResponseDescription string `json:"ResponseDescription"`
}

// MakeB2CPayment is a function that makes a B2C payment using the Mpesa API.
// It takes a B2CPayload struct representing the payment request payload as input,
// and returns a pointer to a B2CResponse struct representing the payment response,
// or a pointer to an ErrorResponse struct representing an error that occurred during the request.
func (d *DarajaApi) MakeB2CPayment(b2c B2CPayload) (*B2CResponse, *ErrorResponse) {
	b2c.CommandID = "BusinessPayment"
	// marshal the struct into a map
	secureResponse, err := performSecurePostRequest[*B2CResponse](b2c, endpointB2CPmtReq, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
