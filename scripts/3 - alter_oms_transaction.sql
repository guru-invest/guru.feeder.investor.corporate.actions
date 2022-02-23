ALTER TABLE wallet.oms_transactions ADD post_event_symbol text NULL ;
ALTER TABLE wallet.oms_transactions ADD event_factor numeric default 1;
ALTER TABLE wallet.oms_transactions ADD post_event_quantity numeric NULL;
ALTER TABLE wallet.oms_transactions ADD post_event_price numeric NULL;
ALTER TABLE wallet.oms_transactions ADD event_date timestamp default '2001-01-01';
ALTER TABLE wallet.oms_transactions ADD event_name text default 'PADRAO';



CREATE or replace FUNCTION wallet.oms_transactions_after_insert() RETURNS trigger AS $$
    BEGIN
        NEW.post_event_quantity := NEW.quantity;
        NEW.post_event_price := NEW.price;
        NEW.post_event_symbol := NEW.symbol;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER oms_transactions_trigger BEFORE INSERT ON wallet.oms_transactions
    FOR EACH ROW EXECUTE FUNCTION wallet.oms_transactions_after_insert();