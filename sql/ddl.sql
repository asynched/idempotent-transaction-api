CREATE TABLE IF NOT EXISTS "accounts" (
  -- UUID
  "id" VARCHAR(36) PRIMARY KEY NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "cpf" VARCHAR(11) UNIQUE NOT NULL,
  "balance" DECIMAL(10, 2) NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transactions" (
  -- UUID
  "id" VARCHAR(36) PRIMARY KEY NOT NULL,
  "amount" DECIMAL(10, 2) NOT NULL,
  "payer_id" VARCHAR(36) NOT NULL,
  "payee_id" VARCHAR(36) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "fk_payer_id" FOREIGN KEY ("payer_id") REFERENCES "accounts" ("id"),
  CONSTRAINT "fk_payee_id" FOREIGN KEY ("payee_id") REFERENCES "accounts" ("id")
);

CREATE INDEX IF NOT EXISTS "payer_id" ON "transactions" ("payer_id");

CREATE INDEX IF NOT EXISTS "payee_id" ON "transactions" ("payee_id");