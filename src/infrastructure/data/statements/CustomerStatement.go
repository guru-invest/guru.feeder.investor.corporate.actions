package statements

const CustomerGetAll = `
	SELECT 
		DISTINCT customer_code
	FROM 
		wallet.oms_transactions
`
