ALTER TABLE wallet.cei_transactions ADD post_event_symbol text NULL ;
ALTER TABLE wallet.cei_transactions ADD event_factor numeric default 1;
ALTER TABLE wallet.cei_transactions ADD post_event_quantity numeric NULL;
ALTER TABLE wallet.cei_transactions ADD post_event_price numeric NULL;
ALTER TABLE wallet.cei_transactions ADD event_date timestamp default '2001-01-01';
ALTER TABLE wallet.cei_transactions ADD event_name text default 'PADRAO';

update wallet.cei_transactions set post_event_symbol = symbol, post_event_quantity = quantity, post_event_price = price;

CREATE or replace FUNCTION wallet.cei_transactions_after_insert() RETURNS trigger AS $$
    BEGIN
        NEW.post_event_quantity := NEW.quantity;
        NEW.post_event_price := NEW.price;
        NEW.post_event_symbol := NEW.symbol;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER cei_transactionss_trigger BEFORE INSERT ON wallet.cei_transactions
    FOR EACH ROW EXECUTE FUNCTION wallet.cei_transactions_after_insert();