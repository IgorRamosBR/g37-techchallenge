CREATE TABLE IF NOT EXISTS public.customers (
	"id" serial primary key,
	"created_at" timestamptz not null,
	"updated_at" timestamptz not null,
	"name" text not null,
	"cpf" text not null,
	"email" text not null
);

CREATE TABLE IF NOT EXISTS public.products (
	"id" serial primary key,
	"created_at" timestamptz not null,
	"updated_at" timestamptz not null,
	"name" text not null,
	"sku_id" text,
	"description" text,
	"category" text,
	"price" numeric
);

CREATE TABLE IF NOT EXISTS public.orders (
	"id" serial primary key,
	"created_at" timestamptz not null,
	"updated_at" timestamptz not null,
	"coupon" text NULL,
	"discount" numeric NULL,
	"total_amount" numeric NULL,
	"customer_id" integer NULL,
	"status" text NULL,
	CONSTRAINT "FK_customers_orders" FOREIGN KEY (customer_id) REFERENCES public.customers(id)
);

CREATE TABLE IF NOT EXISTS public.order_items (
	"id" serial primary key,
	"product_id" integer,
	"quantity" integer,
	"type" text,
	"order_id" int,
	CONSTRAINT "FK_order_items_product" FOREIGN KEY (product_id) REFERENCES public.products(id),
	CONSTRAINT "FK_orders_items" FOREIGN KEY (order_id) REFERENCES public.orders(id)
);

