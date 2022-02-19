ALTER TABLE wallet.manual_transactions ADD post_event_symbol text NULL ;
ALTER TABLE wallet.manual_transactions ADD event_factor numeric default 1;
ALTER TABLE wallet.manual_transactions ADD post_event_quantity int4 NULL;
ALTER TABLE wallet.manual_transactions ADD post_event_price numeric NULL;
ALTER TABLE wallet.manual_transactions ADD event_date timestamp default '2001-01-01';
ALTER TABLE wallet.manual_transactions ADD event_name text default 'PADRAO';

CREATE or replace FUNCTION wallet.manual_transactions_after_insert() RETURNS trigger AS $$
    BEGIN
        NEW.post_event_quantity := NEW.quantity;
        NEW.post_event_price := NEW.price;
        NEW.post_event_symbol := NEW.symbol;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER manual_transactionss_trigger BEFORE INSERT ON wallet.manual_transactions
    FOR EACH ROW EXECUTE FUNCTION wallet.manual_transactions_after_insert();