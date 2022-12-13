ALTER TABLE wallet.cei_proventos DROP COLUMN "date";
ALTER TABLE wallet.cei_proventos ADD initial_date timestamp NULL;
ALTER TABLE wallet.cei_proventos ADD com_date timestamp NULL;
ALTER TABLE wallet.cei_proventos ADD payment_date timestamp NULL;

ALTER TABLE wallet.oms_proventos DROP COLUMN "date";
ALTER TABLE wallet.oms_proventos ADD initial_date timestamp NULL;
ALTER TABLE wallet.oms_proventos ADD com_date timestamp NULL;
ALTER TABLE wallet.oms_proventos ADD payment_date timestamp NULL;

ALTER TABLE wallet.manual_proventos DROP COLUMN "date";
ALTER TABLE wallet.manual_proventos ADD initial_date timestamp NULL;
ALTER TABLE wallet.manual_proventos ADD com_date timestamp NULL;
ALTER TABLE wallet.manual_proventos ADD payment_date timestamp NULL;