package statements

const GetTransactionsByCustomerCodes = `
	SELECT 
		id, 
		customer_code, 
		broker_id, 
		symbol, 
		quantity, 
		price, 
		amount, 
		side, 
		trade_date, 
		post_event_quantity, 
		post_event_price, 
		post_event_symbol, 
		event_factor, 
		event_date, 
		event_name
	FROM 
		wallet.oms_transactions
	WHERE
		customer_code in @customerCodes
	ODER BY
		trade_date ASC
`
