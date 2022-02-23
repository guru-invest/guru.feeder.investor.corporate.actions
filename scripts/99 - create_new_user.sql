CREATE USER guru_corporate_actions WITH ENCRYPTED PASSWORD <mypass>;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA financial, wallet TO guru_corporate_actions;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA financial, wallet TO guru_corporate_actions;