CREATE TABLE IF NOT EXISTS public.customers (
	"id" serial primary key,
	"name" text not null,
	"cpf" text not null,
	"email" text not null,
	"created_at" timestamptz not null,
	"updated_at" timestamptz not null
);

CREATE TABLE IF NOT EXISTS public.products (
	"id" serial primary key,
	"name" text not null,
	"sku_id" text,
	"description" text,
	"category" text,
	"price" numeric "created_at" timestamptz not null,
	"updated_at" timestamptz not null
);

CREATE TABLE IF NOT EXISTS public.orders (
	"id" serial primary key,
	"coupon" text,
	"total_amount" numeric,
	"customer_id" integer,
	"status" text,
	"created_at" timestamptz not null,
	CONSTRAINT "FK_customers_orders" FOREIGN KEY (customer_id) REFERENCES public.customers(id)
);

CREATE TABLE IF NOT EXISTS public.order_items (
	"id" serial primary key,
	"order_id" int,
	"product_id" integer,
	"quantity" integer,
	"type" text,
	CONSTRAINT "FK_order_items_product" FOREIGN KEY (product_id) REFERENCES public.products(id),
	CONSTRAINT "FK_order" FOREIGN KEY (order_id) REFERENCES public.orders(id)
);