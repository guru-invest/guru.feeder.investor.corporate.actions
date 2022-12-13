package statements

const CorporateActionsGetAll = `
	SELECT 
		ca.ticker, 
		ca.description, 
		ca.value, 
		ca.payment_date, 
		ca.com_date, 
		ca.target_ticker, 
		ca.calculated_factor, 
		ca.initial_date
	FROM 
		financial.corporate_actions ca
	WHERE
		ca.com_date > current_timestamp - interval '5 years' 
		and ca.description in @descriptions
	ORDER BY
		ca.com_date DESC
`
