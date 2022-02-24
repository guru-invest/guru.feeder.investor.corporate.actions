CREATE TABLE wallet.cei_proventos (
	id text NOT NULL,
	customer_code text NULL,
	broker_id numeric NULL,
	symbol text NULL,
	quantity numeric NULL,
	value numeric NULL,
	amount numeric NULL,
	"date" timestamp NULL,
	"event" text NULL,
	CONSTRAINT cei_proventos_pkey PRIMARY KEY (id)
);