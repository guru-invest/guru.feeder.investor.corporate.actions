CREATE TABLE wallet.oms_proventos (
	id serial4 NOT NULL,
	customer_code text NULL,
	broker_id numeric NULL,
	symbol text NULL,
	quantity numeric NULL,
	value numeric NULL,
	amount numeric NULL,
	"date" timestamp NULL,
	"event" text NULL,
	CONSTRAINT oms_proventos_pkey PRIMARY KEY (id)
);