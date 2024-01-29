CREATE TABLE IF NOT EXISTS public.products (
	"id" serial primary key,
	"name" text not null,
	"sku_id" text,
	"description" text,
	"category" text,
	"price" numeric,
	"created_at" timestamptz not null,
	"updated_at" timestamptz not null
);