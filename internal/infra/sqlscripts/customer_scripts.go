package sqlscripts

const GetCustomerByIdQuery = `
	SELECT 
		c.id,
		c.name, 
		c.cpf, 
		c.email,
		c.created_at,
		c.updated_at,
	FROM public.customers as c
	WHERE c.id = $1
`

const GetCustomerByCPFQuery = `
	SELECT 
		c.id,
		c.name, 
		c.cpf, 
		c.email,
		c.created_at,
		c.updated_at,
	FROM public.customers as c
	WHERE c.cpf = $1
`

const InsertCustomer = `
	INSERT INTO public.customers(name, cpf, email, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
`
