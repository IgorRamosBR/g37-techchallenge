CREATE TABLE IF NOT EXISTS public.customers (
	"id" serial primary key,
	"name" text not null,
	"cpf" text not null,
	"email" text not null,
	"created_at" timestamptz not null,
	"updated_at" timestamptz not null
);