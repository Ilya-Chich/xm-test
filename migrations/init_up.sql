DROP TABLE IF EXISTS companies;
CREATE TABLE companies (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   name VARCHAR(15) UNIQUE NOT NULL,
   description VARCHAR(3000),
   amount_of_employees INT NOT NULL,
   registered BOOLEAN NOT NULL,
   type VARCHAR(20) NOT NULL CHECK (type IN ('Corporations', 'NonProfit', 'Cooperative', 'Sole Proprietorship'))
);
INSERT INTO companies (name, description, amount_of_employees, registered, type)
VALUES ('XM', 'A leading tech company.', 500, true, 'Corporations');