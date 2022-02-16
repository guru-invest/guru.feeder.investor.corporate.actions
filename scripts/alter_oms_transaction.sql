ALTER TABLE wallet.oms_transactions ADD post_event_symbol text NULL;
ALTER TABLE wallet.oms_transactions ADD event_factor numeric NULL;
ALTER TABLE wallet.oms_transactions ADD post_event_quantity int4 NULL;
ALTER TABLE wallet.oms_transactions ADD post_event_price numeric NULL;
ALTER TABLE wallet.oms_transactions ADD event_date timestamp NULL;
ALTER TABLE wallet.oms_transactions ADD event_name text NULL;