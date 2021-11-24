package qdiif

import (
	"github.com/PumpkinSeed/iif"
)

type TransactionHeader struct {
	// A unique number that identifies the transaction.
	TransactionID int `iif:"TRNSID"`

	//  A second unique number that works with TRNSID to identify the transaction.
	Timestamp Time `iif:"TIMESTAMP"`

	// (Required) A keyword that identifies the type of transaction. These
	// keywords can appear in the TRNSTYPE field:
	// - BEGINBALCHECK
	//   Transactions that create a beginning balance in a balance sheet account
	// - BILL
	//   Bills from vendors
	// - BILL REFUND
	//   Refunds from a vendor
	// - CASH REFUND
	//   Cash refunds you give to customers
	// - CASH SALE
	//   Sales receipts
	// - CCARD REFUND
	//   Refunds you receive on credit card charges
	// - CHECK
	//   Checks
	// - CREDIT CARD
	//   Charges you make on a credit card
	// - CREDIT MEMO
	//   Credit you give to customers for merchandise they return
	// - DEPOSIT
	//   Bank or money market deposits
	// - ESTIMATES
	//   Estimates or bids
	// - GENERAL JOURNAL
	//   General journal transactions
	// - INVOICE
	//   Invoices
	// - PAYMENT
	//   Customer payments
	// - PURCHORD
	//   Purchase orders
	// - TRANSFER
	//   Transfers of funds from one balance sheet account to another
	TransactionType string `iif:"TRNSTYPE"`

	// The date of the transaction. The date is always in MM/DD/YY format. For
	// example, 1/30/94.
	Date Date `iif:"DATE"`

	// (Required) The name of the account assigned to the transaction.)
	Account string `iif:"ACCNT"`

	// The name of the customer, vendor, payee, or employee.
	Name string `iif:"NAME"`

	// The name of the class that applies to the transaction. If the class is a
	// subclass, the class name includes the names of the parent classes,
	// beginning with the highest level class. A colon (:) separates each class
	// name.
	Class string `iif:"CLASS"`

	// (Required) The amount of the transaction. Debit amounts are always
	// positive, credit amounts are always negative.
	Amount float64 `iif:"AMOUNT"`

	// The number of the transaction. For checks, the number is the check
	// number; for invoices, the number is the invoice number; etc.
	DocNum string `iif:"DOCNUM"`

	// The memo text associated with the transaction.
	Memo string `iif:"MEMO"`

	// Indicates whether the transaction has cleared. These keywords can appear
	// in the CLEAR field:
	// - Y
	//   Yes. The transaction has cleared.
	// - N
	//   No. The transaction hasn't cleared.
	Clear Bool `iif:"CLEAR"`

	// Indicates whether a check, invoice, credit memo, or sales receipt has
	// been marked as "To be printed." These keywords can appear in the TOPRINT
	// field:
	// - Y
	//   Yes
	// - N
	//   No
	ToPrint Bool `iif:"TOPRINT"`

	// The first line of the customer's, vendor's, payee's, or employee's address.
	Address1 string `iif:"ADDR1"`

	// The second line of the customer's, vendor's, payee's, or employee's address.
	Address2 string `iif:"ADDR2"`

	// The third line of the customer's, vendor's, payee's, or employee's address.
	Address3 string `iif:"ADDR3"`

	// The fourth line of the customer's, vendor's, payee's, or employee's address.
	Address4 string `iif:"ADDR4"`

	// The fifth line of the customer's, vendor's, payee's, or employee's address.
	Address5 string `iif:"ADDR5"`

	// (Bills and invoices only) The due date of the bill payment or invoice
	// payment. The date is always in MM/DD/YY format. For example, 1/30/98.
	DueDate Date `iif:"DUEDATE"`

	// (Invoices only) The terms of the invoice.
	Terms string `iif:"TERMS"`

	// (Invoices only) Indicates whether an invoice has been paid in full. These
	// keywords can appear in the PAID field:
	// - Y
	//   Yes. The invoice has been paid in full.
	// - N
	//   No. The invoice is either partially paid or unpaid.
	Paid Bool `iif:"PAID"`

	// (Sales receipts only) The method your customer used to pay for the merchandise.
	PaymentMethod string `iif:"PAYMETH"`

	// (Invoices and sales receipts only) The method you used to ship the merchandise.
	ShipVia string `iif:"SHIPVIA"`

	// (Invoices and sales receipts only) The shipping date. If you are creating
	// an import file, enter the date in MM/DD/YY format. For example, 1/30/94.
	ShipDate Date `iif:"SHIPDATE"`

	// (Invoices and sales receipts only) The initials of the sales
	// representative or employee who made the sale.
	Rep string `iif:"REP"`

	// (Invoices, credit memos, and cash sales only) The location where you
	// delivered the merchandise—free of charge—to a carrier for shipment.
	Fob string `iif:"FOB"`

	// (Invoices and credit memos only) The customer's purchase order number.
	PurchaseOrderNumber string `iif:"PONUM"`

	// (Invoices, credit memos, and sales receipts only) The title that appears
	// on the invoice, credit memo, or sales receipt.
	InvoiceTitle string `iif:"INVTITLE"`

	// (Invoices, credit memos, and sales receipts only) Your message to the
	// customer—as it appears on the invoice, credit memo, or sales receipt.
	InvoiceMemo string `iif:"INVMEMO"`

	// (Invoices and sales receipts only) The first line of the customer's shipping address.
	ShippingAddress1 string `iif:"SADDR1"`

	// (Invoices and sales receipts only) The second line of the customer's shipping address.
	ShippingAddress2 string `iif:"SADDR2"`

	// (Invoices and sales receipts only) The third line of the customer's shipping address.
	ShippingAddress3 string `iif:"SADDR3"`

	// (Invoices and sales receipts only) The fourth line of the customer's shipping address.
	ShippingAddress4 string `iif:"SADDR4"`

	// (Invoices and sales receipts only) The fifth line of the customer's shipping address.
	ShippingAddress5 string `iif:"SADDR5"`

	// (Invoices and sales receipts only) Indicates whether the customer whose
	// name appears in the transaction is taxable.
	// - Y
	//   Yes. The customer is taxable.
	// - N
	//   No. The customer is not taxable.
	NameIsTaxable Bool `iif:"NAMEISTAXABLE"`
}

func (th TransactionHeader) GetType() iif.Type {
	return iif.Trns
}
