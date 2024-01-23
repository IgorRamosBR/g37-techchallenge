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
	WHERE o.status <> 'DONE'
	ORDER BY array_position(array['READY','IN_PROGRESS','RECEIVED'], o.status), o.created_at ASC
	LIMIT $1 OFFSET $2
`

const FindOrderItems = `
	SELECT
		oi.id,
		p.id,
		p.name, 
		p.sku_id, 
		p.description,
		p.category,
		p.price,
		p.created_at,
		p.updated_at,
		oi.quantity,
		oi.type
	FROM public.order_items oi
	LEFT JOIN public.products p ON oi.product_id = p.id
	WHERE oi.order_id = $1
`

const FindOrderStatusByIdQuery = `
	SELECT
		o.status
	FROM public.orders o
	WHERE o.id = $1	
`

const InsertOrderCmd = `
	INSERT INTO public.orders(coupon, total_amount, customer_id, status, created_at)
	VALUES ($1, $2, $3, $4, $5) RETURNING id
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
