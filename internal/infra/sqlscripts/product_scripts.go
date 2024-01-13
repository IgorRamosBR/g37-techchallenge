package sqlscripts

const GetAllProductsQuery = `
	SELECT 
		p.id,
		p.name, 
		p.sku_id, 
		p.description,
		p.category,
		p.price,
		p.created_at,
		p.updated_at
	FROM public.products as p
	ORDER BY p.name ASC
	LIMIT %d OFFSET %d
`

const GetProductsByCategoryQuery = `
	SELECT 
		p.id,
		p.name, 
		p.sku_id, 
		p.description,
		p.category,
		p.price,
		p.created_at,
		p.updated_at
	FROM public.products as p
	WHERE p.category = '%s'
	ORDER BY p.name ASC
	LIMIT %d OFFSET %d
`

const GetProductByIdQuery = `
	SELECT 
		p.id,
		p.name, 
		p.sku_id, 
		p.description,
		p.category,
		p.price,
		p.created_at,
		p.updated_at
	FROM public.products as p
	WHERE p.id = %d
`

const InsertProductCmd = `
	INSERT INTO public.products(name, sku_id, description, category, price, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

const UpdateProductCmd = `
	UPDATE public.products
	SET name = $2, sku_id = $3, description = $4, category = $5, price = $6, created_at = $7, updated_at = $8
	WHERE id = $1
`

const DeleteProductCmd = `
	DELETE FROM public.products
	WHERE id = $1
`
