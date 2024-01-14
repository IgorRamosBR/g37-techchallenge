package sqlscripts

const FindAllOrdersQuery = `
	SELECT 
		o.id,
		o.coupon,
		o.total_amount,
		o.status,
		o.created_at,
		c.id,
		c.name, 
		c.cpf, 
		c.email,
		c.created_at,
		c.updated_at	
	FROM public.orders o
	LEFT JOIN public.customers c ON o.customer_id = c.id
	ORDER BY ORDER BY array_position(array['READY','IN_PROGRESS','RECEIVED'], o.status), o.created_at ASC
	WHERE status <> DONE
	LIMIT $1 OFFSET $2
`

const FindOrderStatusByIdQuery = `
	SELECT
		o.status,
	FROM public.orders o
	WHERE o.id = $1	
`

const InsertOrderCmd = `
	INSERT INTO public.orders(coupon, total_amount, customer_id, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
`

const InsertOrderItemCmd = `
	INSERT INTO public.order_items(order_id, product_id, quantity, type)
	VALUES ($1, $2, $3, $4)
`

const UpdateOrderStatusCmd = `
	UPDATE public.orders
	SET status = $2
	WHERE id = $1
`
